package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	github "github.com/anfern777/cicd-dashboard/constants"
)

type SchiService interface {
	FindAll(*gin.Context) ([]entity.SourceCodeHostIntegration, error)
	FindByProject(ctx *gin.Context, projectID uint) (entity.SourceCodeHostIntegration, error)
	Save(*gin.Context, entity.SourceCodeHostIntegration) error
	ListEnvironments(ctx *gin.Context, schi entity.SourceCodeHostIntegration) ([]Environment, error)
}

type EnvironmentsData struct {
	TotalCount   int           `json:"total_count"`
	Environments []Environment `json:"environments"`
}

type Environment struct {
	Name      string `json:"name"`
	TargetUrl string `json:"url"`
	Reachable bool   `json:"reachable"`
	Version   string `json:"version"`
}

type EnvUrlData struct {
	Value string `json:"value"`
}

type schiService struct {
}

func NewSchiService() SchiService {
	return &schiService{}
}

func (service *schiService) FindAll(ctx *gin.Context) ([]entity.SourceCodeHostIntegration, error) {
	var schis []entity.SourceCodeHostIntegration
	err := GetDB(ctx).Find(&schis).Error
	if err != nil {
		return schis, err
	}
	return schis, nil
}

func (service *schiService) FindByProject(ctx *gin.Context, projectID uint) (entity.SourceCodeHostIntegration, error) {
	var schi entity.SourceCodeHostIntegration
	err := GetDB(ctx).First(&schi, "project_id = ?", projectID).Error
	if err != nil {
		return schi, err
	}
	return schi, nil
}

func (service *schiService) Save(ctx *gin.Context, schi entity.SourceCodeHostIntegration) error {
	encryptedKey, err := encrypt(schi.Secret, []byte(os.Getenv("ENCRYPT_KEY")))
	if err != nil {
		return err
	}

	schi.Secret = encryptedKey
	session := GetDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})

	return session.Save(&schi).Error
}

func (service *schiService) ListEnvironments(ctx *gin.Context, schi entity.SourceCodeHostIntegration) ([]Environment, error) {

	decryptedSecret, err := decrypt(schi.Secret, []byte(os.Getenv("ENCRYPT_KEY")))
	if err != nil {
		return nil, err
	}

	urlEnvs, _ := url.JoinPath(github.BaseApiUrl, schi.Owner, schi.Repo, "environments")
	req, err := http.NewRequest("GET", urlEnvs, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decryptedSecret))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %s", err.Error())
	}

	var data EnvironmentsData
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return nil, fmt.Errorf("error occured during unmarshalling. %v", err)
	}

	var wg sync.WaitGroup
	for _, env := range data.Environments {
		wg.Add(1)
		go func(env Environment, envData *EnvironmentsData) {
			defer wg.Done()

			urlEnvUrl, _ := url.JoinPath(github.BaseApiUrl, schi.Owner, schi.Repo, "environments", env.Name, "variables", "URL")
			fmt.Println("URL: ", urlEnvUrl)
			req, err := http.NewRequest("GET", urlEnvUrl, nil)
			if err != nil {
				fmt.Printf("error executing request: %s", err.Error())
				return
			}
			req.Header.Set("Accept", "application/vnd.github+json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", decryptedSecret))
			req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("error executing request: %s", err.Error())
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error reading response body: %s", err.Error())
				return
			}

			var envUrlData EnvUrlData
			err = json.Unmarshal([]byte(body), &envUrlData)
			if err != nil {
				fmt.Printf("error occured during unmarshalling. %v", err)
				return
			}

			res, err := http.Get(envUrlData.Value)
			if err != nil {
				fmt.Printf("error occured performing request. %v", err)
			}

			var isReachable bool
			if res.StatusCode != 200 {
				isReachable = false
			} else {
				isReachable = true
			}

			for i, obj := range envData.Environments {
				if obj.Name == env.Name {
					envData.Environments[i].TargetUrl = envUrlData.Value
					envData.Environments[i].Reachable = isReachable
				}
			}
			resp.Body.Close()
		}(env, &data)
	}
	wg.Wait()

	return data.Environments, nil
}

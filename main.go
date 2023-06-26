package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/anfern777/cicd-dashboard/controller"
	middleware "github.com/anfern777/cicd-dashboard/middleware"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := url.URL{
		User:     url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", os.Getenv("DB_SERVER"), os.Getenv("DB_PORT")),
		Path:     os.Getenv("DB_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database: Connection Successful")
	}

	router := gin.Default()
	router.Use(
		middleware.DatabaseConnect(db),
	)

	var (
		userService    service.UserService       = service.NewUserService()
		userController controller.UserController = controller.NewUserController(userService)

		schiService    service.SchiService       = service.NewSchiService()
		schiController controller.SchiController = controller.NewSchiController(schiService)

		cpiService    service.CpiService       = service.NewCpiService()
		cpiController controller.CpiController = controller.NewCpiController(cpiService)
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})

	router.GET("/auth", func(ctx *gin.Context) {
		email, password, _ := ctx.Request.BasicAuth()
		user, err := userService.FindByEmail(ctx, email)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = userService.CheckPassword(user, password)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		token, err := userService.JwtGenerateToken(user.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, token)
		}
	})

	users := router.Group("/user")
	users.Use(middleware.JwtValidate())
	{
		users.GET("/", func(ctx *gin.Context) {
			users, err := userController.FindAll(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, users)
			}
		})

		users.GET("/:email", func(ctx *gin.Context) {
			user, err := userController.FindByEmail(ctx, ctx.Param("email"))
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, user)
			}
		})

		users.POST("/", func(ctx *gin.Context) {
			err := userController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User submitted successfuly"})
			}
		})
	}

	schis := router.Group("/schi")
	schis.Use(middleware.JwtValidate())
	{
		schis.GET("/", func(ctx *gin.Context) {
			schis, err := schiController.FindAll(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, schis)
			}
		})

		schis.POST("/", func(ctx *gin.Context) {
			err := schiController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "SCHI submitted successfuly"})
			}
		})
	}

	cpis := router.Group("/cpi")
	cpis.Use(middleware.JwtValidate())
	{
		cpis.GET("/", func(ctx *gin.Context) {
			cpis, err := cpiController.FindAll(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, cpis)
			}
		})

		cpis.POST("/", func(ctx *gin.Context) {
			err := cpiController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "CPI submitted successfuly"})
			}
		})
	}

	// init app
	router.Run("0.0.0.0:8080")
}

package service

import (
	"errors"
	"os"
	"time"

	"github.com/anfern777/cicd-dashboard/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Save(*gin.Context, entity.User) error
	FindAll(*gin.Context) ([]entity.User, error)
	FindByEmail(*gin.Context, string) (entity.User, error)

	CheckPassword(entity.User, string) error
	HashPassword(string) (string, error)

	JwtGenerateToken(email string) (string, error)
	JwtValidateToken(tknStr string) (jwt.Claims, error)

	GetAllUserRoles() []string
}

type userService struct {
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewUserService() UserService {
	return &userService{}
}

func (service *userService) Save(ctx *gin.Context, user entity.User) error {
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	session := getDB(ctx).Session(&gorm.Session{FullSaveAssociations: true})
	return session.Save(&user).Error
}

func (service *userService) FindAll(ctx *gin.Context) ([]entity.User, error) {
	var users []entity.User
	err := getDB(ctx).Find(&users).Error
	return users, err
}

func (service *userService) FindByEmail(ctx *gin.Context, email string) (entity.User, error) {
	var user entity.User
	err := getDB(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func (service *userService) GetAllUserRoles() []string {
	return []string{"guest", "admin"}
}

func (service *userService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (service *userService) CheckPassword(user entity.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (service *userService) JwtGenerateToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}

func (service *userService) JwtValidateToken(tknStr string) (jwt.Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("could not parse token")
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

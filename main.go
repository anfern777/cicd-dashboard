package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/anfern777/cicd-dashboard/controller"
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect (dsn url.URL) gin.HandlerFunc {
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Println("Database: Connection Successful")
	}
  
	return func (c *gin.Context) {
	  c.Set ("DB", db)
	  c.Next ()
	}
  }

func main() {

	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Error loading .env file")
    }

	dsn := url.URL{
		User: url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")),
		Scheme: "postgres",
		Host: fmt.Sprintf("%s:%s", os.Getenv("DB_SERVER"), os.Getenv("DB_PORT")),
		Path: os.Getenv("DB_NAME"),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	router := gin.Default()
	router.Use(DatabaseConnect(dsn))


	// Auto migrate and update schemas
	// db.AutoMigrate(&User{}, &SourceCodeHostIntegration{}, &CloudProviderIntegration{})

	// db.Create(&User{Email: "test@email.com", Privilege: Guest, Password: "test123"})

	var(
		userService service.UserService = service.NewUserService()
		userController controller.UserController = controller.NewUserController(userService)

		schiService service.SchiService = service.NewSchiService()
		schiController controller.SchiController = controller.NewSchiController(schiService)

		cpiService service.CpiService = service.NewCpiService()
		cpiController controller.CpiController = controller.NewCpiController(cpiService)
	)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
	
	users := router.Group("/user")
	{
		users.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.FindAll(ctx))
		})

		users.GET("/:email", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,userController.FindByEmail(ctx, ctx.Param("email")))
		})

		users.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.Save(ctx))
		})
	}

	schis := router.Group("/schi")
	{
		schis.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, schiController.FindAll(ctx))
		})

		schis.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, schiController.Save(ctx))
		})
	}

	cpis := router.Group("/cpi")
	{
		cpis.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, cpiController.FindAll(ctx))
		})

		cpis.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, cpiController.Save(ctx))
		})
	}

	// init app
	router.Run("0.0.0.0:8080")
}

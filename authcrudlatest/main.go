package main

import (
	"authcrudlatest/config"
	"authcrudlatest/controllers"
	"authcrudlatest/middleware"
	"authcrudlatest/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.Data{})

	r := gin.Default()

	//r.Static("/static/*")
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/uploads", "./uploads") // Serve the uploads directory as static files

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", controllers.Index)

	r.GET("/logout", controllers.Logout)

	unauth := r.Group("/")
	unauth.Use(middleware.AlreadyAuthenticated())
	{
		unauth.GET("/login", controllers.LoginForm)
		unauth.POST("/login", controllers.Login)
		unauth.GET("/register", controllers.RegisterForm)
		unauth.POST("/register", controllers.Register)
	}

	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/dashboard", controllers.Dashboard)

		auth.GET("/data", controllers.CreateDataForm)
		auth.POST("/data", controllers.CreateData)
		auth.POST("/showdel/:id", controllers.DeleteData)

		auth.GET("/show", controllers.ReadData)

		auth.GET("/show/:id", controllers.UpdateDataById)
		auth.POST("/show/:id", controllers.UpdateDataByIdSave)
		//auth.PUT("/show/:id", controllers.UpdateData)

	}

	r.Run(":8080")
}

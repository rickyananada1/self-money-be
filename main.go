package main

import (
	"Self-Money_BE/controllers"
	"Self-Money_BE/database"
	_ "Self-Money_BE/docs"
	"Self-Money_BE/middlewares"
	"Self-Money_BE/models"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}
	database.GlobalDB.AutoMigrate(&models.User{}, &models.Pengeluaran{})
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			protected.GET("/profile", controllers.Profile)
		}
		pengeluaran := api.Group("/").Use(middlewares.Authz())
		{
			pengeluaran.POST("pengeluaran/create", controllers.CreatePengeluaran, middlewares.Authz())
			pengeluaran.PUT("pengeluaran/update", controllers.UpdatePengeluaran)
			pengeluaran.DELETE("pengeluaran/delete", controllers.DeletePengeluaran)
			pengeluaran.GET("pengeluaran", controllers.GetAllPengeluaran)
			pengeluaran.GET("pengeluaranByJenis", controllers.GetPengeluaranByJenisPengeluaran)
			pengeluaran.GET("pengeluaranByTanggal", controllers.GetPengeluaranByTanggal)
			pengeluaran.GET("pengeluaranById", controllers.GetPengeluaranById)
			pengeluaran.GET("GetPengeluaranFilterByJenisPengeluaran", controllers.GetPengeluaranFilterByJenisPengeluaran)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

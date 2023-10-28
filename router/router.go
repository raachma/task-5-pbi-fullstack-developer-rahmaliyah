package router

import (
	"github.com/cperdiansyah/golang-h8-dts/final-project/app"
	"github.com/cperdiansyah/golang-h8-dts/final-project/controller"
	"github.com/cperdiansyah/golang-h8-dts/final-project/middleware"
	"github.com/cperdiansyah/golang-h8-dts/final-project/repository"
	"github.com/cperdiansyah/golang-h8-dts/final-project/service"
	"github.com/gin-gonic/gin"
)

// func PhotoRouter(router *gin.Engine) {
// 	db := app.NewDB()

// 	repo := repository.NewPhotoRepository(db)
// 	srv := service.NewPhotoService(repo)
// 	ctrl := controller.NewPhotoController(srv)

// 	photoRouter := router.Group("/photos", middleware.AuthMiddleware())

//		{
//			photoRouter.POST("/", ctrl.Create)
//			photoRouter.GET("/", ctrl.GetAll)
//			{
//				photoRouter.PUT("/:id", middleware.PhotoMiddleware(srv), ctrl.Update)
//				photoRouter.DELETE("/:id", middleware.PhotoMiddleware(srv), ctrl.Delete)
//			}
//		}
//	}
func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := app.NewDB()

	repo := repository.NewPhotoRepository(db)
	srv := service.NewPhotoService(repo)
	ctrl := controller.NewPhotoController(srv)

	photoRouter := r.Group("/photos", middleware.AuthMiddleware())
	// router.Group("/photos", middleware.AuthMiddleware())
	{
		photoRouter.POST("/", ctrl.Create)
		photoRouter.GET("/", ctrl.GetAll)
		{
			photoRouter.PUT("/:id", middleware.PhotoMiddleware(srv), ctrl.Update)
			photoRouter.DELETE("/:id", middleware.PhotoMiddleware(srv), ctrl.Delete)
		}
	}

	// user := r.Group("/users")
	// {
	// 	user.POST("/register", controllers.RegisterUser)
	// 	user.POST("/login", controllers.LoginUser)
	// 	user.PUT("/:userId", controllers.UpdateUser)
	// 	user.DELETE("/:userId", controllers.DeleteUser)
	// }

	// photos := r.Group("/photos")
	// {
	// 	photos.POST("/", controllers.CreatePhoto)
	// 	photos.GET("/", controllers.GetPhotos)
	// 	photos.PUT("/:photoId", controllers.UpdatePhoto)
	// 	photos.DELETE("/:photoId", controllers.DeletePhoto)
	// }

	return r
}
func InitRouter() {
	r := SetupRouter()
	r.Run(":8080")
}

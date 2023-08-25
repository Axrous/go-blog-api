package app

import (
	"go-blog-api/controller"
	"go-blog-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	authController controller.AuthController, 
	userController controller.UserController, 
	postController controller.PostController, 
	commentController controller.CommentController) *httprouter.Router{

	router := httprouter.New()

	//Auth Controller
	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	//User Controller
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:userId", userController.FindById)

	//Post Controller
	router.POST("/api/posts", postController.Create)
	router.GET("/api/posts", postController.FindAll)
	router.GET("/api/posts/:postId", postController.FindById)
	router.PUT("/api/posts/:postId", postController.Update)
	router.DELETE("/api/posts/:postId", postController.Delete)

	//Comment Controller
	router.POST("/api/comments", commentController.Create)
	router.GET("/api/comments", commentController.FindByPostId)
	router.GET("/api/comments/:commentId", commentController.FindById)
	router.PUT("/api/comments/:commentId", commentController.Update)
	router.DELETE("/api/comments/:commentId", commentController.Delete)

	
	router.PanicHandler = exception.ErrorHandler
	
	return router
}
package router

import (
	"github.com/gin-gonic/gin"
	"sqlDblecture/controllers"
)
func HandleRequest(){
	rout := gin.Default()
	rout.GET("/getAllBooks",controllers.GetAllBooks)
	rout.GET("/getBooksByTitle",controllers.GetBooksByTitle)
 	rout.Run()
}
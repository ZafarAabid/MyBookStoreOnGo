package router

import (
	"github.com/gin-gonic/gin"
	"sqlDblecture/controllers"
)
func HandleRequest(){
	rout := gin.Default()
	rout.GET("/getAllBooks",controllers.GetAllBooks)
	rout.GET("/getBooksByTitle",controllers.GetBooksByTitle)
/*	OTHER REQUEST TYPES
	rout.POST("/getBooksByTitle",controllers.GetBooksByTitle)
	rout.PUT("/getBooksByTitle",controllers.GetBooksByTitle)
	rout.DELETE("/getBooksByTitle",controllers.GetBooksByTitle)
	rout.PATCH("/getBooksByTitle",controllers.GetBooksByTitle)
	rout.HEAD("/getBooksByTitle",controllers.GetBooksByTitle)
	rout.OPTIONS("/getBooksByTitle",controllers.GetBooksByTitle)
*/

	rout.Run()
}
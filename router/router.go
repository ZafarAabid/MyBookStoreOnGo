package router

import (
	"github.com/gin-gonic/gin"
	"sqlDblecture/controllers"
)
func HandleRequest(){
	rout := gin.Default()
	rout.GET("/Books",controllers.GetAllBooks)
	rout.GET("/BookByTitleAsPathVar",controllers.GetBooksByTitleByPathVar)
	rout.GET("BookByTitleAsParam/:title",controllers.GetBookByTitleByParam)
/*	OTHER REQUEST TYPES
	rout.POST("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
	rout.PUT("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
	rout.DELETE("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
	rout.PATCH("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
	rout.HEAD("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
	rout.OPTIONS("/getBooksByTitle",controllers.GetBooksByTitleByPathVar)
*/

	rout.Run()
}
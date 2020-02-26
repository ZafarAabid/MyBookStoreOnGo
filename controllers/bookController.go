package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sqlDblecture/config"
	"sqlDblecture/models"
)

func GetAllBooks(context *gin.Context){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println("################IF1")
		fmt.Println(err)

	}else{
		bookModel:=models.BookModel{Db:db}
		books,err := bookModel.FindAll()
		fmt.Printf("\nBooks  %v ,  %T\n",books,books)
		if err!=nil{
			fmt.Println("################IF2")
			fmt.Println(err)
		}else{
			context.JSON(200,books)
		}
	}
}

func GetBooksByTitle(ctx *gin.Context){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println("################GBBT1")
		fmt.Println(err)
	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.SearchBooksByTitle(ctx.Query("title"))
		if err!=nil{
			fmt.Println("################GBBT2")
			fmt.Println(err)
		}else{
			fmt.Printf("\nBooks  %v ,  %T\n",books,books)
			ctx.JSON(200,books)
		}
	}
}
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sqlDblecture/config"
	"sqlDblecture/models"
)

func GetAllBooks(w http.ResponseWriter,r *http.Request){
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
			json.NewEncoder(w).Encode(books)
		}
	}
}

func GetBooksByTitle(w http.ResponseWriter,r *http.Request){
	db,err:=config.GetMySQLDB()
	if err!=nil{
		fmt.Println("################GBBT")
		fmt.Println(err)
	}else{
		bookModel:=models.BookModel{Db:db}
		books,err:=bookModel.SearchBooksByTitle(r.URL.Query().Get("title"))
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Printf("\nBooks  %v ,  %T\n",books,books)
			json.NewEncoder(w).Encode(books)
		}
	}
}
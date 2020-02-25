package router
import (
	"log"
	"net/http"
	"sqlDblecture/controllers"
)
func HandleRequest(){
	http.HandleFunc("/getBooksByTitle",controllers.GetBooksByTitle)
	http.HandleFunc("/getAllBooks",controllers.GetAllBooks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
package main
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	Id string `json:"Id"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Id:"1", Title:"Test Title", Desc:"Test Description", Content:"Hello World"},
		Article{Id:"2",Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	articles := Articles{
		Article{Id:"1", Title:"Test Title", Desc:"Test Description", Content:"Hello World"},
		Article{Id:"2",Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	//fmt.Fprintf(w, "Key: "+ key)
	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func testPostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Test Post Endpoint Hit")
	fmt.Println("Endpoint POST: POST Articles Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Homepage Endpoint Hit")
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
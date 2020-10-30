package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"

	"log"
)


//we need a homepage func a function to handle requests and a main function to call the handlefunc
type article struct{
	Id 	  string `json:"Id"`
	Title string `json:"Title"`
	Desc string  `json:"Desc"`
	Content string `json:"Content"`
}

var articles []article


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to my web app")

	fmt.Println("endpoint hit: home page")

}
func returnArticles(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("endpoint hit: article page")
	json.NewEncoder(w).Encode(articles)
}
// func handleRequests(){
// 	http.HandleFunc("/",homePage)
// 	http.HandleFunc("/articles",returnArticles )
// 	log.Fatal(http.ListenAndServe(":10000", nil))
// }
func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticles)
	myRouter.HandleFunc("/article",createArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleArticles(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	key := vars["Id"]

	for _, article :=  range articles{
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}
	}

}




func createArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    articles = append(articles, article)

    json.NewEncoder(w).Encode(article)
	fmt.Println("Someone did a POST request")
}	

func deleteArticle(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]


	for index, article := range articles {
		if article.Id == id {

            articles = append(articles[:index], articles[index+1:]...)
        }
	}

}

func updateArticle(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var art article
	json.Unmarshal(reqBody, &art)
	
	for index, article :=  range articles{
		if article.Id == id{
			articles[index] = art
			
		}else {
			articles = append(articles, art)
			

		}
	}

	fmt.Println("Someone did a PUT request")

    


}
func main(){
	articles = []article{
		article{Id: "1",Title:"first", Desc: "some description", Content:"Article content"},
		article{Id: "2",Title: "second", Desc: "some description", Content: "article content"},

	}
	handleRequests() 
 
}
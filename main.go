package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
 
)
// Book Struct (Model)
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

//Init books variable as a slice Book struct

var books []Book


// Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get single book
func getBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) // Get params
	//Loop through books and find by id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) //MOCK ID - Random (not safe)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//Edit existing book
func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	//Loop through books and find by id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	//Loop through books and find by id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main(){
	// Init Mux Router
	r := mux.NewRouter()

	//mock data - todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "23499203", Title: "Sometimes a Great Notion", Author: &Author {Firstname: "Ken", Lastname: "Kesey"}})
	books = append(books, Book{ID: "2", Isbn: "90998398", Title: "One Flew Over The Cuckoo's Nest", Author: &Author {Firstname: "Ken", Lastname: "Kesey"}})
	books = append(books, Book{ID: "3", Isbn: "55434534", Title: "East of Eden", Author: &Author {Firstname: "John", Lastname: "Steinbeck"}})

	// Create Route Handlers / Endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

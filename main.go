// Simple GoLang CRUD App. Test using Postman
package main
  
import (
    "fmt"
    "log"
    "encoding/json"
    "math/rand"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type Book struct{
    ID string `json:"id"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

type Author struct{
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

var books []Book
  

func getBooks(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)

}

func deleteBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range books {

        if item.ID == params["id"]{
            books = append(books[:index],books[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r) 
    for _,item := range books{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            break
        }
    }
}

func createBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(100000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r) 
    
    for index, item := range books{
        if item.ID == params["id"]{
            books = append(books[:index], books[index+1:]...)
            var book Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = params["id"]
            books = append(books,book)
            json.NewEncoder(w).Encode(book)

        }
    }
}

// Main function
func main() {
    r := mux.NewRouter()

    books = append(books, Book{ID:"1", Title: "Book one", Author : &Author{Firstname:"John", Lastname:"Doe"}})
    books = append(books, Book{ID:"2", Title: "Book two", Author : &Author{Firstname:"Mark", Lastname:"Doe"}})

    // handlers
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
    
    fmt.Printf("Starting server at port 8000...\n")
    log.Fatal(http.ListenAndServe(":8000", r))


}
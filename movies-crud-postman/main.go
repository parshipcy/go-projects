package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie // Create a variable named movies that can hold many "Movie" values, but keep it empty for now.

func getMovies(w http.ResponseWriter, r *http.Request) {
	// This tells the client: “I am sending JSON data.”
	// Without this, the browser or client may not treat the response as JSON.
	// "Content-Type: application/json" tells the client to treat the response as JSON data.
	w.Header().Set("Content-Type", "application/json")

	// Takes the movies slice (your list of movies). Converts it into JSON format
	// Writes that JSON directly to w. Sending it back as the HTTP response body
	// You do not need fmt.Println or return here. Writing to w is enough.
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Extracts values from the URL. For /movies/1, it gives: params["id"] == "1"
	for index, item := range movies { //loop
		if item.ID == params["id"] {
			// movies[:index] → everything before the movie
 			// movies[index+1:] → everything after the movie
			// append joins them, skipping the deleted movie
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Extracts values from the URL. For /movies/1, it gives: params["id"] == "1"
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item) // NewEncoder → prepares a JSON encoder // Encode → converts Go data to JSON and sends it directly to w (w is where the output will go (HTTP response))
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie

	// Why _ =? -> Decode returns an error -> You are intentionally ignoring it for now
	_ = json.NewDecoder(r.Body).Decode(&movie) // Reads the request body (JSON sent by client) -> Converts JSON into a Go Movie struct -> Stores the result inside movie
	movie.ID = strconv.Itoa(rand.Intn(100000000)) //Generates a random number -> Converts it to a string ->Assigns it as the movie ID
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range //delete the movie with the id that u have sent
	//add a new movie - the movie that we send in the body of postman
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie //create an empty movie struct
			_ = json.NewDecoder(r.Body).Decode(&movie) //Read JSON from the request body and fill the struct. We are ignoring errors here (fine for learning).
			movie.ID = params["id"] //Ensure the ID stays the same as the URL.
			movies = append(movies, movie) //add the updated movie back to the list
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	// creates a new HTTP router using Gorilla Mux. r will receive incoming HTTP requests and decide which handler function should run.
	r := mux.NewRouter()

	// Adds a movie to the movies slice. append works even if movies is nil. Director: &Director{...} stores a pointer to a Director struct.
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie 1", Director: &Director{Firstname: "Parship", Lastname: "Chowdhury"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438445", Title: "Movie 2", Director: &Director{Firstname: "Suman", Lastname: "Das"}})

	// HandleFunc is a method on the router. It tells the server which function should run for a given URL.
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r)) // Start the HTTP server on port 8000, and if it fails, log the error and stop the program.
}


/*

mux.Vars only returns variables that appear inside {} in the route pattern.
If you define more variables, You will get more values:
r.HandleFunc("/movies/{id}/{year}", handler)

Request:
/movies/10/2024

Then:
params := mux.Vars(r)

Will return:
map[string]string{
  "id":   "10",
  "year": "2024",
}

You can access:
params["id"]
params["year"]

----------------------------------------------------------------

What HandleFunc expects:
func(w http.ResponseWriter, r *http.Request)

So your handler must look like this:
func getMovies(w http.ResponseWriter, r *http.Request) {
    // handle request here
}

----------------------------------------------------------------

func getMovies(w http.ResponseWriter, r *http.Request) {
This is an HTTP handler function.
It runs when someone hits the route mapped to getMovies.
w is used to send data back to the client.
r contains the incoming request means what user sent (URL, headers, method, etc.).

----------------------------------------------------------------

Encode converts data. NewEncoder(w) decides where the converted data goes. Let me explain:

Step 1: w (http.ResponseWriter)
w represents the HTTP response
Anything written to w is sent to the client
Think:
“w is the pipe to the browser/Postman.”

Step 2: json.NewEncoder(w)
Creates a JSON writer
Tells it:
“When you produce JSON, write it into w.”
At this point:
No conversion yet
Only setup

Step 3: .Encode(movie)
This is where conversion happens.
Encode:
Takes movie (Go struct)
Converts it into JSON text
Writes that JSON into w
Because the encoder was created with w

*/

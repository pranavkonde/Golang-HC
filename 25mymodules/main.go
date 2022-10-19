package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")
	// indirect in mod suggests library is not used
	// sum stores hashes which means to store the version of the library
	// go-mod-cache-download-github.com-(all packages are stored here)
	// anything coming from web is in slice of byte format
	// go mod tidy- package is directly been used(indirect to direct)
	// go mod verify - used to verify the sum file withthe packages been used
	// go installs - will list all packages installed
	// go list -m - will list packages for current directory
	// go mod why github.com/gorilla/mux - will display all the files using that package
	// go mod graph - will show graph for above if more than one package is used
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang</h1>"))
}

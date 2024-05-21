package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/courses", getCourses)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	fmt.Println("Servicio iniciado en el puerto: ")
	fmt.Println(8000)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /users")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /courses")
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

/*

	Primera forma de hacer peticiones mediante http

	func main() {
		port := ":3000"

		Se sustituye el uso de http por el router de gorilla mux
		http.HandleFunc("/users", getUsers)
		http.HandleFunc("/courses", getCourses)


		fmt.Println("Servicio iniciado en el puerto: ")
		fmt.Println(port)
		err := http.ListenAndServe(port, nil)

		if err != nil {
			fmt.Println("Error: ")
			fmt.Println(err)
		}
	}
	Se sustituye por el uso de gorilla mux
	func getUsers(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /users")
		io.WriteString(w, "This is my user endpoint")
	}

	func getCourses(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got /courses")
		io.WriteString(w, "This is my course endpoint")
	}
*/

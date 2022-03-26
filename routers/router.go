package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes main routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// Stubs for BarPad server
	router = InitBarPadRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		fmt.Println("Got connection...")
		w.Write([]byte("Welcome to Test API server"))
	})

	// router.HandleFunc("/api/v1/version", func(w http.ResponseWriter, rq *http.Request) {
	// 	fmt.Println("Got GET request...")
	// 	w.Write([]byte("{\"status\":\"ok\"}"))
	// })

	// router.HandleFunc("/api/v1/post", func(w http.ResponseWriter, rq *http.Request) {
	// 	fmt.Println("Got POST request...")
	// 	defer rq.Body.Close()
	// 	body, err := ioutil.ReadAll(rq.Body)
	// 	fmt.Println(string(body))

	// 	// fmt.Println("Sleeping for 500 milliseconds before responding...")
	// 	// time.Sleep(500 * time.Millisecond)
	// 	if err != nil {
	// 		http.Error(w, "Can't read POST body.", http.StatusBadRequest)
	// 		return
	// 	}
	// 	w.Write([]byte("{\"status\":\"ok\"}"))
	// })

	// router.HandleFunc("/api/v1/gettimeout", func(w http.ResponseWriter, rq *http.Request) {
	// 	fmt.Println("Got response, will sleep 30 seconds before responding...")
	// 	time.Sleep(30 * time.Second)
	// 	w.Write([]byte("{\"status\":\"timeout\"}"))
	// })

	return router
}

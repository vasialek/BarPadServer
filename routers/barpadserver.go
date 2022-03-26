package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vasialek/testserver/helpers"
	"github.com/vasialek/testserver/models"
	"github.com/vasialek/testserver/services"
)

var tableService = services.NewTableService()

// InitBarPadRoutes initalizes routes for BarPad server
func InitBarPadRoutes(r *mux.Router) *mux.Router {
	// Heartbeat
	r.HandleFunc("/api/v1/heartbeat", func(w http.ResponseWriter, rq *http.Request) {
		// defer rq.Body.Close()
		// data, err := ioutil.ReadAll(rq.Body)
		// if err != nil {
		// 	w.Write([]byte("No request data"))
		// 	return
		// }

		var heartbeatRq models.HeartbeatRequest
		err := helpers.DecodeAndDump(rq, &heartbeatRq)
		if err != nil {
			w.Write([]byte("Incorret JSON request"))
			return
		}
		// if err = json.Unmarshal(data, &heartbeatRq); err != nil {
		// 	w.Write([]byte("Incorret JSON request"))
		// 	return
		// }

		// tableService := services.NewTableService()
		resp := tableService.HandleHeartbeatRequest(&heartbeatRq)

		json, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte("Error serializing response on Heartbeat request."))
			return
		}
		w.Write(json)
	}).Methods("POST")

	// Request for waiter
	r.HandleFunc("/api/v1/request/waiter", func(w http.ResponseWriter, rq *http.Request) {
		defer rq.Body.Close()

		var waiterRq models.WaiterRequest
		if err := helpers.Decode(rq, &waiterRq); err != nil {
			w.Write([]byte("Incorrect JSON request"))
			return
		}

		resp := tableService.HandleWaiterRequest(&waiterRq)
		json, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte("Error serializing respons on Waiter request."))
			return
		}
		w.Write(json)
	}).Methods("POST")

	// Request for bill
	r.HandleFunc("/api/v1/request/bill", func(w http.ResponseWriter, rq *http.Request) {
		defer rq.Body.Close()
		dump(rq)
		w.Write([]byte("{\"Status\":true\"}"))
	}).Methods("POST")

	r.HandleFunc("/api/v1/status", func(w http.ResponseWriter, rq *http.Request) {
		defer rq.Body.Close()

		var statusRq models.StatusRequest
		if err := helpers.Decode(rq, &statusRq); err != nil {
			w.Write([]byte("Incorrect JSON request"))
			return
		}

		resp := tableService.HandleStatusRequest(&statusRq)

		json, err := json.Marshal(resp)
		if err != nil {
			w.Write([]byte("Error serializing Status response"))
			return
		}

		w.Header().Set("content-type", "application/json")
		w.Write(json)
	}).Methods("GET")

	return r
}

func dump(rq *http.Request) {
	fmt.Println(time.Now().Format("15:04:05"))
	body, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		fmt.Printf("Error parsing JSON request: %v\n", err)
		return
	}
	dumpba(body)
}

func dumpba(body []byte) {
	fmt.Println(time.Now().Format("15:04:05"))
	fmt.Println(string(body))
	fmt.Println("")
}

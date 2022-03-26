package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vasialek/testserver/helpers"
	"github.com/vasialek/testserver/models"
	"github.com/vasialek/testserver/repositories"
	"github.com/vasialek/testserver/routers"
)

func main() {
	fmt.Println("Starting test API server...")

	// url := "127.0.0.1:8079"
	url := "192.168.0.102:8079"
	server := http.Server{
		Addr:    url,
		Handler: routers.InitRoutes(),
	}

	err := server.ListenAndServe()
	// err := server.ListenAndServeTLS("certs\\dev-server.crt", "certs\\dev-server.key")
	if err != nil {
		log.Fatal(err)
	}
}

func testRepository() {
	repository := repositories.NewTableRepository()
	tables, err := repository.GetListOfTables()
	if err != nil {
		panic(err)
	}

	repository.UpdateTable("FakeTableId", &models.Table{
		LastHeartbeatAt: time.Now(),
	})

	helpers.DumpTables(tables)
}

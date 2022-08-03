package main

import (
	"assignment-3/handler"
	"assignment-3/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func writefile() {
	data := model.RandomValueStatus()
	file, _ := json.MarshalIndent(data, "", "")
	err := ioutil.WriteFile("file/status.json", file, 0644)
	if err != nil {
		fmt.Printf("%v \n", err)
	}
}

func main() {
	wd, _ := os.Getwd()
	//check is debug true
	if strings.Contains(wd, "cmd") {
		os.Chdir("..")
		os.Chdir("..")
	}
	ticker := time.NewTicker(14 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				writefile()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/assignment3", handler.MainHandler)

	fmt.Println("Now listening on port 127.0.0.1:8080")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

package main

import (
	"encoding/json"
	"os"
	"net/http"
	"fmt"
	"log"
	"io"
)

func start() {

	resp, err := http.Get(fmt.Sprintf("https://%v/v1/start/merge", os.Getenv("fixed_server")))
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rec map[string]string
	json.Unmarshal(body, &rec)	

	os.Setenv("server", rec["server"])
	os.Setenv("task", rec["task"])


	if _,b := os.LookupEnv("server"); b {
		fmt.Println("server present")
	}
	if _,b := os.LookupEnv("task"); b {
		fmt.Println("task present")
	}	
}
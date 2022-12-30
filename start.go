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
	
	var bearer string
	if t, ok := os.LookupEnv("token"); ok {
		bearer = "Bearer " + t
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%v/v1/start/%v", os.Getenv("fixed_server"), os.Getenv("tool")), nil)
	if err != nil {
		log.Fatal(err)
	}


	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)
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


	if elem, ok := rec["server"]; ok {
		os.Setenv("server", elem)
		fmt.Println("server present")
	}

	if elem, ok := rec["task"]; ok {
		os.Setenv("task", elem)
		fmt.Println("task present")
	}

}
package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"io"
)

func download(totalSize int64) {
	
	var bearer string
	if t, ok := os.LookupEnv("token"); ok {
		bearer = "Bearer " + t
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%v/v1/download/%v", os.Getenv("server"), os.Getenv("task")), nil)
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

	file, err := os.Create("merged.pdf")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(body)
	file.Close()
}
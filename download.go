package main

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"io"
)

func download(totalSize int64, numFiles int) {	
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

	var file *os.File
	switch(os.Getenv("tool")){
	case "merge":
		file, err = createPdf(os.Getenv("default_name"))
	case "split":
		if numFiles > 1 {
			file, err = createZip(os.Getenv("default_name"))
		} else {
			file, err = createPdf(os.Getenv("default_name"))
		}		
	}

	if err != nil {
		log.Fatal(err)
	}
	file.Write(body)
	file.Close()
}
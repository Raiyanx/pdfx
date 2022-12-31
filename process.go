package main

import (
	"encoding/json"
	"bytes"
	"os"
	"net/http"
	"fmt"
	"log"
	"io"
)

type entry map[string]string

func process(server_filenames map[string]string) {
		
	theMap := map[string]any{
		"task":os.Getenv("task"),
		"tool":os.Getenv("tool"),
	}

	entries := make([]entry, 0)

	for key, val := range server_filenames {
		fmt.Println(key, val)
		newEntry := entry{
			"server_filename":val,
			"filename":key,
		}
		entries = append(entries, newEntry)
	}

	theMap["files"] = entries


	postBody, _ := json.Marshal(theMap)

	responseBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%v/v1/process", os.Getenv("server")), responseBody)
	if err != nil {
		log.Fatal(err)
	}	

	var bearer string
	if t, ok := os.LookupEnv("token"); ok {
		bearer = "Bearer " + t
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()


	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rec map[string]string
	json.Unmarshal(body, &rec)	

	fmt.Println(rec)

	if elem, ok := rec["download_filename"]; ok {
		fmt.Println("download_filename present")
		fmt.Println(elem)
	}

}
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

func auth() {
	postBody, _ := json.Marshal(map[string]string{
		"public_key":os.Getenv("public_key"),
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(fmt.Sprintf("https://%v/v1/auth", os.Getenv("fixed_server")), "application/json", responseBody)
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


	if elem, ok := rec["token"]; ok {
		os.Setenv("token", elem)
		fmt.Println("token present")
	}
}
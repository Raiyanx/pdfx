package main

import (
	"encoding/json"
	"bytes"
	"os"
	"net/http"
	"fmt"
	"log"
	"io"
	"mime/multipart"
)

func upload(pdfs []string) (map[string]string, int64) {

	server_filenames := make(map[string]string)
	var totalSize int64 = 0

	for _, pdf := range pdfs {


		buf := new(bytes.Buffer)
		w := multipart.NewWriter(buf)


		if fd, e := os.Open(pdf); e != nil {
			log.Fatal(e)		
		} else {
			fStat, err := fd.Stat()
			if err != nil {
				log.Fatal(err)
			}
			size := fStat.Size()
			totalSize = totalSize + size
			fdata := make([]byte, size)	
			fd.Read(fdata)
			fd.Close()	
			part, err := w.CreateFormFile("file", pdf) 
			if err != nil {
				log.Fatal(err)
			}
			part.Write(fdata)
		}
		

		part, err := w.CreateFormField("task") 
		if err != nil {
			log.Fatal(err)
		}
		part.Write([]byte(os.Getenv("task")))

		w.Close()


		var bearer string
		if t, ok := os.LookupEnv("token"); ok {
			bearer = "Bearer " + t
		}
		
		req, err := http.NewRequest("POST", fmt.Sprintf("https://%v/v1/upload", os.Getenv("server")), buf)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", w.FormDataContentType())	
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

		if elem, ok := rec["server_filename"]; ok {
			fmt.Println("server_filename present")
			server_filenames[pdf] = elem
		}
	}
	return server_filenames, totalSize
}
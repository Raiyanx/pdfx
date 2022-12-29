package main

import (
	"bytes"
	"fmt"
	"os"
	"log"
	"net/http"
	"io"
	"encoding/json"
	"encoding/csv"
	"io/ioutil"
)

func printUsage() {
	usage := `Usage:
	pdfx merge [files]
	pdfx split filename [Integer/Integer-Integer]	
`
	fmt.Println(usage)
}

func setEnv() {
	fd, e := os.Open("env.csv")
	if e != nil {
		log.Fatal(e)
	}
	defer fd.Close()

	r := csv.NewReader(fd)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		os.Setenv(record[0], record[1])
	}
}





func main() {

	setEnv()
	fmt.Println(os.Getenv("public_key"))

	postBody, _ := json.Marshal(map[string]string{
		"public_key":os.Getenv("public_key"),
	})
	responseBody := bytes.NewBuffer(postBody)


	resp, err := http.Post(fmt.Sprintf("https://%v/v1/auth", os.Getenv("fixed_server")), "application/json", responseBody)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	fmt.Println(sb)


	if _,b := os.LookupEnv("token"); b {
		fmt.Println("token present")
	}

	printUsage()

}




















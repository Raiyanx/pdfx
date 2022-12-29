package main

import (
	"os"
	"encoding/csv"
	"log"
	"io"
)

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
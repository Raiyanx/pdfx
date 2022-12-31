package main

import (
	"os"
)

func main() {

	pdfs := []string{"q1.pdf","q2.pdf"}
	os.Setenv("tool", "merge")

	setEnv()
	auth()
	start()	
	server_filenames, totalSize := upload(pdfs)
	process(server_filenames)
	download(totalSize)
}




















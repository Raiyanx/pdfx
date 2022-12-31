package main

import (
	"os"
	"fmt"
	"flag"
)

func main() {
	setEnv()

	dn := flag.String("dn", "newfile.pdf", "Default name used for new files")
	flag.Parse()
	os.Setenv("default_name", *dn)

	args := make([]string, 0)
	for _, s := range os.Args[1:] {
		if s[0] != '-' {
			args = append(args, s)
		}
	}

	pdfs := make([]string, 0)
	
	if len(args) == 0 {
		printUsage()
		return
	} else if args[0] == "merge" {
		if len(args) < 3 {
			fmt.Println("Please select at least two files for merging")
		} else {
			for i:=1; i<len(args); i++ {
				pdfs = append(pdfs, args[i])
			}
			os.Setenv("tool", "merge")
			startProcess(pdfs)
		}
	} else if args[0] == "split" {
		fmt.Println("using split")
	} else {
		printUsage()
		return
	}
}


func startProcess(pdfs []string) {
	auth()
	start()	
	server_filenames, totalSize := upload(pdfs)
	process(server_filenames)
	download(totalSize)
}





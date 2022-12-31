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
	ranges := make([]string, 0)

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
			startProcess(pdfs, "merge", []string{})
		}
	} else if args[0] == "split" {
		if len(args) < 2 {
			fmt.Println("Please select a file for splitting")
		} else if len(args) < 3 {
			fmt.Println("Please specify at least one range")
		} else {
			pdfs = append(pdfs, args[1])
			for i:=2; i<len(args); i++ {
				ranges = append(ranges, args[i])
			}
			os.Setenv("tool", "split")
			startProcess(pdfs, "split", ranges)
		}
	} else {
		printUsage()	
		return
	}
}


func startProcess(pdfs []string, tool string, ranges []string) {
	auth()
	start()	
	server_filenames, totalSize := upload(pdfs)
	switch(tool){
	case "merge":
		merge(server_filenames)
	case "split":
		split(server_filenames, ranges)
	default:
		printUsage()	
	}
	download(totalSize, len(ranges))
}





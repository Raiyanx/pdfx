package main

import (
	"fmt"
	"os"
	"github.com/raiyanx/pdfx/merge"
	"github.com/raiyanx/pdfx/split"
)

func printUsage() {
	usage := `Usage:
	pdfx merge [files]
	pdfx split filename [Integer/Integer-Integer]	
`
	fmt.Println(usage)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
	} else if args[0] == "merge" {
		merge.MergePrint()
	} else if args[0] == "split" {
		split.SplitPrint()
	} else {
		printUsage()
	}
}

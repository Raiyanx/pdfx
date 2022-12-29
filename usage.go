package main

import (
	"fmt"
)

func printUsage() {
	usage := `Usage:
	pdfx merge [files]
	pdfx split filename [Integer/Integer-Integer]	
`
	fmt.Println(usage)
}
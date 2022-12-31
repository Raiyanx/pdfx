package main

import (
	"fmt"
)

func printUsage() {
	usage := `Usage:
	pdfx [-flags] merge [files]
	pdfx [-flags] split filename [Integer/Integer-Integer]	
`
	fmt.Println(usage)
}
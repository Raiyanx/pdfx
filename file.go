package main

import (
	"os"
	"strconv"
)

func createFile(s string) (*os.File, error) {
	if len(s) >= 4 {
		if s[len(s)-4:len(s)] == ".pdf" {
			s = s[:len(s)-4]
		}	
	}
	if fd, err := os.Open(s + ".pdf"); err != nil {  
		fd.Close()
		return os.Create(s + ".pdf")
	} else {
		var i int64 = 1
		for {
			newname := s + "_" + strconv.FormatInt(i, 10) + ".pdf"
			if fd, err := os.Open(newname); err != nil {  
				fd.Close()
				return os.Create(newname)
			}
			i++
		}
	} 
}
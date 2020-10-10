package main

import (
	"fmt"
)

func main() {
	value := 5
	avgVal := (value + 1) / 2

	for i := 0; i < value; i++ {
		print := ""
		for j := 0; j < value; j++ {
			if j == 0 || j == (value - 1) || i == (avgVal - 1) {
				if j < (value - 1) {
					print += "* "
				} else {
					print += "*"
				}	
			} else {
				if j < (value - 1) {
					print += "- "
				} else {
					print += "-"
				}
			}
		}
		fmt.Println(print)
	}
}
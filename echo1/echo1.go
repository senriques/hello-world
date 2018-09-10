package main

// Shawn Enriques
// Assignment 1.3 echo command line arguments.

import (
	"fmt"
	"os"
)

func main() 	{
	ags:= os.Args[1:]
	for ii, ag := range ags	{
		if ii < len(ags)	{
			fmt.Printf("%s ", ag)
		} else {
			fmt.Printf("%s", ag)
		}
	}
	fmt.Printf("\n")
}

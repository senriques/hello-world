package main

// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.6 Test a simple library package

import	(
  "fmt"
  "c:/Go/src/mkPkg/test1"	// import test1
)

func main() {
  out := test1.DoubleValue(8)	// Call function
  fmt.Printf("out = %d\n", out)	// should return 16
  }
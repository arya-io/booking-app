// Package main is the starting point of an executable Go program.
// 'main' is a special package name that indicates this is a runnable program.
package main

// Importing the 'fmt' package, which provides functions for formatted I/O, including printing to the console.
import "fmt"

// The main function is the entry point of every Go program.
// Execution of the program begins here.
func main() {

// Println outputs the string and appends a newline automatically.
fmt.Println("Hello World")

// Print outputs the string without adding a newline at the end.
fmt.Print("Hello World")
}
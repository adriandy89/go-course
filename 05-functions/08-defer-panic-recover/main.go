package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("OCURRIO UN ERROR!", error)
		}
	}()

	// Check if the data file exists
	if fi, err := os.Stat("data2.txt"); os.IsNotExist(err) || fi.IsDir() {
		panic("Error: data.txt file not found.")
	}
	// Open the data file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("closed data.txt")
	}()
	// Read the contents of the file
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Print the contents of the file
	fmt.Println(string(buf[:n]))
}

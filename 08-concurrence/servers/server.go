package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	begin := time.Now()

	chanel := make(chan string)
	// var chanel chan string

	servers := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.wikipedia.net",
	}

	// for _, server := range servers {
	// 	checkServer(server)
	// }

	for _, server := range servers {
		go checkServer(server, chanel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-chanel)
	}

	end := time.Since(begin)
	fmt.Println(end)

}

func checkServer(server string, chanel chan string) {
	_, err := http.Get(server)

	if err != nil {
		chanel <- server + " No disponible"
	} else {
		chanel <- server + " Disponible"
	}
}

// func checkServer(server string) {
// 	_, err := http.Get(server)

// 	if err != nil {
// 		fmt.Println(server, " No disponible")
// 	} else {
// 		fmt.Println(server, " Disponible")
// 	}
// }

package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"log"
	"flag"
	"time"
)

func main() {

	server := flag.String("server", "http://localhost:4000", "http server to test")
	rate := flag.Float64("rate", 1, "requests per second")

	flag.Parse()

	fmt.Println("Server ", *server)
	fmt.Println("Rate", *rate)

	//var interval := float32()
	interval := 1 / *rate
	fmt.Println("Interval", interval)
	sleeptime := interval * 1000
	fmt.Println("sleep time", sleeptime)
	var succeeded int

	for start:= time.Now(); time.Since(start) < 15 * time.Second; {
		resp, err := http.Get("http://localhost:4000/")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		s := string(body)
		if s == "OK" {
			succeeded += 1
		}
		fmt.Println(s)
		time.Sleep(time.Duration(sleeptime) *  time.Millisecond)
		//time.Sleep(900 *  time.Millisecond)
		//time.Sleep(500 *  time.Millisecond)
	}
	// ...
	fmt.Println("Successful requests:", succeeded)
}
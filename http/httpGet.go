package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("https://sufinama.org/ebooks/?subcategory=epics")
	if err != nil {

		fmt.Printf("epics")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

	}
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func String(data []byte) {
	panic("unimplemented")
}

package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	response, _ := http.Get("http://bencrevis.uk/")
	page, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s", page)
}
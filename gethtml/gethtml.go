package gethtml

import (
	"golang.org/x/net/html"
	"net/http"
	"fmt"
)

func Gethtmlcode(s string) (*html.Node){
	resp, err := http.Get(s)
	if err != nil{
		fmt.Println("Unable to GET the domain  ",err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil{
		fmt.Println("Error while parsing the response  ",err)
	}

	return doc
}
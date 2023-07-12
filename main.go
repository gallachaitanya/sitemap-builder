package main

import (
	"encoding/xml"
	"exercise5/gethtml"
	"exercise5/links"
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/maps"
)

func main(){
	domain := flag.String("d","https://www.calhoun.io","Url that need to be parsed for site builder")
	flag.Parse()

	doc := gethtml.Gethtmlcode(*domain)

	m := make(map[string]bool)
	maps.Copy(m,links.FindLinks(doc, m, *domain))

	type loc struct{
		Val string `xml:"loc"`
	}
	var xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
	type urlset struct{
		Urls []loc `xml:"url"`
		Xmlns string `xml:"xmlns,attr"`
	}

	toXml := urlset{
		Xmlns: xmlns,
	}

	for k := range links.Parselink(m,*domain){
		toXml.Urls = append(toXml.Urls, loc{k})
	}

	fmt.Print(xml.Header)
	enc:= xml.NewEncoder(os.Stdout)
	enc.Indent("","  ")
	if err := enc.Encode(toXml); err != nil{
		fmt.Println("Error in xml encoding..",err)
	}
	fmt.Println()


}
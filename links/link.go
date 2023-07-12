package links

import (
	"strings"

	"golang.org/x/net/html"
)

func FindLinks(n *html.Node, m map[string]bool, domain string) map[string]bool{

	if n.Type == html.ElementNode && n.Data == "a"{
			
		for _, a := range n.Attr{
			if a.Key == "href"{
				if _, ok := m[a.Val]; !ok{
					if strings.Contains(a.Val,"http"){
						if strings.Contains(a.Val,domain){
							m[a.Val] = false
						}	
						}else{
							if _,ok :=m[domain + a.Val]; !ok{
							m[domain + a.Val] = false
							}
						}
					
				}
			}
		}
		
	}
	for c := n.FirstChild; c != nil; c= c.NextSibling {
		FindLinks(c,m,domain)
	}

	return m
}
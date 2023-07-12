package links

import (
	"exercise5/gethtml"

	"golang.org/x/exp/maps"
)

func Parselink(m map[string]bool, domain string) map[string]bool {
	var temp map[string]bool = m
	for k, v := range m {   
		if !v {
				tmpnode := gethtml.Gethtmlcode(k)
			    tmpmap := FindLinks(tmpnode,temp,domain)
				maps.Copy(temp,tmpmap)
				temp[k] = true
			}
			
			}	
	for _, v := range temp{
		if !v{
			
			Parselink(temp,domain)
		}
	}
	return temp
}
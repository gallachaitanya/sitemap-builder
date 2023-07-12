package links

import (
	"exercise5/gethtml"
	"fmt"

	"golang.org/x/exp/maps"
)

func Parselink1(m map[string]bool, domain string) map[string]bool {
	var temp map[string]bool = m
	ch := make(chan map[string]bool)
	ch <- temp
	for k, v := range m {   
		if !v {
			valu := k
			go func(ch chan map[string]bool){
				sampmap := <- ch
				tmpnode := gethtml.Gethtmlcode(valu)
			    tmpmap := FindLinks(tmpnode,sampmap,domain)
				maps.Copy(sampmap,tmpmap)
				sampmap[valu] = true
				ch <- sampmap
			}(ch)
			}
			}
	// for v:= range ch{
	// 	maps.Copy(temp,v)
	// }
	temp = <- ch	
	for _, v := range temp{
		if !v{
			fmt.Println("entered...")
			Parselink(temp,domain)
		}
	}
	return temp
}
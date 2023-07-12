package test

import(
	"testing"
	"exercise5/gethtml"
	"exercise5/links"
	"golang.org/x/exp/maps"
)

func BenchmarkParse(b *testing.B) {
	doc := gethtml.Gethtmlcode("https://www.calhoun.io")

	m := make(map[string]bool)
	maps.Copy(m,links.FindLinks(doc, m, "https://www.calhoun.io"))

	for n := 0; n < b.N; n++ {
		links.Parselink(m, "https://www.calhoun.io")
}

	 
}
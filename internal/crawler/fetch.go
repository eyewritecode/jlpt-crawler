package crawler

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type GrammarItem struct {
	Word       string
	DetailLink string
}

func DownloadGrammarCard(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download flash cards")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to download flash cards")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	doc.Find("#jl-grammar tbody tr.jl-row").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td").Eq(1).Text()
		wordDetailUrl, exists := s.Find("td").Eq(1).Find("a").Attr("href")
		if !exists {
			fmt.Errorf("word detail page missing")
		}
		fmt.Printf("%s: %s\n", title, wordDetailUrl)
	})
	return "", nil
}

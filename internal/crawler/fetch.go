package crawler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GrammarItem struct {
	Word       string
	DetailLink string
}

func FetchAllGrammarItems(startURL string) ([]GrammarItem, error) {
	var allItems []GrammarItem

	visited := make(map[string]bool)
	nextURL := startURL

	for nextURL != "" {
		if visited[nextURL] {
			break
		}
		visited[nextURL] = true

		items, newNextURL, err := ParseGrammarPage(nextURL)
		if err != nil {
			return allItems, err
		}
		allItems = append(allItems, items...)
		nextURL = newNextURL
	}
	return allItems, nil
}

func ParseGrammarPage(url string) ([]GrammarItem, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch page: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, "", fmt.Errorf("bad status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	var items []GrammarItem
	var nextURL string

	doc.Find("#jl-grammar tbody tr").Each(func(i int, s *goquery.Selection) {
		word := strings.TrimSpace(s.Find("td").Eq(2).Text())
		detailURL, exists := s.Find("td").Eq(1).Find("a").Attr("href")
		if exists && word != "" {
			items = append(items, GrammarItem{
				Word:       word,
				DetailLink: detailURL,
			})
		}
	})

	doc.Find(".pagination .page-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		link := s.Find("a")
		if link.Length() == 0 {
			return true
		}
		text := strings.TrimSpace(link.Text())
		if text == "→" {
			href, exists := link.Attr("href")
			if exists && href != "" {
				nextURL = href
				return false
			}
		}
		return true
	})
	return items, nextURL, nil
}

func DownloadGrammarCard(detailURL, word, destDir string) error {
	res, err := http.Get(detailURL)
	if err != nil {
		return fmt.Errorf("failed to fetch detail page: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("bad status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("failed to parse detail HTML: %w", err)
	}

	imgSrc, exists := doc.Find(".grammar-thumbnail-cont a").Attr("href")
	if !exists || imgSrc == "" {
		return fmt.Errorf("image not found for %s", word)
	}

	imgRes, err := http.Get(imgSrc)
	if err != nil {
		return fmt.Errorf("failed to download image: %w", err)
	}
	defer imgRes.Body.Close()
	if imgRes.StatusCode != 200 {
		return fmt.Errorf("bad status fetching image: %d", imgRes.StatusCode)
	}

	fileName := fmt.Sprintf("%s.jpg", word)
	filePath := filepath.Join(destDir, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, imgRes.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	return nil
}

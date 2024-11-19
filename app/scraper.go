package app

import (
	"path/filepath"
	"strings"

	colly "github.com/gocolly/colly"
)

func isValidImage(url string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(url))
	for _, v := range validExtensions {
		if ext == v {
			return true
		}
	}
	return false
}

func Grabber(url string) map[int]string {
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(5),
	)
	photos := make(map[int]string)

	c.OnHTML("img", func(e *colly.HTMLElement) {

		imgSrc := e.Attr("src")
		imgSrc = e.Request.AbsoluteURL(imgSrc)

		if isValidImage(imgSrc) {
			photos[len(photos)] = imgSrc
		}

	})

	c.Visit(url)
	c.Wait()
	return photos
}

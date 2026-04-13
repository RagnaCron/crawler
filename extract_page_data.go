package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	var data PageData
	u, err := url.Parse(pageURL)
	if err != nil {
		log.Fatalln(err)
	}

	urls, err := getURLsFromHTML(html, u)
	if err != nil {
		log.Fatalln(err)
	}

	imageUrls, err := getImagesFromHTML(html, u)
	if err != nil {
		log.Fatalln(err)
	}

	data = PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  urls,
		ImageURLs:      imageUrls,
	}

	return data
}

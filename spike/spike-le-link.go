package spike

import (
  "fmt"
  "log"

  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
  doc, err := goquery.NewDocument("https://moz.com/blog/meta-data-templates-123")
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("meta").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
	  metatagProperty, _ := s.Attr("property")
	  metatagContent, _ := s.Attr("content")
    fmt.Printf("Metatag %s: %s\n", metatagProperty, metatagContent)
  })
}

func main() {
  ExampleScrape()
}

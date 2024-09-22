package crawler_test

import (
	"fmt"
	"gosearch/pkg/crawler"
	"testing"
)

func TestDocumentSerialize(t *testing.T) {
	want := []byte{91, 123, 34, 73, 68, 34, 58, 49, 50, 51, 44, 34, 85, 82, 76, 34, 58, 34, 103, 111, 34, 44, 34, 84, 105, 116, 108, 101, 34, 58, 34, 84, 104, 101, 32, 71, 111, 34, 125, 93}
	docs := []crawler.Document{
		{
			ID:    123,
			URL:   "go",
			Title: "The Go",
		},
	}
	got := crawler.DocumentSerialize(&docs)
	for i := range want {
		if got[i] != want[i] {
			t.Fail()
		}
	}
}

func TestDocumentDeSerialize(t *testing.T) {
	docs := []byte{91, 123, 34, 73, 68, 34, 58, 49, 50, 51, 44, 34, 85, 82, 76, 34, 58, 34, 103, 111, 34, 44, 34, 84, 105, 116, 108, 101, 34, 58, 34, 84, 104, 101, 32, 71, 111, 34, 125, 93}
	want := []crawler.Document{
		{
			ID:    123,
			URL:   "go",
			Title: "The Go",
		},
	}
	got := crawler.DocumentDeSerialize(docs)
	fmt.Println(got)
	for i := range want {
		if got[i] != want[i] {
			t.Fail()
		}
	}
}

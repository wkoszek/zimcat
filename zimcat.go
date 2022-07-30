package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	zim "github.com/akhenakh/gozim"

	"jaytaylor.com/html2text"
)

var (
	z *zim.ZimReader
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		print("usage!")
		os.Exit(1)
	}
	path := flag.Args()[0]

	z, err := zim.NewReader(path, false)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(.*)This article is issued from Wikipedia.*Additional terms may apply for the media files.(.*)`)

	z.ListTitlesPtrIterator(func(idx uint32) {
		a, err := z.ArticleAtURLIdx(idx)
		if err != nil || a.EntryType == zim.DeletedEntry {
			return
		}

		if a.Namespace == 'A' {
			htmldata, err := a.Data()
			if err != nil {
				log.Fatal(err.Error())
			}
			htmlstrraw := string(htmldata)
			if len(htmlstrraw) <= 0 {
				return
			}
			htmlstr := htmlstrraw

			text, err := html2text.FromString(htmlstr, html2text.Options{PrettyTables: false})
			if err != nil {
				log.Fatal(err)
			}

			if len(text) <= 0 {
				return
			}

			text = a.Title + "\n\n" + re.ReplaceAllString(text, "$1$2")

			fmt.Print(text)
		}
	})
}

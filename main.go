package main

import (
	"flag"
	"log"
	"os"

	"github.com/Gophers254/goWebCrawler/crawler"
)

func main() {
	// get file name and line number when the code crashes
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Println("WARNING | Please enter a valid URI to crawl")
		os.Exit(1)
	}

	crawler.Crawl(args[0])

}

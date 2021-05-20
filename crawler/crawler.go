package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Crawl public function which does the actual crawling
// and fetches different quotes from the target URI
func Crawl(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("WARNING | The following error has occured while fetching contents from URI: %s. Error: %v\n", uri, err.Error())
	}

	// close the openned resource before exiting the Crawl function
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

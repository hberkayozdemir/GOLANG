/* package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://ytdfinans.web.app",
}

func fetch(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}

func crawl() {
	for _, url := range urls {
		go fetch(url)
	}
}

func main() {
	crawl()

}
bu şekilde herhangi bir değer dönmedi birden çok url var hepsini anlık alamadı bunun için wait group kullanuyoruz

*/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://ytdfinans.web.app",
}

func fetch(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	wg.Done()
	// her işi bitirdiğinde burda bitsin yenisin geçsin hesabıı

}

func crawl() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("Crawling Urls finished")

}

func main() {
	crawl()

}

package workers

import (
	"fmt"
	"math/rand"
	"time"
)

func getPage(url string) (int, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3000) 
	time.Sleep(time.Duration(n) * time.Millisecond)
	return len(url), nil
}
func worker(urlCh chan string, sizeCh chan string, id int) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("%s has length %d, id: (%d)", url, length, id)
		} else {
			sizeCh <- fmt.Sprintf("Error getting %s: %s, id:(%d)", url, err, id)
		}
	}
}
func generator(url string, urlCh chan string) {
	urlCh <- url
}

func Initialize() {
	urls := []string{
		"http://www.google.com.com/",
		"http://www.facebook.com.com/",
		"http://www.twitter.com.com/",
		"http://www.instagram.com.com/",
	}

	sizeCh := make(chan string)
	urlCh := make(chan string)

	for id := 0; id < 10; id++ {
		go worker(urlCh, sizeCh, id)
	}

	for _, url := range urls {
		go generator(url, urlCh)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

}

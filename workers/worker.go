package workers

import "fmt"


func getPage(url string) (int, error){
  return len(url), nil
}
func worker(urlCh chan string, sizeCh chan string, id int) {
  for{
    url := <- urlCh
    length, err := getPage(url)
    if err == nil {
      sizeCh <- fmt.Sprintf("%s has length %d, id: (%d)", url, length, id)
    } else {
      sizeCh <- fmt.Sprintf("Error getting %s: %s, id:(%d)", url, err, id)
    }
  }
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
    go worker(urlCh,sizeCh, id)
  }

  for _,url := range urls {
    urlCh <- url  
  }
  for i:= 0; i < len(urls); i++ {
    fmt.Printf("%s\n", <-sizeCh)
  }
  
  
}

//waitgroup用于等待一组线程的结束

package main

import (
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		//父线程调用Add方法来设定应等待的线程的数量
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}

	//wait方法一直阻塞直到waitgroup计数器减到0
	wg.Wait()
}

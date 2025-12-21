package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func fetchLength(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("status não OK para %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.bing.com",
	}

	type result struct {
		url    string
		length int
		err    error
	}

	results := make(chan result, len(urls))
	var wg sync.WaitGroup

	start := time.Now()
	for _, url := range urls {
		wg.Add(1)
		// lança goroutine para cada requisição
		go func(u string) {
			defer wg.Done()
			length, err := fetchLength(u)
			results <- result{url: u, length: length, err: err}
		}(url)
	}

	// fecha o canal quando todas as goroutines finalizam
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.err != nil {
			fmt.Printf("erro ao buscar %s: %v\n", res.url, res.err)
			continue
		}
		fmt.Printf("%s length: %d\n", res.url, res.length)
	}
	fmt.Printf("elapsed: %dms\n", time.Since(start).Milliseconds())
}

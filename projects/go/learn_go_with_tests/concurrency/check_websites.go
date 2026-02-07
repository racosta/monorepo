// Package concurrency provides functions for checking the availability of websites concurrently.
package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns a map
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

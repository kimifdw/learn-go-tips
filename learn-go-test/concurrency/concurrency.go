package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results

}

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

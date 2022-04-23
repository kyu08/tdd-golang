package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

  for _, u := range urls {
    results[u] := wc(u)
  }

  return results
}
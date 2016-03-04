package concurrency

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	flag.Parse()

	sites := flag.Args()
	start := time.Now()

	WithWG(sites)
	// WithoutWG(sites)
	fmt.Printf("Entire process took %s\n", time.Since(start))
}

// WithWG runs the fetcher with wait groups
func WithWG(sites []string) {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(len(sites))

	for _, site := range sites {
		go func(site string) {
			defer wg.Done()
			begin := time.Now()
			if _, err := http.Get(site); err != nil {
				fmt.Println(site, err)
				return
			}
			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}(site)

	}
	wg.Wait()
	fmt.Printf("Entire process took %s\n", time.Since(start))
}

// WithoutWG runs the fetcher without wait groups or concurrency
func WithoutWG(sites []string) {
	start := time.Now()
	for _, site := range sites {
		fmt.Println(site)
		begin := time.Now()
		if _, err := http.Get(site); err != nil {
			fmt.Println(site, err)
			return
		}
		fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
	}
	fmt.Printf("Entire process took %s\n", time.Since(start))
}

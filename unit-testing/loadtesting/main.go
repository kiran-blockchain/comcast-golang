package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "sync"
    "time"
    "golang.org/x/sync/errgroup"
)

const (
    numRequests      = 10000
    concurrentUsers  = 100
    requestURL       = "http://localhost:4000"
)

func main() {
    var g errgroup.Group
    var mu sync.Mutex

    successes := 0
    failures := 0

    start := time.Now()
    
    // Create a buffered channel to limit concurrency
    sem := make(chan struct{}, concurrentUsers)

    for i := 0; i < numRequests; i++ {
        sem <- struct{}{}
        g.Go(func() error {
            defer func() { <-sem }()

            resp, err := http.Get(requestURL)
            if err != nil {
                mu.Lock()
                failures++
                mu.Unlock()
                return err
            }
            defer resp.Body.Close()
            
            _, err = ioutil.ReadAll(resp.Body)
            if err != nil {
                mu.Lock()
                failures++
                mu.Unlock()
                return err
            }

            if resp.StatusCode == http.StatusOK {
                mu.Lock()
                successes++
                mu.Unlock()
            } else {
                mu.Lock()
                failures++
                mu.Unlock()
            }

            return nil
        })
    }

    // Wait for all requests to finish
    if err := g.Wait(); err != nil {
        log.Fatalf("Error occurred during load test: %v", err)
    }

    duration := time.Since(start)

    fmt.Printf("Completed %d requests in %v\n", numRequests, duration)
    fmt.Printf("Successes: %d\n", successes)
    fmt.Printf("Failures: %d\n", failures)
}

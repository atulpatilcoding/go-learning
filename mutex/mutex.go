package main

import (
"fmt"
"sync"
)

type Post struct {
views int
mu    sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
defer wg.Done()

p.mu.Lock()
p.views++
p.mu.Unlock()
}

func main() {
myPost := Post{views: 0}
var wg sync.WaitGroup

for i := 0; i < 100; i++ {
wg.Add(1)
go myPost.inc(&wg)
}

wg.Wait()
fmt.Println(myPost.views)
}
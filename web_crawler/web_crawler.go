package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	timeOut = 5 * time.Second
)

type web_crawler struct {
	bq                    BlockingQueue
	longLivedWorkersCount int
	startNode             string
	nodes                 map[string][]string
	visited               map[string]bool
	visitedMutex          sync.Mutex
	totalUrlsVisited      int
}

func newWebCrawler() *web_crawler {

	wc := &web_crawler{}
	wc.nodes = make(map[string][]string)
	wc.nodes["google"] = []string{"facebook", "linkedin", "bookmyshow"}
	wc.nodes["facebook"] = []string{"instagram", "whatsapp", "snapchat"}
	wc.nodes["linkedin"] = []string{"naukri", "instahyre", "coderoundai"}
	wc.nodes["instagram"] = []string{"google"}
	wc.nodes["naukri"] = []string{"google"}
	wc.longLivedWorkersCount = 3
	wc.bq = NewBlockingQueue(3)
	wc.startNode = "google"
	wc.visited = make(map[string]bool)
	wc.visitedMutex = sync.Mutex{}
	wc.totalUrlsVisited = 0

	return wc
}

func (wc *web_crawler) run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	wc.visited[wc.startNode] = true
	ok := wc.bq.Offer(wc.startNode, timeOut)
	if !ok {
		panic("Unable to push start element into queue")
	}

	for i := 0; i < wc.longLivedWorkersCount; i++ {
		g.Go(func() error {
			return wc.executeBFS()
		})
	}

	if err := g.Wait(); err != nil {
		panic("BFS Failed")
	}

}

// worker functions
func (wc *web_crawler) executeBFS() error {

	for {

		ele, ok := wc.bq.Poll(timeOut)
		if !ok {
			//Queue is empty return
			return nil
		}

		currentNode := ele.(string)
		fmt.Println("current node", currentNode)

		neighbors := wc.nodes[currentNode]

		for _, value := range neighbors {
			wc.visitedMutex.Lock()
			if wc.visited[value] == false {
				wc.visited[value] = true
				wc.totalUrlsVisited++
				wc.bq.Offer(value, timeOut)
			}
			wc.visitedMutex.Unlock()
		}

	}
}

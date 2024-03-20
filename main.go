package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"bufio"
	"github.com/adhocore/chin"
)

func main() {
	var dumpPath, query string
	var wg sync.WaitGroup
	var inputText string
	fmt.Println("✨ GO-FTS Engine running ✨")
	
	fmt.Println("🚀 Enter your text to be searched: ")
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", inputText, "search query")
	flag.Parse()
	s := chin.New().WithWait(&wg)

	start := time.Now()

	go s.Start()

	fmt.Println("✨ Loading documents ✨")
	wg.Add(1)
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()

	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	fmt.Println("✨ Indexing documents ✨")
	start = time.Now()

	wg.Add(1)
	idx := make(index)
	idx.add(docs)
	wg.Done()

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	fmt.Println("✨ Searching documents ✨")

	wg.Add(1)
	matchedIDs := idx.search(query)
	wg.Done()

	s.Stop()
	wg.Wait()

	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}

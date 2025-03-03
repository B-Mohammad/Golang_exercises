package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type pair struct {
	path string
	hash string
}

type fileList []string
type result map[string]fileList

func calcHash(path string) pair {
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer fi.Close()
	hash := md5.New()

	if _, err := io.Copy(hash, fi); err != nil {
		log.Fatal(err)
	}
	return pair{path, fmt.Sprintf("%x", hash.Sum(nil))}
}

func searchFiles(root string, limitC chan bool, pairC chan<- pair, wg *sync.WaitGroup) error {
	defer wg.Done()
	search := func(path string, info fs.FileInfo, err error) error {

		//ignore err param
		if err != nil {
			return err
		}
		if info.Mode().IsDir() && path != root {
			wg.Add(1)
			go searchFiles(path, limitC, pairC, wg)
			return filepath.SkipDir
		}

		if info.Mode().IsRegular() && info.Size() > 0 {
			wg.Add(1)
			go processF(path, pairC, wg, limitC)

		}

		return nil
	}

	limitC <- true
	defer func() { <-limitC }()

	return filepath.Walk(root, search)

}

func run(root string) result {
	workers := 2 * runtime.GOMAXPROCS(0)

	limitC := make(chan bool, workers)
	pairC := make(chan pair)
	resultC := make(chan result)
	wg := new(sync.WaitGroup)

	go collectResults(resultC, pairC)

	wg.Add(1)
	if err := searchFiles(root, limitC, pairC, wg); err != nil {
		return nil
	}

	wg.Wait()
	// close(pathC)

	// for range workers {
	// 	<-doneC
	// }

	close(pairC)

	return <-resultC

}

func collectResults(resultC chan<- result, pairC <-chan pair) {
	result := make(result)

	for h := range pairC {
		result[h.hash] = append(result[h.hash], h.path)
	}

	resultC <- result
}

func processF(path string, pairC chan<- pair, wg *sync.WaitGroup, limit chan bool) {
	defer wg.Done()

	limit <- true
	defer func() { <-limit }()

	pairC <- calcHash(path)

}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Root dir not find!")
	}
	root := os.Args[1]

	if res := run(root); res != nil {

		for k, v := range res {
			if len(v) > 1 {
				fmt.Println(k[len(k)-7:])

				for _, path := range v {
					fmt.Println("	", path)
				}
			}

		}

	}
}

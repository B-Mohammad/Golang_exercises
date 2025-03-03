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

func searchFiles(root string, pathC chan<- string) error {
	search := func(path string, info fs.FileInfo, err error) error {

		//ignore err param
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() && info.Size() > 0 {
			pathC <- path

		}

		return nil
	}

	return filepath.Walk(root, search)

}

func run(root string) result {
	pathC := make(chan string)
	doneC := make(chan bool)
	pairC := make(chan pair)
	resultC := make(chan result)

	workers := 2 * runtime.GOMAXPROCS(0)

	for range workers {
		go processF(pathC, pairC, doneC)
	}

	go collectResults(resultC, pairC)

	if err := searchFiles(root, pathC); err != nil {
		return nil
	}

	close(pathC)

	for range workers {
		<-doneC
	}

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

func processF(pathC <-chan string, pairC chan<- pair, doneC chan<- bool) {
	for p := range pathC {
		pairC <- calcHash(p)
	}
	doneC <- true

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

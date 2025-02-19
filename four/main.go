package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	bow := make(map[string]int)

	for scanner.Scan() {

		bow[scanner.Text()]++
	}

	fmt.Println("total unique:", len(bow))

	type kv struct {
		word  string
		count int
	}

	var sob []kv
	for k, v := range bow {
		sob = append(sob, kv{k, v})
	}

	sort.SliceStable(sob, func(i, j int) bool {
		return sob[i].count > sob[j].count
	})

	for _, kv := range sob {
		fmt.Println("word:", kv.word, "with count:", kv.count)
	}

}

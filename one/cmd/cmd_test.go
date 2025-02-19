package cmd

import (
	"testing"
)

func TestSayHello(t *testing.T) {

	subtests := []struct {
		names  []string
		result string
	}{
		{result: "Hello, World!"},
		{names: []string{"test"}, result: "Hello, test!"},
		{names: []string{"test", "test2"}, result: "Hello, test test2!"},
	}

	for _, item := range subtests {

		if got := Say(item.names); got != item.result {
			t.Errorf("wanted %s, (%v) but got %s", item.result, item.names, got)
		}

	}
}

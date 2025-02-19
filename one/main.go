package main

import (
	"fmt"
	"main/cmd"
	"os"
)

func main() {

	fmt.Println(cmd.Say(os.Args[1:]))

}

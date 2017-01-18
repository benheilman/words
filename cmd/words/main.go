// Read words and find all combinations that can be made using hex representation
package main

import (
	"github.com/mitchellh/go-linereader"
	"os"
	"words"
	"fmt"
)

func main() {
	ins := []<-chan string{}
	outs := []chan string{}
	for i := 0; i < 8; i++ {
		out := make(chan string, 10)
		outs = append(outs, out)
		ins = append(ins, words.Matcher(i, out))
	}
		
	merged := words.Merge(ins...)
	words.Duplicate(linereader.New(os.Stdin).Ch, outs...)
	for o := range merged {
		fmt.Println(o)
	}
}

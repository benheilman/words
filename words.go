// Read words and find all combinations that can be made using hex representation
package words

import (
	"regexp"
	"strings"
	"sync"
)

func Matcher(length int, in <-chan string) <-chan string {
	regex := regexp.MustCompile("^[a-f]+$")
	out := make(chan string, 10)
	go func() {
		for s := range in {
			if length == strings.Count(s, "")-1 && regex.MatchString(s) {
				out <- s
			}
		}
		close(out)
	}()
	return out
}

func Duplicate(input <-chan string, outputs ...chan string) {
	go func() {
	for in := range input {
		for _, o := range outputs {
			o <- in
		}
	}
	for _, o := range outputs {
		close(o)
	}
	}()
}

func Merge(inputs ...<-chan string) <-chan string {
	wg := new(sync.WaitGroup)
	out := make(chan string, 10)

	output := func(c <-chan string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(inputs))

	for _, in := range inputs {
		go output(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

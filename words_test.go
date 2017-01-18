package words

import (
	"testing"
	"strconv"
)

var (
	chanCount int = 10	
)

func TestDuplicate(t *testing.T) {
	outputs := []chan string{}
	for i := 0; i < chanCount; i++ {
		outputs = append(outputs, make(chan string, 10))
	}
	input := make(chan string, 10)
	input <- "test"
	Duplicate(input, outputs...)
	for i := 0; i < chanCount; i++ {
		message, ok := <- outputs[i]
		if !ok {
			t.Errorf("output channel closed")
		}
		if message != "test" {
			t.Errorf("message was not test")
		}
	}
	close(input)
	for i := 0; i < chanCount; i++ {
		_, ok := <- outputs[i]
		if ok {
			t.Errorf("output channel was not closed after input was")
		}
	}
}

func TestMerge(t *testing.T) {
	ins := []<-chan string{}
	for i := 0; i < chanCount; i++ {
		in := make(chan string, 10)
		in <- strconv.Itoa(i)
		ins = append(ins, in)
		close(in)
	}
	out := Merge(ins...)
	count := 0
	for _ = range out {
		count++
	}
	if count != chanCount {
		t.Errorf("Merged channel had %d messages instead of %d", count, chanCount)
	}	
}

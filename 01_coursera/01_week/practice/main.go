package main

import (
	"bufio"
	"fmt"
	"os"
)

func mapUniq() {
	in := bufio.NewScanner(os.Stdin)
	counter := map[string]int{}
	for in.Scan() {
		txt := in.Text()
		if counter[txt] == 0 {
			fmt.Println(txt)
		}
		counter[txt] += 1
	}
}

func sortedUniq() {
	var prev string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		txt := in.Text()
		if txt < prev {
			panic("Data not sorted!")
		}
		if txt != prev {
			fmt.Println(txt)
		}
		prev = txt
	}
}

func main() {
	sortedUniq()
}

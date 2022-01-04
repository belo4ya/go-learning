package main

import (
	"bufio"
	"fmt"
	"io"
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

func sortedUniq(input io.Reader, output io.Writer) error {
	var prev string
	in := bufio.NewScanner(input)
	for in.Scan() {
		txt := in.Text()
		if txt < prev {
			return fmt.Errorf("data not sorted")
		}
		if txt != prev {
			if _, err := fmt.Fprintln(output, txt); err != nil {
				return err
			}
		}
		prev = txt
	}
	return nil
}

func main() {
	err := sortedUniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

type Stack struct {
	stack []os.DirEntry
}

func NewStack(entries ...os.DirEntry) *Stack {
	s := new(Stack)
	s.stack = entries
	return s
}

func (s *Stack) Push(el os.DirEntry) {
	s.stack = append(s.stack, el)
}

func (s *Stack) Pop() os.DirEntry {
	i := len(s.stack) - 1
	el := s.stack[i]
	s.stack = s.stack[:i]
	return el
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	const (
		branch   string = "───"
		vertical string = "├"
		corner   string = "└"
	)

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	stack := NewStack()
	for _, file := range files {
		stack.Push(file)
	}

	var tree string
	for stack.Len() != 0 {
		file := stack.Pop()
		fmt.Println(file.Name())
	}

	if _, err := fmt.Fprint(out, tree); err != nil {
		return err
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

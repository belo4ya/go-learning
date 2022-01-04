package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type StackEl struct {
	Entry  os.DirEntry
	Parent string
	IsLast bool
}

func (e *StackEl) Level() int {
	return len(strings.Split(e.Parent, string(os.PathSeparator))) - 1
}

type Stack struct {
	stack []StackEl
}

func NewStack(entries ...StackEl) *Stack {
	s := new(Stack)
	s.stack = entries
	return s
}

func (s *Stack) Push(el StackEl) {
	s.stack = append(s.stack, el)
}

func (s *Stack) Pop() StackEl {
	i := len(s.stack) - 1
	el := s.stack[i]
	s.stack = s.stack[:i]
	return el
}

func (s *Stack) Seek() StackEl {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func getSizeLabel(entry os.DirEntry) (string, error) {
	if entry.IsDir() {
		return "", nil
	}

	info, err := entry.Info()
	if err != nil {
		return "", err
	}

	size := info.Size()
	if size == 0 {
		return " (empty)", nil
	}
	return fmt.Sprintf(" (%db)", size), nil
}

func readDir(name string, includeF bool) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	if !includeF {
		var dirs []os.DirEntry
		for _, entry := range entries {
			if entry.IsDir() {
				dirs = append(dirs, entry)
			}
		}
		return dirs, nil
	}

	return entries, nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	const (
		vertical string = "│"
		commonB  string = "├───"
		lastB    string = "└───"
	)

	files, err := readDir(path, printFiles)
	if err != nil {
		return err
	}

	stack := NewStack()
	if len(files) != 0 {
		stack.Push(StackEl{files[len(files)-1], path, true})
		for i := len(files) - 2; i > -1; i-- {
			stack.Push(StackEl{files[i], path, false})
		}
	}

	tree := ""
	levelMap := map[int]bool{}
	for stack.Len() != 0 {
		el := stack.Pop()
		entry, parent, level, isLast := el.Entry, el.Parent, el.Level(), el.IsLast

		if entry.IsDir() {
			newParent := filepath.Join(parent, entry.Name())
			files, err := readDir(newParent, printFiles)
			if err != nil {
				return err
			}
			if len(files) != 0 {
				stack.Push(StackEl{files[len(files)-1], newParent, true})
				for i := len(files) - 2; i > -1; i-- {
					stack.Push(StackEl{files[i], newParent, false})
				}
			}
		}

		sizeLabel, err := getSizeLabel(entry)
		if err != nil {
			return err
		}

		branch := commonB
		if isLast {
			branch = lastB
		}

		levelMap[level] = isLast
		space := ""
		for i := 0; i < level; i++ {
			if v, ok := levelMap[i]; ok {
				if v {
					space += "\t"
				} else {
					space += vertical + "\t"
				}
			}
		}

		tree += space + branch + entry.Name() + sizeLabel + "\n"
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

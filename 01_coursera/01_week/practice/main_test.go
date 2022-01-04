package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)

	if err := sortedUniq(in, out); err != nil {
		t.Errorf("test for OK Failed - error")
	}
	if result := out.String(); result != testOkResult {
		t.Errorf("test for OK Failed - results not match\n %v %v", result, testOkResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)

	if err := sortedUniq(in, out); err == nil {
		t.Errorf("test for ERROR Failed - not error")
	}
}

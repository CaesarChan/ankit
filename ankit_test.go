package ankit

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type MockNote []string

func (m MockNote) Fields() []string {
	return m
}

func TestOneNoteReader(t *testing.T) {
	var tests = []struct {
		fields []string
		csv    string
	}{
		{
			[]string{"field1", "field2", "field3"},
			"field1,field2,field3\n",
		},
		{
			[]string{"field1", "field2\n\"content\"", "field3"},
			"field1,\"field2\n\"\"content\"\"\",field3\n",
		},
	}

	var buf bytes.Buffer
	for _, tt := range tests {
		buf.Reset()
		r := MockNote(tt.fields)
		Copy(&buf, OneNoteReader(r))

		if buf.String() != tt.csv {
			t.Errorf("got %q, want %q", buf.String(), tt.csv)
		}
	}
}

func TestFileName(t *testing.T) {
	path := "/Users/SM1701/caesar/LeetCode-Go/leetcode/0020.Two-Sum"

	dirNames := strings.Split(path, "/")
	lastDirName := dirNames[len(dirNames)-1]
	fileName := strings.TrimLeft(lastDirName, "0")
	fileName = strings.Replace(fileName, ".", ". ", -1)
	questionName := strings.Replace(fileName, "-", " ", -1)
	fmt.Println("fileName", questionName)
}

func TestTrimLeft(t *testing.T) {
	s := "0001000200"
	left := strings.TrimLeft(s, "0")
	fmt.Println(left)
}

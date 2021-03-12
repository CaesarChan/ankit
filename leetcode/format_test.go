package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetAllFiles(t *testing.T) {
	_, err := RenameFiles("/Users/SM1701/caesar/LeetCode-Go/leetcode")
	if err != nil {
		t.Errorf("RenameFiles err: %v", err)
	}
}

func TestSubstr(t *testing.T) {
	str := "Shortest Unsorted Continuous Subarray.go"
	//str := "Shortest Unsorted Continuous Subarray_test.go"
	ok := strings.HasSuffix(str, "_test.go")
	fmt.Println(ok)

}
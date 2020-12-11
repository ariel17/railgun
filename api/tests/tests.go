package tests

import (
	"io/ioutil"
	"testing"
)

func GetGoldenFile(t *testing.T, path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return content
}

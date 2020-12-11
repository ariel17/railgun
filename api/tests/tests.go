package tests

import (
	"io/ioutil"
	"testing"
)

// GetGoldenFile reads the content of the file in the given path.
// Source: https://medium.com/@jarifibrahim/golden-files-why-you-should-use-them-47087ec994bf
func GetGoldenFile(t *testing.T, path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return content
}

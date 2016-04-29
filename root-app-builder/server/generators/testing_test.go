package generators_test

import (
	"io/ioutil"
	"os"
	"testing"
)

/*
 * Creates a temporary file and returns its name and a function
 * removing the file from the file system.
 */
func testTempFile(t *testing.T) (string, func()) {
	file, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	file.Close()

	return file.Name(), func() {
		os.Remove(file.Name())
	}
}

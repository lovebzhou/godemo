package main

import (
	"io"
	"log"
	"os"
)

func testWrite() {
	t, err := os.Create("./testwrite.data")
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	t.WriteString("testWrite")
	for i := 0; i < 5; i++ {
		t.Write([]byte(`func (f *File) Write(b []byte) (n int, err error) `))
		t.WriteAt([]byte(`func (f *File) WriteAt(b []byte, off int64) (n int, err error) `), 40)
		t.WriteString("func (f *File) WriteString(s string) (n int, err error) ")
	}
}

func testCopy() {
	f, err := os.Open("./testwrite.data")
	if err != nil {
		return
	}
	defer f.Close()

	t, err := os.Create("./testcopy.data")
	if err != nil {
		log.Fatal(err)
	}
	defer t.Close()

	io.Copy(t, f)
	t.WriteString("append by test copy")
}

func main() {
	testWrite()
	testCopy()
}

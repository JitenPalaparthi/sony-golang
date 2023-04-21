package main

import (
	"fmt"
	"os"
)

func main() {
	mf1 := &MyFile{FN: "demo.txt"}
	fmt.Fprintln(mf1, "Hello World!")

	// w1 := bufio.NewWriter(os.Stdout)
	// w1.Write([]byte("Hello World"))
	// w1.Close()

}

type MyFile struct {
	FN string
}

func (mf *MyFile) Write(p []byte) (n int, err error) {
	f, err := os.Create(mf.FN)

	if err != nil {
		return 0, err
	}
	n, err = f.Write(p)
	f.Close()
	return n, err
}

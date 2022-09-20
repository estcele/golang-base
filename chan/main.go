package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var sizes = make(chan int64)

const PATH_SCAN = "/Users/estcele/Downloads/"

// var files = make(map[string]int64)
var files sync.Map

func main() {
	// files := sync.Map{}
	fis, err := ioutil.ReadDir(PATH_SCAN)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	wg := sync.WaitGroup{}
	for _, fi := range fis {
		wg.Add(1)
		go func(fi os.FileInfo) {
			defer wg.Done()
			if !fi.IsDir() {
				// files[fi.Name()] = fi.Size()
				files.Store(fi.Name(), fi.Size())
				sizes <- fi.Size()
			}
		}(fi)
	}
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	fmt.Println("File Name	Size")
	files.Range(func(key, value any) bool {
		fmt.Printf("%s\t%d\n", key, value)
		return true
	})

	fmt.Printf("Total\t%d\n", total)

}

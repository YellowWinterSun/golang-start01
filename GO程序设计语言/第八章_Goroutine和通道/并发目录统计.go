package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func walkDir(dir string, fileSizeChannal chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			// 当前是文件夹
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizeChannal)
		} else {
			// 文件大小 单位：字节
			fmt.Printf("> %s\t%d\n", filepath.Join(dir, entry.Name()), entry.Size())
			fileSizeChannal <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	if entries, err := ioutil.ReadDir(dir); err != nil {
		fmt.Println(fmt.Errorf("Error:\n%v\n", err))
	} else {
		return entries
	}
	return nil
}

func main() {
	var filePaths []string = make([]string, 0)
	filePaths = append(filePaths, "D:\\test", "D:\\test2", "D:\\test3")

	fileSizeChannal := make(chan int64, len(filePaths))
	var wg sync.WaitGroup

	for _, path := range filePaths {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()
			fmt.Println("---", path)
			walkDir(path, fileSizeChannal)
		}(path)
	}

	// 等待关闭
	go func() {
		wg.Wait()
		fmt.Println("fileChannal -------> 关闭")
		close(fileSizeChannal)
	}()

	// 统计用的异步
	var nFiles, nBytes int64
	for size := range fileSizeChannal {
		nFiles++
		nBytes += size
	}

	fmt.Printf("%d files\t%d bytes\n", nFiles, nBytes)
}

package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path   string
	md5sum [md5.Size]byte
	err    error
}

func sumFiles(root string, done <-chan struct{}) (<-chan result, <-chan error) {
	resultCh := make(chan result)
	errCh := make(chan error, 1)

	go func() {
		// so this is the main go routine that will spwan many other go routines
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
				// check by creating a non regular file what the fuck it will retrun
			}

			wg.Add(1)
			go func() {
				// this is useful if one go routine is reading a big file and other a 2 byte file.
				// what is the use of this go routine??
				data, err := ioutil.ReadFile(path)
				select {
				case resultCh <- result{path, md5.Sum(data), err}: // add printf here to see how the flow work in this go routine
				case <-done: // here as well
				}
				wg.Done()
			}()
			select {
			case <-done:
				return errors.New("walk cancelled")
			default:
				return nil
			}
		})
		// the walk has returned here
		// who is monitoring the return of this go routine...why not wg.Add(1) for this one, put sleep and check
		go func() {
			wg.Wait()
			close(resultCh) // closing the write end of the channel
		}()
		errCh <- err
	}()
	return resultCh, errCh
}

func md5All(root string) (map[string][md5.Size]byte, error) {
	// create a done channel and md5all will call the done channel when it receives the results
	// it may do so even before receiveing all the values from result channel and error channel

	done := make(chan struct{})
	defer close(done)
	resultCh, errCh := sumFiles(root, done)

	m := make(map[string][md5.Size]byte)
	for r := range resultCh {
		if r.err != nil {
			return nil, r.err // this means een if one of the files shows error will it exit ??
		}
		m[r.path] = r.md5sum
	}

	if err := <-errCh; err != nil {
		return nil, err
	}

	return m, nil

}

func main() {
	m, err := md5All(os.Args[1])
	if err != nil {
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for out := range paths {
		fmt.Printf("%x : %s", m[paths[out]], out)
	}
}

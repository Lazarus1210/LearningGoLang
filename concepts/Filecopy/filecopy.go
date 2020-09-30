package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const kb = 1024
const twoHundKb = 200 * kb
const chunkSize int64 = twoHundKb

type data struct {
	offsetNum int
	d         byte
}

var dataCh chan data

func main() {
	srcFile := os.Args[1]
	//This is the source file name
	srcFileHandle, err := os.Open(srcFile)
	if err != nil {
		log.Printf("error opening the src file: %s", os.Args[1])
		log.Println(err)
		return
	}

	srcFileInfo, err := srcFileHandle.Stat()
	if err != nil {
		log.Printf("cannot stat the src file: %s\n", srcFile)
		log.Println(err)
		return
	}

	if !srcFileInfo.Mode().IsRegular() {
		log.Printf("%s : not a regualr file, this program is for regular file\n", srcFile)
		return
	}

	srcFileSize := srcFileInfo.Size()
	// take a new buffered reader
	reader := bufio.NewReader(srcFileHandle)
	buffer := make([]byte, srcFileSize)

	// take a new buffered reader and then read the number of bytes in the file
	bytesRead, err := reader.Read(buffer)
	if int64(bytesRead) != srcFileSize {
		log.Printf("bytes read did not match actual: %d  read : %d\n", srcFileSize, bytesRead)
		log.Println(err)
	}

	fmt.Printf("bytes read from the file : %d \n", bytesRead)
	fmt.Printf("%c\n", buffer[0])
	fmt.Printf("%c\n", buffer[srcFileSize-1])
	fmt.Println(srcFileSize % chunkSize)
	completeChunks := srcFileSize / chunkSize
	remainder := srcFileSize % chunkSize
	dataCh = make(chan data)

	for i := 0; int64(i) <= completeChunks; i++ {
		// spawn the go routines here
		go readFileAtchunk(srcFileHandle, i, int64(i)*chunkSize)
	}

}

func readFileAtchunk(fileHandle *os.File, chunkNum int, chunkoffset int64) (err error) {
	if fileHandle == nil {
		return err
	}

	// Go to the offset and read the chunck
	filename := fmt.Sprintf("%d_tempfile", chunkNum)
	//create a tempfile with this name
	tmpfile, err := os.Create(filename)
	if err != nil {
		log.Printf("cannot create file name : %s\n", filename)
		log.Println(err)
		return err
	}
	//move the pointer of this file to the offset
	fileHandle.Seek(chunkoffset, 0)
	//create a new reader for this goroutine
	reader := bufio.NewReader(fileHandle)
	i := 0

	for {
		if int64(i) >= chunkSize {
			log.Printf("counter has read the chunksize ")
			break
		}
		br, err := reader.ReadByte()
		if err == io.EOF {
			log.Println(err)
			break
		}
		select {
		case dataCh <- data{chunkNum, br}:
		default:
			return nil
		}

	}
	return nil
}

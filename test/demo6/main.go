package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var HK = [...]string{"HK", "AU"}

var UK = [...]string{"CN", "AU"}

func ifExist(key string, array []string) bool {
	sort.Strings(array)
	index := sort.SearchStrings(array, key)
	if index < len(array) && array[index] == key {
		return true
	}
	return false
}

func fileSplit(indicator string, arr []string, srcFile string, delimiter string, index string, channel chan string) {

	// open target file

	tgtFileIO, tgtErr := os.OpenFile(srcFile+"."+indicator, os.O_CREATE|os.O_WRONLY, 0666)

	if tgtErr != nil {

		log.Fatal("open file error :", tgtErr)

		return

	}

	defer tgtFileIO.Close()

	writer := bufio.NewWriter(tgtFileIO)

	// open source file

	srcFileIO, err := os.Open(srcFile)

	if err != nil {

		log.Fatal("open file error :", err)

		return

	}

	defer srcFileIO.Close()

	reader := bufio.NewReader(srcFileIO)

	// looping read line to split

	for {

		line, _, readErr := reader.ReadLine()

		if readErr == io.EOF {
			break
		}

		lineStr := string(line)

		arrIndex, _ := strconv.ParseInt(index, 0, 8)

		strArr := strings.Split(lineStr, delimiter)
		if strArr != nil && (len(strArr) > int(arrIndex)) {
			flag := strArr[arrIndex]
			if ifExist(flag, arr) {
				writer.WriteString(lineStr + "\n")
			}
		}
	}

	flushErr := writer.Flush()

	if flushErr != nil {

		log.Fatal("flush error:", flushErr)

		return

	}

	channel <- indicator

}

func main() {

	// -s="source.txt"

	var srcFile string

	flag.StringVar(&srcFile, "s", "", "source file")

	// -d="target.txt"

	var delimiter string

	flag.StringVar(&delimiter, "d", "", "delimiter")

	// -i="index"

	var index string

	flag.StringVar(&index, "i", "", "index")

	flag.Parse()

	channel := make(chan string, 2)

	go fileSplit("HK", HK[:], srcFile, delimiter, index, channel)

	go fileSplit("UK", UK[:], srcFile, delimiter, index, channel)

	for i := 0; i < 2; i++ {

		log.Println(<-channel)

	}

	log.Println("all completed")

}

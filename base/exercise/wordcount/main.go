package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 实现一个wordcount程序， 输入是文件路径，程序需要读取该文件，并输出文件内单词的数量。（文件内容均为英文）
*/

func CountWords(url string) int {
	file, err := os.Open(url)
	if err != nil {
		fmt.Println("文件打开失败")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)

	count := 0

	var err1 error = nil

	for {
		if err1 == io.EOF {
			break
		}

		readString, err := reader.ReadString('\n')
		err1 = err

		fields := strings.Fields(readString)

		count += len(fields)

	}

	return count

}

func main() {
	var url = "/Users/houjichao/Work/tmp/wordcount.txt"
	fmt.Println(CountWords(url))
}

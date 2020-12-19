package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {

	if len(commit) > 7 {
		commit = commit[:7]
	}
	fmt.Printf("xml2csv v%s (%s)\n", version, commit)

	var textPattern string
	var binaryPattern string
	var sizeStr string
	var outputPath string
	var help bool

	flag.StringVar(&textPattern, "t", "", "text pattern")
	flag.StringVar(&binaryPattern, "b", "", "binary pattern")
	flag.StringVar(&sizeStr, "s", "", "size")
	flag.StringVar(&outputPath, "o", "", "output file path")
	flag.BoolVar(&help, "h", false, "Help")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if sizeStr == "" || outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	size, err := parseSize(sizeStr)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := readBytePattern(textPattern, binaryPattern)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = write(file, bytes, size)
	if err != nil {
		log.Fatal(err)
	}
}

func write(file io.Writer, bytes []byte, size int64) error {

	byteLen := int64(len(bytes))
	writer := bufio.NewWriter(file)

	for i := int64(0); i < size; i++ {
		err := writer.WriteByte(bytes[i%byteLen])
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}

func readBytePattern(textPattern string, binaryPattern string) ([]byte, error) {

	if textPattern != "" {
		return []byte(textPattern), nil
	}

	return hex.DecodeString(binaryPattern)
}

func parseSize(sizeStr string) (int64, error) {

	sizeStr = strings.ToUpper(sizeStr)

	if strings.HasSuffix(sizeStr, "GB") {
		baseSize, err := parseInt64(sizeStr[:len(sizeStr)-len("GB")])
		if err != nil {
			return -1, err
		}

		return baseSize * 1024 * 1024 * 1024, nil
	}

	if strings.HasSuffix(sizeStr, "MB") {
		baseSize, err := parseInt64(sizeStr[:len(sizeStr)-len("MB")])
		if err != nil {
			return -1, err
		}

		return baseSize * 1024 * 1024, nil
	}

	if strings.HasSuffix(sizeStr, "KB") {
		baseSize, err := parseInt64(sizeStr[:len(sizeStr)-len("KB")])
		if err != nil {
			return -1, err
		}

		return baseSize * 1024, nil
	}

	return parseInt64(sizeStr)
}

func parseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

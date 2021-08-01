package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

var (
	Version = "dev"
	Commit  = "dev"
)

func main() {

	var textPattern string
	var binaryPattern string
	var sizeStr string
	var outputPath string
	var help bool

	flag.StringVarP(&textPattern, "text", "t", "", "Text pattern.")
	flag.StringVarP(&binaryPattern, "binary", "b", "", "Binary pattern. This is specified in hexadecimal. (ex. 0A0F)\nIf neither --text nor --binary is specified, the value will be 00.")
	flag.StringVarP(&sizeStr, "size", "s", "", "Size. Can be specified in KB, MB, or GB. (ex. 1gb)")
	flag.StringVarP(&outputPath, "output", "o", "", "Output file path.")
	flag.BoolVarP(&help, "help", "h", false, "Help.")
	flag.Parse()
	flag.CommandLine.SortFlags = false
	flag.Usage = func() {
		fmt.Printf("createfile v%s (%s)\n\n", Version, Commit)
		fmt.Fprint(os.Stderr, "Usage:\n  createfile [flags]\n\nFlags\n")
		flag.PrintDefaults()
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if sizeStr == "" || outputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	err := run(textPattern, binaryPattern, sizeStr, outputPath)
	if err != nil {
		log.Fatal(err)
	}
}

func run(textPattern string, binaryPattern string, sizeStr string, outputPath string) error {

	size, err := parseSize(sizeStr)
	if err != nil {
		return err
	}

	bytes, err := readBytePattern(textPattern, binaryPattern)
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return write(file, bytes, size)
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

	return writer.Flush()
}

func readBytePattern(textPattern string, binaryPattern string) ([]byte, error) {

	if textPattern != "" {
		return []byte(textPattern), nil
	}

	if binaryPattern != "" {
		return hex.DecodeString(binaryPattern)
	}

	return []byte{0x00}, nil
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

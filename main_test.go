package main

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadBytePattern_text(t *testing.T) {

	textPattern := "aあ"
	binaryPattern := ""

	bytes, err := readBytePattern(textPattern, binaryPattern)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if !reflect.DeepEqual(bytes, []byte{0x61, 0xe3, 0x81, 0x82}) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestReadBytePattern_binary(t *testing.T) {

	textPattern := ""
	binaryPattern := "61e38182"

	bytes, err := readBytePattern(textPattern, binaryPattern)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if !reflect.DeepEqual(bytes, []byte{0x61, 0xe3, 0x81, 0x82}) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestReadBytePattern_none(t *testing.T) {

	textPattern := ""
	binaryPattern := ""

	bytes, err := readBytePattern(textPattern, binaryPattern)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if !reflect.DeepEqual(bytes, []byte{0x00}) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestParseSize(t *testing.T) {

	sizeStr := "123"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 123 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_GB(t *testing.T) {

	sizeStr := "1GB"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1073741824 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_gb(t *testing.T) {

	sizeStr := "1gb"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1073741824 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_GB_parseError(t *testing.T) {

	sizeStr := "1ggb"

	_, err := parseSize(sizeStr)
	if err == nil || err.Error() != `strconv.ParseInt: parsing "1G": invalid syntax` {
		t.Fatal("failed test\n", err)
	}
}

func TestParseSize_MB(t *testing.T) {

	sizeStr := "1MB"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1048576 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_mb(t *testing.T) {

	sizeStr := "1mb"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1048576 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_MB_parseError(t *testing.T) {

	sizeStr := "mb"

	_, err := parseSize(sizeStr)
	if err == nil || err.Error() != `strconv.ParseInt: parsing "": invalid syntax` {
		t.Fatal("failed test\n", err)
	}
}

func TestParseSize_KB(t *testing.T) {

	sizeStr := "1KB"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1024 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_kb(t *testing.T) {

	sizeStr := "1kb"

	size, err := parseSize(sizeStr)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	if size != 1024 {
		t.Fatal("failed test\n", size)
	}
}

func TestParseSize_KB_parseError(t *testing.T) {

	sizeStr := "kkb"

	_, err := parseSize(sizeStr)
	if err == nil || err.Error() != `strconv.ParseInt: parsing "K": invalid syntax` {
		t.Fatal("failed test\n", err)
	}
}

func TestWrite(t *testing.T) {

	buf := new(bytes.Buffer)

	write(buf, []byte{0x00}, 1)

	bytes := buf.Bytes()
	if !reflect.DeepEqual(buf.Bytes(), []byte{0x00}) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestWrite_repeat(t *testing.T) {

	buf := new(bytes.Buffer)

	write(buf, []byte{0x00, 0x01}, 10)

	bytes := buf.Bytes()
	if !reflect.DeepEqual(buf.Bytes(), []byte{0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01}) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestMain(t *testing.T) {

	output := createTempFile(t)
	defer os.Remove(output)

	os.Args = []string{
		"createfile",
		"-t", "abc",
		"-s", "10",
		"-o", output,
	}

	main()

	bytes := readBytes(t, output)
	if !reflect.DeepEqual(bytes, []byte("abcabcabca")) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestRun(t *testing.T) {

	output := createTempFile(t)
	defer os.Remove(output)

	err := run("abc", "", "10", output)
	if err != nil {
		t.Fatal("failed test\n", err)
	}

	bytes := readBytes(t, output)
	if !reflect.DeepEqual(bytes, []byte("abcabcabca")) {
		t.Fatal("failed test\n", bytes)
	}
}

func TestRun_parseSizeError(t *testing.T) {

	output := createTempFile(t)
	defer os.Remove(output)

	err := run("abc", "", "10m", output)
	if err == nil || err.Error() != `strconv.ParseInt: parsing "10M": invalid syntax` {
		t.Fatal("failed test\n", err)
	}
}

func TestRun_readBytePatternError(t *testing.T) {

	output := createTempFile(t)
	defer os.Remove(output)

	err := run("", "0G", "10", output)
	if err == nil || err.Error() != `encoding/hex: invalid byte: U+0047 'G'` {
		t.Fatal("failed test\n", err)
	}
}

func TestRun_fileNotFound(t *testing.T) {

	output := createTempFile(t)
	defer os.Remove(output)

	err := run("abc", "", "10", filepath.Join(output, "__")) // 存在しないディレクトリ
	if err == nil {
		t.Fatal("failed test\n", err)
	}

	pathErr := err.(*os.PathError)
	if pathErr.Path != filepath.Join(output, "__") || pathErr.Op != "open" {
		t.Fatal("failed test\n", err)
	}
}

func createTempFile(t *testing.T) string {

	tempFile, err := os.CreateTemp("", "dat")
	if err != nil {
		t.Fatal("craete file failed\n", err)
	}

	err = tempFile.Close()
	if err != nil {
		t.Fatal("write file failed\n", err)
	}

	return tempFile.Name()
}

func readBytes(t *testing.T, name string) []byte {

	bo, err := os.ReadFile(name)
	if err != nil {
		t.Fatal("read failed\n", err)
	}

	return bo
}

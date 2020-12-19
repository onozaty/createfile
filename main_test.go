package main

import (
	"bytes"
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

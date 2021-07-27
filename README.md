# createfile

[![GitHub license](https://img.shields.io/github/license/onozaty/createfile)](https://github.com/onozaty/createfile/blob/main/LICENSE)
[![Test](https://github.com/onozaty/createfile/actions/workflows/test.yaml/badge.svg)](https://github.com/onozaty/createfile/actions/workflows/test.yaml)

The createfile can create files of a specified size.  
You can easily create large file.

## Usage

```
$ createfile -t abc -s 1GB -o 1gb.txt
```

or 

```
$ createfile -b 616263 -s 1GB -o 1gb.txt
```

The arguments are as follows.

```
Usage of createfile:
  -t, --text string     text pattern      
  -b, --binary string   binary pattern    
  -s, --size string     size
  -o, --output string   output file path  
  -h, --help            Help
```

## Install

You can download the binary from the following.

* https://github.com/onozaty/createfile/releases/latest

## License

MIT

## Author

[onozaty](https://github.com/onozaty)

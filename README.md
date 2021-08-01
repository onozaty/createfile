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
Usage:
  createfile [flags]

Flags
  -t, --text string     Text pattern.
  -b, --binary string   Binary pattern. This is specified in hexadecimal. (ex. 0A0F)
                        If neither --text nor --binary is specified, the value will be 00.
  -s, --size string     Size. Can be specified in KB, MB, or GB. (ex. 1gb)
  -o, --output string   Output file path.
  -h, --help            Help.
```

## Install

You can download the binary from the following.

* https://github.com/onozaty/createfile/releases/latest

## License

MIT

## Author

[onozaty](https://github.com/onozaty)

# covercmp
The covercmp command displays code coverage changes between unit tests.
This plugin is heavily inspired by [benchcmp](https://godoc.org/golang.org/x/tools/cmd/benchcmp)

To measure the coverage impact of a change, use 'go test' to run coverage before and after the change:
```bash
go test -count=1 -cover > before.txt
# make changes
go test -count=1 -cover > after.txt
```

Then feed the coverage results to covercmp:
```bash
covercmp go before.txt after.txt
```

Covercmp will summarize and display the coverage changes, in a format like this:
```bash
package            before cov     after cov     delta
compress/bzip2     87.5%          97.5%         10.0%
compress/flate     82.7%          93.8%         11.1%
compress/gzip      89.9%          89.9%         0.0%
compress/lzw       87.2%          87.2%         0.0%
compress/zlib      84.7%          84.7%         0.0%
```

## Installation
- Download the latest binary for your OS release from the [releases page](https://github.com/arithran/covercmp/releases)
- Rename the file to `covercmp`
- Make the file executable (`chmod +x covercmp`)
- Checkout the help menu for usage instructions `covercmp help`
- (Optional Step) Move it to a folder in your PATH variable. (`mv covercmp /bin`)


## Supported Langages

- [x] [Golang](./driver/golang)
- [ ] TBD

## Usage
```bash
covercmp go before.txt after.txt
```

## Contribute
1. Pull requests are welcome
2. The API is open enough to support other Langages. See [Golang driver](./driver/golang) implementation for more details

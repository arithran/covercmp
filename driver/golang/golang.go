package golang

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/arithran/covercmp/driver"
)

var re = regexp.MustCompile(`\d+\.\d+`)

// Driver //
type Driver struct{}

// New //
func New() driver.Parser {
	return &Driver{}
}

// Parse //
func (*Driver) Parse(r io.Reader) (driver.Set, error) {
	cs := make(driver.Set)
	scan := bufio.NewScanner(r)
	ord := 0
	for scan.Scan() {
		if c, err := parseLine(scan.Text()); err == nil {
			c.Order = ord
			ord++
			cs[c.Package] = c
		}
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}

func parseLine(line string) (driver.Coverage, error) {
	// split by tab character
	fields := strings.Split(line, "\t")

	// trim any spaces
	for k, v := range fields {
		fields[k] = strings.TrimSpace(v)
	}

	switch fields[0] {
	case "ok":
		var percent float64
		strPercent := re.FindAllString(fields[3], 1)
		percent, _ = strconv.ParseFloat(strPercent[0], 32)

		return driver.Coverage{
			Package: fields[1],
			Pecent:  float32(percent),
		}, nil
	case "?":
		return driver.Coverage{
			Package: fields[1],
			Pecent:  0,
		}, nil
	default:
		// lets assume that this is harmless
		return driver.Coverage{}, errors.New("invalid line")
	}
}

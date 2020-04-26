package golang

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/arithran/covercmp/driver"
)

// Fix later
// func TestDriver_Parse(t *testing.T) {
// 	testReader, cleanup := getTestReader(t)
// 	defer cleanup()
//
// 	type args struct {
// 		r io.Reader
// 	}
// 	tests := []struct {
// 		name    string
// 		d       *Driver
// 		args    args
// 		want    []driver.Coverage
// 		wantErr bool
// 	}{
// 		{
// 			name: "testing golang parse",
// 			d:    &Driver{},
// 			args: args{
// 				r: testReader,
// 			},
// 			want: []driver.Coverage{
// 				{Package: "compress/bzip2", Pecent: 87.5, Order: 0},
// 				{Package: "compress/flate", Pecent: 93.8, Order: 1},
// 				{Package: "compress/gzip", Pecent: 89.9, Order: 2},
// 				{Package: "compress/lzw", Pecent: 87.2, Order: 3},
// 				{Package: "compress/zlib", Pecent: 84.7, Order: 4},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := &Driver{}
// 			got, err := d.Parse(tt.args.r)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Driver.Parse() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Driver.Parse() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    driver.Coverage
		wantErr bool
	}{
		{
			name: "successfully parsed a line",
			args: args{
				line: "ok  	compress/bzip2	0.186s	coverage: 87.5% of statements",
			},
			want: driver.Coverage{
				Package: "compress/bzip2",
				Pecent:  87.5,
			},
			wantErr: false,
		},
		{
			name: "successfully parsed a (cached) line",
			args: args{
				line: "ok  	compress/bzip2	(cached)	coverage: 87.5% of statements",
			},
			want: driver.Coverage{
				Package: "compress/bzip2",
				Pecent:  87.5,
			},
			wantErr: false,
		},
		{
			name: "successfully parsed a [no test files] line",
			args: args{
				line: "?   	github.com/arithran/testpackage	[no test files]",
			},
			want: driver.Coverage{
				Package: "github.com/arithran/testpackage",
				Pecent:  0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTestReader(t *testing.T) (io.Reader, func()) {
	t.Helper()

	f, err := os.Open("./testdata/out.txt")
	if err != nil {
		t.Error(err)
	}

	return f, func() {
		f.Close()
	}
}

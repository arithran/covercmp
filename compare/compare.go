package compare

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/arithran/covercmp/driver"
)

// Cmp will compare the before and after coverage
func Cmp(parser driver.Parser, beforePath, afterPath string) error {
	// open before and after files
	beforeFile, err := os.Open(beforePath)
	if err != nil {
		return err
	}
	defer beforeFile.Close()
	afterFile, err := os.Open(afterPath)
	if err != nil {
		return err
	}
	defer afterFile.Close()

	// parse coverage
	beforeCov, err := parser.Parse(beforeFile)
	if err != nil {
		return err
	}
	afterCov, err := parser.Parse(afterFile)
	if err != nil {
		return err
	}

	cmps := correlate(beforeCov, afterCov)
	sort.Sort(ByParseOrder(cmps))

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 5, ' ', 0)
	defer w.Flush()

	// header
	fmt.Fprint(w, "package\tbefore cov\tafter cov\tdelta\n")
	_ = cmps
	for _, cmp := range cmps {
		fmt.Fprintf(w, "%s\t%.1f%%\t%.1f%%\t%.1f%%\n", cmp.Package(), cmp.Before.Pecent, cmp.After.Pecent, cmp.Delta())

	}
	return nil

}

// CoverageCmp is a pair of coverage
type CoverageCmp struct {
	Before driver.Coverage
	After  driver.Coverage
}

func correlate(before, after driver.Set) []CoverageCmp {
	cmps := make([]CoverageCmp, 0, len(after))
	for k := range after {
		cmps = append(cmps, CoverageCmp{
			Before: before[k],
			After:  after[k],
		})
	}

	return cmps
}

func (cmp *CoverageCmp) Package() string { return cmp.After.Package }
func (cmp *CoverageCmp) Delta() float32  { return cmp.After.Pecent - cmp.Before.Pecent }

// ByParseOrder sorts CoverageCmp to match the order in
// which the Before coverages were presented to driver.Parser.
type ByParseOrder []CoverageCmp

func (x ByParseOrder) Len() int           { return len(x) }
func (x ByParseOrder) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x ByParseOrder) Less(i, j int) bool { return x[i].Before.Order < x[j].Before.Order }

package camelotgo

import (
	"os/exec"
	"strconv"
	"strings"
)

type Camelot struct {
	globalOptions []string
}

// Usage: camelot [OPTIONS] COMMAND [ARGS]...
//
//   Camelot: PDF Table Extraction for Humans
//
// Options:
//   --version                       Show the version and exit.
//   -q, --quiet TEXT                Suppress logs and warnings.
//   -p, --pages TEXT                Comma-separated page numbers. Example: 1,3,4
//                                   or 1,4-end or all.
//   -pw, --password TEXT            Password for decryption.
//   -o, --output TEXT               Output file path.
//   -f, --format [csv|excel|html|json|markdown|sqlite]
//                                   Output file format.
//   -z, --zip                       Create ZIP archive.
//   -split, --split_text            Split text that spans across multiple cells.
//   -flag, --flag_size              Flag text based on font size. Useful to
//                                   detect super/subscripts.
//   -strip, --strip_text TEXT       Characters that should be stripped from a
//                                   string before assigning it to a cell.
//   -M, --margins FLOAT...          PDFMiner char_margin, line_margin and
//                                   word_margin.
//   --help                          Show this message and exit.
//
// Commands:
//   lattice  Use lines between text to parse the table.
//   stream   Use spaces between text to parse the table.
func NewCamelot(options ...GlobalOption) Camelot {
	var globalOptions []string

	for _, option := range options {
		globalOptions = append(globalOptions, option()...)
	}

	return Camelot{globalOptions: globalOptions}
}

func (c Camelot) Exec() ([]byte, error) {
	cmd := exec.Command("camelot", c.globalOptions...)

	return cmd.CombinedOutput()
}

type GlobalOption func() []string

func Version() GlobalOption {
	return func() []string {
		return []string{"--help"}
	}
}

func Quiet(o string) GlobalOption {
	return func() []string {
		return []string{"--quiet", o}
	}
}

func Pages(o string) GlobalOption {
	return func() []string {
		return []string{"--pages", o}
	}
}

func Password(o string) GlobalOption {
	return func() []string {
		return []string{"--password", o}
	}
}

func Output(o string) GlobalOption {
	return func() []string {
		return []string{"--output", o}
	}
}

func Format(o string) GlobalOption {
	return func() []string {
		return []string{"--format", o}
	}
}

func Zip() GlobalOption {
	return func() []string {
		return []string{"--zip"}
	}
}

func SplitText() GlobalOption {
	return func() []string {
		return []string{"--split_text"}
	}
}

func FlagSize(o int) GlobalOption {
	return func() []string {
		return []string{"--flag_size", strconv.Itoa(o)}
	}
}

func StripText(o string) GlobalOption {
	return func() []string {
		return []string{"--strip_text", o}
	}
}

func Margins(o float64) GlobalOption {
	return func() []string {
		return []string{"--margins", strconv.FormatFloat(o, 'f', -1, 64)}
	}
}

func Help() GlobalOption {
	return func() []string {
		return []string{"--help"}
	}
}

// Usage: camelot lattice [OPTIONS] FILEPATH
//
// Use lines between text to parse the table.
//
// Options:
// -R, --table_regions TEXT        Page regions to analyze. Example:
// 																x1,y1,x2,y2 where x1, y1 -> left-top and x2,
// 																y2 -> right-bottom.
// -T, --table_areas TEXT          Table areas to process. Example: x1,y1,x2,y2
// 																where x1, y1 -> left-top and x2, y2 ->
// 																right-bottom.
// -back, --process_background     Process background lines.
// -scale, --line_scale INTEGER    Line size scaling factor. The larger the
// 																value, the smaller the detected lines.
// -copy, --copy_text [h|v]        Direction in which text in a spanning cell
// 																will be copied over.
// -shift, --shift_text [|l|r|t|b]
// 																Direction in which text in a spanning cell
// 																will flow.
// -l, --line_tol INTEGER          Tolerance parameter used to merge close
// 																vertical and horizontal lines.
// -j, --joint_tol INTEGER         Tolerance parameter used to decide whether
// 																the detected lines and points lie close to
// 																each other.
// -block, --threshold_blocksize INTEGER
// 																For adaptive thresholding, size of a pixel
// 																neighborhood that is used to calculate a
// 																threshold value for the pixel. Example: 3,
// 																5, 7, and so on.
// -const, --threshold_constant INTEGER
// 																For adaptive thresholding, constant
// 																subtracted from the mean or weighted mean.
// 																Normally, it is positive but may be zero or
// 																negative as well.
// -I, --iterations INTEGER        Number of times for erosion/dilation will be
// 																applied.
// -res, --resolution INTEGER      Resolution used for PDF to PNG conversion.
// -plot, --plot_type [text|grid|contour|joint|line]
// 																Plot elements found on PDF page for visual
// 																debugging.
// --help                          Show this message and exit.
func (c Camelot) Lattice(inputFile string, options ...LatticeOption) ([]byte, error) {
	var latticeOptions []string

	for _, option := range options {
		latticeOptions = append(latticeOptions, option()...)
	}

	cmd := exec.Command("camelot", append(append(append(c.globalOptions, "lattice"), latticeOptions...), inputFile)...)

	return cmd.CombinedOutput()
}

type LatticeOption func() []string

func LatticeTableRegions(o string) LatticeOption {
	return func() []string {
		return []string{"--table_regions", o}
	}
}

func LatticeTableAreas(o string) LatticeOption {
	return func() []string {
		return []string{"--table_areas", o}
	}
}

func ProcessingBackground() LatticeOption {
	return func() []string {
		return []string{"--processing_background"}
	}
}

func LineScale(o int) LatticeOption {
	return func() []string {
		return []string{"--line_scale", strconv.Itoa(o)}
	}
}

type CopyTextOption string

const (
	CopyTextHorizontal = "h"
	CopyTextVertical   = "v"
)

func CopyText(o CopyTextOption) LatticeOption {
	return func() []string {
		return []string{"--copy_text", string(o)}
	}
}

type ShiftTextOption string

const (
	ShiftTextLeft   = "l"
	ShiftTextTop    = "t"
	ShiftTextRight  = "r"
	ShiftTextBottom = "b"
)

func ShiftText(o []ShiftTextOption) LatticeOption {
	return func() []string {
		options := make([]string, len(o))
		for i, option := range o {
			options[i] = string(option)
		}

		return []string{"--shift_text", strings.Join(options, ",")}
	}
}

func LineTolerance(o int) LatticeOption {
	return func() []string {
		return []string{"--line_tol", strconv.Itoa(o)}
	}
}

func JointTolerance(o int) LatticeOption {
	return func() []string {
		return []string{"--joint_tol", strconv.Itoa(o)}
	}
}

func ThresholdBlocksize(o int) LatticeOption {
	return func() []string {
		return []string{"--threshold_blocksize", strconv.Itoa(o)}
	}
}

func ThresholdConstant(o int) LatticeOption {
	return func() []string {
		return []string{"--threshold_constant", strconv.Itoa(o)}
	}
}

func Iterations(o int) LatticeOption {
	return func() []string {
		return []string{"--iterations", strconv.Itoa(o)}
	}
}

func Resolution(o int) LatticeOption {
	return func() []string {
		return []string{"--resolution", strconv.Itoa(o)}
	}
}

type LatticePlotTypeOption string

const (
	LatticePlotTypeText    = "text"
	LatticePlotTypeGrid    = "grid"
	LatticePlotTypeContour = "contour"
	LatticePlotTypeJoint   = "joint"
	LatticePlotTypeLine    = "line"
)

func LatticePlotType(o LatticePlotTypeOption) LatticeOption {
	return func() []string {
		return []string{"--plot_type", string(o)}
	}
}

func LatticeHelp() LatticeOption {
	return func() []string {
		return []string{"--help"}
	}
}

// Usage: camelot stream [OPTIONS] FILEPATH
//
//   Use spaces between text to parse the table.
//
// Options:
//   -R, --table_regions TEXT        Page regions to analyze. Example:
//                                   x1,y1,x2,y2 where x1, y1 -> left-top and x2,
//                                   y2 -> right-bottom.
//   -T, --table_areas TEXT          Table areas to process. Example: x1,y1,x2,y2
//                                   where x1, y1 -> left-top and x2, y2 ->
//                                   right-bottom.
//   -C, --columns TEXT              X coordinates of column separators.
//   -e, --edge_tol INTEGER          Tolerance parameter for extending textedges
//                                   vertically.
//   -r, --row_tol INTEGER           Tolerance parameter used to combine text
//                                   vertically, to generate rows.
//   -c, --column_tol INTEGER        Tolerance parameter used to combine text
//                                   horizontally, to generate columns.
//   -plot, --plot_type [text|grid|contour|textedge]
//                                   Plot elements found on PDF page for visual
//                                   debugging.
//   --help                          Show this message and exit.
func (c Camelot) Stream(inputFile string, options ...StreamOption) ([]byte, error) {
	var streamOptions []string

	for _, option := range options {
		streamOptions = append(streamOptions, option()...)
	}

	cmd := exec.Command("camelot", append(append(append(c.globalOptions, "stream"), streamOptions...), inputFile)...)

	return cmd.CombinedOutput()
}

type StreamOption func() []string

func StreamTableRegions(o string) StreamOption {
	return func() []string {
		return []string{"--table_regions", o}
	}
}

func StreamTableAreas(o string) StreamOption {
	return func() []string {
		return []string{"--table_areas", o}
	}
}

func Columns(o string) StreamOption {
	return func() []string {
		return []string{"--columns", o}
	}
}

func EdgeTolerance(o int) StreamOption {
	return func() []string {
		return []string{"--edge_tol", strconv.Itoa(o)}
	}
}

func RowTolerance(o int) StreamOption {
	return func() []string {
		return []string{"--row_tol", strconv.Itoa(o)}
	}
}

func ColumnTolerance(o int) StreamOption {
	return func() []string {
		return []string{"--column_tol", strconv.Itoa(o)}
	}
}

type StreamPlotTypeOption string

const (
	StreamPlotTypeText     = "text"
	StreamPlotTypeGrid     = "grid"
	StreamPlotTypeContour  = "contour"
	StreamPlotTypeTextEdge = "textedge"
)

func StreamPlotType(o StreamPlotTypeOption) StreamOption {
	return func() []string {
		return []string{"--plot_type", string(o)}
	}
}

func StreamHelp() StreamOption {
	return func() []string {
		return []string{"--help"}
	}
}

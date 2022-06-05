package camelotgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var helpText = `Usage: camelot [OPTIONS] COMMAND [ARGS]...

  Camelot: PDF Table Extraction for Humans

Options:
  --version                       Show the version and exit.
  -q, --quiet TEXT                Suppress logs and warnings.
  -p, --pages TEXT                Comma-separated page numbers. Example: 1,3,4
                                  or 1,4-end or all.
  -pw, --password TEXT            Password for decryption.
  -o, --output TEXT               Output file path.
  -f, --format [csv|excel|html|json|markdown|sqlite]
                                  Output file format.
  -z, --zip                       Create ZIP archive.
  -split, --split_text            Split text that spans across multiple cells.
  -flag, --flag_size              Flag text based on font size. Useful to
                                  detect super/subscripts.
  -strip, --strip_text TEXT       Characters that should be stripped from a
                                  string before assigning it to a cell.
  -M, --margins FLOAT...          PDFMiner char_margin, line_margin and
                                  word_margin.
  --help                          Show this message and exit.

Commands:
  lattice  Use lines between text to parse the table.
  stream   Use spaces between text to parse the table.
`

func TestCamelot(t *testing.T) {
	b, err := NewCamelot(Help()).Exec()
	require.NoError(t, err)
	assert.Equal(t, helpText, string(b))
}

var latticeHelpText = `Usage: camelot lattice [OPTIONS] FILEPATH

  Use lines between text to parse the table.

Options:
  -R, --table_regions TEXT        Page regions to analyze. Example:
                                  x1,y1,x2,y2 where x1, y1 -> left-top and x2,
                                  y2 -> right-bottom.
  -T, --table_areas TEXT          Table areas to process. Example: x1,y1,x2,y2
                                  where x1, y1 -> left-top and x2, y2 ->
                                  right-bottom.
  -back, --process_background     Process background lines.
  -scale, --line_scale INTEGER    Line size scaling factor. The larger the
                                  value, the smaller the detected lines.
  -copy, --copy_text [h|v]        Direction in which text in a spanning cell
                                  will be copied over.
  -shift, --shift_text [|l|r|t|b]
                                  Direction in which text in a spanning cell
                                  will flow.
  -l, --line_tol INTEGER          Tolerance parameter used to merge close
                                  vertical and horizontal lines.
  -j, --joint_tol INTEGER         Tolerance parameter used to decide whether
                                  the detected lines and points lie close to
                                  each other.
  -block, --threshold_blocksize INTEGER
                                  For adaptive thresholding, size of a pixel
                                  neighborhood that is used to calculate a
                                  threshold value for the pixel. Example: 3,
                                  5, 7, and so on.
  -const, --threshold_constant INTEGER
                                  For adaptive thresholding, constant
                                  subtracted from the mean or weighted mean.
                                  Normally, it is positive but may be zero or
                                  negative as well.
  -I, --iterations INTEGER        Number of times for erosion/dilation will be
                                  applied.
  -res, --resolution INTEGER      Resolution used for PDF to PNG conversion.
  -plot, --plot_type [text|grid|contour|joint|line]
                                  Plot elements found on PDF page for visual
                                  debugging.
  --help                          Show this message and exit.
`

func TestLattice(t *testing.T) {
	b, err := NewCamelot().Lattice("", LatticeHelp())
	require.NoError(t, err)
	assert.Equal(t, latticeHelpText, string(b))
}

var streamHelpText = `Usage: camelot stream [OPTIONS] FILEPATH

  Use spaces between text to parse the table.

Options:
  -R, --table_regions TEXT        Page regions to analyze. Example:
                                  x1,y1,x2,y2 where x1, y1 -> left-top and x2,
                                  y2 -> right-bottom.
  -T, --table_areas TEXT          Table areas to process. Example: x1,y1,x2,y2
                                  where x1, y1 -> left-top and x2, y2 ->
                                  right-bottom.
  -C, --columns TEXT              X coordinates of column separators.
  -e, --edge_tol INTEGER          Tolerance parameter for extending textedges
                                  vertically.
  -r, --row_tol INTEGER           Tolerance parameter used to combine text
                                  vertically, to generate rows.
  -c, --column_tol INTEGER        Tolerance parameter used to combine text
                                  horizontally, to generate columns.
  -plot, --plot_type [text|grid|contour|textedge]
                                  Plot elements found on PDF page for visual
                                  debugging.
  --help                          Show this message and exit.
`

func TestStream(t *testing.T) {
	b, err := NewCamelot().Stream("", StreamHelp())
	require.NoError(t, err)
	assert.Equal(t, streamHelpText, string(b))
}

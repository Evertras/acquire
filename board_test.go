package acquire

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintEmptyBoard(t *testing.T) {
	buf := &bytes.Buffer{}
	bufExpected := &bytes.Buffer{}

	fmt.Fprintln(bufExpected, "   A B C D E F G H I")
	fmt.Fprintln(bufExpected, " 1 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 2 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 3 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 4 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 5 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 6 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 7 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 8 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, " 9 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, "10 - - - - - - - - - ")
	fmt.Fprintln(bufExpected, "11 - - - - - - - - - ")

	b := NewBoard()

	b.PrintBoard(buf)

	result := buf.String()
	expected := bufExpected.String()

	if result != expected {
		t.Errorf("Did not get expected board result:\n\n%s\n\n%s", result, expected)

		for i, r := range result {
			if expected[i] != result[i] {
				t.Errorf("Badness %c vs %c at %d", r, expected[i], i)
			}
		}
	}
}

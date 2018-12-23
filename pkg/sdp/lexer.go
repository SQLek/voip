package sdp

import (
	"github.com/SQLek/voip/pkg/ll"
)

//============================================================================
// lexing token
// token-char = %x21 / %x23-27 / %x2A-2B / %x2D-2E /
//              %x30-39 / %x41-5A / %x5E-7E
// token      = 1*(token-char)

var rangesToken = []ll.ByteRange {
	{0x21, 0x21},
	{0x23, 0x27},
	{0x2A, 0x2B},
	{0x2D, 0x2E},
	{0x30, 0x39},
	{0x41, 0x5A},
	{0x5E, 0x7E},
}

func checkToken(reader ll.Reader) (bool, error) {
	return ll.Check(reader, rangesToken)
}

func collectToken(reader ll.Reader) (string, error) {
	return ll.Collect(reader, rangesToken)
}

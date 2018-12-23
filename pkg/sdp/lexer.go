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

//============================================================================
// lexing text
// text =                byte-string
//                       ;default is to interpret this as UTF8 text.
//                       ;ISO 8859-1 requires "a=charset:ISO-8859-1"
//                       ;session-level attribute to be used
//
// byte-string =         1*(%x01-09/%x0B-0C/%x0E-FF)
//                       ;any byte except NUL, CR, or LF

var rangesText = []ll.ByteRange {
	{0x01, 0x09},
	{0x0B, 0x0C},
	{0x0E, 0xFF},
}

func collectText(reader ll.Reader) (string, error) {
	return ll.Collect(reader,rangesText)
}

//============================================================================
// lexing non-ws
// non-ws-string = 1*(VCHAR/%x80-FF)
//                 ;string of visible characters
// from RFC4234
// VCHAR         = %x21-7E
//                 ; visible (printing) characters

var rangesNonWS = []ll.ByteRange {
	{0x21, 0xFF},
}

func collectNonWS(reader ll.Reader) (string, error) {
	return ll.Collect(reader,rangesNonWS)
}

//============================================================================
// lexing digit
// from RFC4234
// DIGIT        =  %x30-39
//              ; 0-9

var rangesDigit = []ll.ByteRange {
	{0x30, 0x39},
}

func collectDigit(reader ll.Reader) (string, error) {
	return ll.Collect(reader,rangesDigit)
}

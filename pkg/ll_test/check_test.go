package ll_test

import (
	"bufio"
	"github.com/SQLek/voip/pkg/ll"
	"strings"
	"testing"
)

var checkTests = []struct{
	b bool
	d string
	s []ll.ByteRange
}{
	{true, "afafa--",[]ll.ByteRange{{'a','a'},{'f','f'}}},
}

func TestCheck(t *testing.T) {

	for _, test := range checkTests {

		reader := bufio.NewReader(strings.NewReader(test.d))
		b, err := ll.Check(reader,test.s)

		if err != nil {
			t.Error("Got error:", err)
		}

		if b != test.b {
			t.Errorf("Expected %t got %t", test.b, b)
		}

	}

}

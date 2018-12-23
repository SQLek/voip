package ll_test

import (
	"bufio"
	"github.com/SQLek/voip/pkg/ll"
	"strings"
	"testing"
)

var keywordTests = []struct{
	b bool
	d string
	k string
}{
	{true, "v=0\r\n--","v=0\r\n"},
}

func TestKeyword(t *testing.T) {

	for _, test := range keywordTests {

		reader := bufio.NewReader(strings.NewReader(test.d))
		b, err := ll.Keyword(reader,test.k)

		if err != nil {
			t.Error("Got error:", err)
		}

		if b != test.b {
			t.Errorf("Expected %t got %t", test.b, b)
		}

	}

}

package ll_test

import (
	"bufio"
	"github.com/SQLek/voip/pkg/ll"
	"strings"
	"testing"
)

var collectTests = []struct{
	e string
	d string
	s []ll.ByteRange
}{
	{"afafa", "afafa--",[]ll.ByteRange{{'a','a'},{'f','f'}}},
}

func TestCollect(t *testing.T) {

	for _, test := range collectTests {

		reader := bufio.NewReader(strings.NewReader(test.d))
		s, err := ll.Collect(reader,test.s)

		if err != nil {
			t.Error("Got error:", err)
		}

		if s != test.e {
			t.Errorf("Expected %s got %s", test.e, s)
		}

	}

}

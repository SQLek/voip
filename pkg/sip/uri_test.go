package sip

import (
	"net/url"
	"testing"
)

var urlTests = []struct {
	urlString string
	host      string
}{
	{
		"sips:bob@biloxi.example.com",
		"biloxi.example.com",
	},
	{
		"client.biloxi.example.com:5061;branch=z9hG4bKnashds7",
		"client.biloxi.example.com:5061",
	},
}

func TestStdUrl(t *testing.T) {

	for _, test := range urlTests {

		u, err := url.Parse(test.urlString)
		if err != nil {
			t.Error(err)
		}
		if u.Host != test.host {
			t.Errorf("Expected '%s' got '%s'\n", test.host, u.Host)
		}
		t.Log(u)

	}

}

package sip

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Dialog interface {
	Close() error

	DemoRead() error

	DemoWrite() error
}

type tcpDialog struct {
	conn net.Conn

	reader *bufio.Reader
}

func Dial(addr string) (Dialog, error) {

	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &tcpDialog{
		conn:   c,
		reader: bufio.NewReader(c),
	}, nil

}

func (tDial *tcpDialog) Close() error {
	return tDial.conn.Close()
}

func (tDial *tcpDialog) DemoRead() error {
	_, err := tDial.reader.WriteTo(os.Stdout)
	return err
}

func (tDial *tcpDialog) DemoWrite() error {

	_, err := fmt.Fprint(tDial.conn,
		"OPTIONS sip:carol@chicago.com SIP/2.0\r\n",
		"Via: SIP/2.0/UDP pc33.atlanta.com;branch=z9hG4bKhjhs8ass877\r\n",
		"Max-Forwards: 70\r\n",
		"To: <sip:carol@chicago.com>\r\n",
		"From: Alice <sip:alice@atlanta.com>;tag=1928301774\r\n",
		"Call-ID: a84b4c76e66710\r\n",
		"CSeq: 63104 OPTIONS\r\n",
		"Contact: <sip:alice@pc33.atlanta.com>\r\n",
		"Accept: application/sdp\r\n",
		"Content-Length: 0\r\n",
	)
	return err
}

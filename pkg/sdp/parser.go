package sdp

import "bufio"

//============================================================================
// LL0 parsing functions
// each commented with corresponding ABNF
// for more info see https://tools.ietf.org/html/rfc4566#section-9

// session-description = proto-version
//                       origin-field
//                       session-name-field
//                       information-field
//                       uri-field
//                       email-fields
//                       phone-fields
//                       connection-field
//                       bandwidth-fields
//                       time-fields
//                       key-field
//                       attribute-fields
//                       media-descriptions
//
func parsePacket(reader bufio.Reader) (data interface{}, err error) {



	return
}

func parse(reader bufio.Reader) (data interface{}, err error) {



	return
}

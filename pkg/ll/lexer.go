package ll

// range of bytes, from to inclusive
type ByteRange struct {
	From, To byte
}

func Check(reader Reader, spec []ByteRange) (bool, error) {

	d, err := reader.Peek(1)
	if err != nil {
		return false, err
	}

	for _, br := range spec {
		if d[0] >= br.From && d[0] <= br.To {
			return true, nil
		}
	}

	return false, nil
}

func Collect(reader Reader, spec []ByteRange) (string, error) {

	i := 0
	var d []byte
	var err error

	charLoop:
	for {

		d, err = reader.Peek(i+1)
		if err != nil {
			return "", err
		}
		b := d[len(d)-1]

		for _, br := range spec {
			if b >= br.From && b <= br.To {
				i++
				continue charLoop
			}
		}

		break

	}

	s := string(d[:i])
	_, err = reader.Discard(i)
	return s,err
}

func Keyword(reader Reader, keyword string) (bool, error) {

	d, err := reader.Peek(len(keyword))
	if err != nil {
		return false, err
	}

	if string(d) == keyword {
		_, err = reader.Discard(len(keyword))
		return true, err
	}
	return false, nil
}

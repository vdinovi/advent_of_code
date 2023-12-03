package lib

import "io"

func ReadAllAsBytes(r io.Reader) (buf []byte, err error) {
	return io.ReadAll(r)
}

func ReadAllAsString(r io.Reader) (string, error) {
	buf, err := ReadAllAsBytes(r)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

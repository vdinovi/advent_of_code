package lib

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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

func ParseInts(str string) (ints []int, err error) {
	strs := strings.Fields(str)
	ints = make([]int, len(strs))
	for i, s := range strs {
		if ints[i], err = strconv.Atoi(s); err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func ParseFloats(str string) (floats []float64, err error) {
	strs := strings.Fields(str)
	floats = make([]float64, len(strs))
	for i, s := range strs {
		if floats[i], err = strconv.ParseFloat(s, 64); err != nil {
			return nil, err
		}
	}
	return floats, nil
}

type ScanLinesFunc[T any] func(line string) (t T, ok bool, err error)

func ScanLines[T any](r io.Reader, f ScanLinesFunc[T]) (ts []T, err error) {
	ts = make([]T, 0)
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if line == "" {
			continue
		}
		t, ok, err := f(line)
		if err != nil {
			return nil, err
		}
		if ok {
			ts = append(ts, t)
		}
	}
	return ts, scan.Err()
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[S ~[]T, T Numeric](ns S) (sum T) {
	for _, n := range ns {
		sum += n
	}
	return sum
}

func AllEqual[S ~[]T, T Numeric](ns S) bool {
	if len(ns) == 0 {
		return true
	}
	for _, n := range ns[1:] {
		if n != ns[0] {
			return false
		}
	}
	return true
}

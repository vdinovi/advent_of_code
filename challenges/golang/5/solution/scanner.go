package solution

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type scanner struct {
	scanner *bufio.Scanner
	seq     *sequence
	mapp    *mapping
	err     error
}

func newScanner(r io.Reader) *scanner {
	return &scanner{
		scanner: bufio.NewScanner(r),
	}
}

var seqRe = regexp.MustCompile(`^(\w+):(?:([\s\d]+))`)
var mapRe = regexp.MustCompile(`(\w+)-to-(\w+) map:`)

func (s *scanner) Scan() bool {
	s.seq = nil
	s.mapp = nil
	if s.scanner.Scan() {
		line := s.scanner.Text()
		line = strings.TrimSpace(line)
		match := seqRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 && len(match[0]) > 2 {
			s.seq, s.err = s.scanSequence(match[0][1], match[0][2])
			return s.err == nil
		}
		match = mapRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 && len(match[0]) > 2 {
			s.mapp, s.err = s.scanMapping(match[0][1], match[0][2])
			return s.err == nil
		}
		s.err = fmt.Errorf("invalid line %s", line)
		return false
	}
	s.err = s.scanner.Err()
	return false
}

func (s *scanner) Sequence() *sequence {
	return s.seq
}

func (s *scanner) Mapping() *mapping {
	return s.mapp
}

func (s *scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

func (s *scanner) scanSequence(name, values string) (*sequence, error) {
	vals := strings.Fields(values)
	seq := &sequence{
		name:  name,
		elems: make([]int, len(vals)),
	}
	for i, s := range vals {
		n, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		seq.elems[i] = int(n)
	}
	return seq, s.scanBlankLine()
}

var rangeRe = regexp.MustCompile(`^(\d+)\s+(\d+)\s+(\d+)`)

func (s *scanner) scanMapping(source, destination string) (*mapping, error) {
	mapp := &mapping{
		source:      source,
		destination: destination,
		trans:       make([][3]int, 0),
	}
	for s.scanner.Scan() {
		line := s.scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			return mapp, s.scanner.Err()
		}
		match := rangeRe.FindAllStringSubmatch(line, -1)
		if len(match) < 1 && len(match[0]) != 4 {
			return nil, fmt.Errorf("invalid line %s", line)
		}
		trans := [3]int{}
		for i, s := range match[0][1:] {
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, err
			}
			trans[i] = int(n)
		}
		mapp.trans = append(mapp.trans, trans)
	}
	return mapp, s.scanner.Err()
}

func (s *scanner) scanBlankLine() error {
	if s.scanner.Scan() {
		line := s.scanner.Text()
		line = strings.TrimSpace(line)
		if line != "" {
			return fmt.Errorf("expected blank line but got %s", line)
		}
	}
	return s.scanner.Err()
}

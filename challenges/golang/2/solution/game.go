package solution

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type color int

const (
	Red = iota
	Blue
	Green
)

func (r *Record) Possible(cubes [3]int) bool {
	for _, note := range r.Notes {
		for i, n := range note {
			if n > cubes[i] {
				return false
			}
		}
	}
	return true
}

func (r *Record) Fewest() (cubes [3]int) {
	for _, note := range r.Notes {
		for i, n := range note {
			cubes[i] = max(cubes[i], n)
		}
	}
	return cubes
}

type Record struct {
	Game  int
	Notes [][3]int
}

func Parse(r io.Reader) ([]Record, error) {
	records := []Record{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		r, err := parseRecord(line)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return records, nil
}

var gameIdRegex = regexp.MustCompile(`Game (\d+):`)
var cubeCountRegex = regexp.MustCompile(`(\d+)\s+(\w+)([,;])?`)

func parseRecord(line string) (r Record, err error) {
	match := gameIdRegex.FindSubmatch([]byte(line))
	if match == nil || len(match) < 2 {
		return r, fmt.Errorf("failed to parse game id from line %s", line)
	}
	id, err := strconv.ParseInt(string(match[1]), 10, 64)
	if err != nil {
		return r, err
	}
	r.Game = int(id)
	matches := cubeCountRegex.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return r, fmt.Errorf("failed to parse info from line %s", line)
	}
	var notes [][3]int
	note := [3]int{}
	for _, match := range matches {
		if len(match) < 4 {
			return r, fmt.Errorf("failed to parse info from line %s", line)
		}
		col, err := getColor(match[2])
		if err != nil {
			return r, err
		}
		count, err := strconv.ParseInt(string(match[1]), 10, 64)
		if err != nil {
			return r, err
		}
		note[col] += int(count)
		switch string(match[3]) {
		case ",":
			continue
		case ";", "":
			notes = append(notes, note)
			note = [3]int{}
		default:
			return r, fmt.Errorf("failed to parse line %s", line)
		}
	}
	r.Notes = notes
	return r, nil
}

func getColor(c string) (color, error) {
	c = strings.TrimSpace(c)
	c = strings.ToLower(c)
	switch c {
	case "red":
		return Red, nil
	case "blue":
		return Blue, nil
	case "green":
		return Green, nil
	default:
		return 0, fmt.Errorf("unknown color %q", c)
	}
}

package internal_test

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/vdinovi/advent_of_code_2023/tool/internal"
	"golang.org/x/net/html"
)

var (
	authPrompt      *regexp.Regexp = regexp.MustCompile(`To play,[\w\-\s]+`)
	answerPrompt    *regexp.Regexp = regexp.MustCompile(`To begin,[\w\-\s]+`)
	completedPrompt *regexp.Regexp = regexp.MustCompile(`Your puzzle answer was[\w\-\s]+`)
)

const (
	uri = "https://adventofcode.com/2022/day/1"
)

func TestDownloadWithoutSessionToken(t *testing.T) {
	buf := &bytes.Buffer{}
	err := internal.Download(buf, uri, "")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	doc, err := html.Parse(buf)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	var target *html.Node
	var search func(n *html.Node)
	search = func(n *html.Node) {
		if authPrompt.Match([]byte(n.Data)) {
			target = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			search(c)
		}
	}
	search(doc)
	if target == nil {
		t.Fatalf("Failed to locate auth prompt")
	}
}

func TestDownloadWithSessionToken(t *testing.T) {
	session, err := internal.GetSessionToken()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	buf := &bytes.Buffer{}
	err = internal.Download(buf, uri, session)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	doc, err := html.Parse(buf)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	var target *html.Node
	var search func(n *html.Node)
	search = func(n *html.Node) {
		if answerPrompt.Match([]byte(n.Data)) || completedPrompt.Match([]byte(n.Data)) {
			target = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			search(c)
		}
	}
	search(doc)
	if target == nil {
		t.Fatalf("Failed to locate answer or completed prompt")
	}
}

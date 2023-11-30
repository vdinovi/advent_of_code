package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	year       = 2023
	month      = time.December
	endOfMonth = 31
)

type language string

const (
	golang language = "golang"
)

var languages = map[string]language{
	"golang": golang,
}

func main() {
	var (
		err          error
		date         time.Time
		templatesDir string
		outputDir    string
		language     language
	)
	lang := flag.String("language", "", "specify choice of language")
	day := flag.Int("day", 0, "specify the day of month")
	templates := flag.String("templates", "", "specify path to templates directory")
	output := flag.String("output", "", "specify path to output directory")
	flag.Parse()

	if language, err = parseLanguage("language", *lang); err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}

	if date, err = parseDate("day", *day); err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}

	if templatesDir, err = parseExistingDir("templates", *templates); err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}

	if outputDir, err = parseExistingDir("output", *output); err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}

	challengeDir := filepath.Join(outputDir, fmt.Sprint(date.Day()))
	if err := generate(language, date, challengeDir, templatesDir); err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}
}

func parseExistingDir(name string, arg string) (string, error) {
	if arg == "" {
		return "", &invalidArgumentError{
			name:   name,
			value:  arg,
			reason: "empty",
		}
	}
	path := filepath.Clean(arg)
	info, err := os.Stat(path)
	if err != nil {
		return "", &invalidArgumentError{
			name:   name,
			value:  arg,
			reason: err.Error(),
		}
	}
	if !info.IsDir() {
		return "", &invalidArgumentError{
			name:   name,
			value:  arg,
			reason: "not a directory",
		}
	}
	return path, nil
}

func parseDate(name string, arg int) (date time.Time, err error) {
	if arg == 0 {
		return date, &invalidArgumentError{
			name:   name,
			value:  fmt.Sprint(arg),
			reason: "missing",
		}
	}
	if arg > endOfMonth {
		return date, &invalidArgumentError{
			name:   name,
			value:  fmt.Sprint(arg),
			reason: "beyond end of month",
		}
	}
	date = time.Date(year, month, arg, 0, 0, 0, 0, time.UTC)
	return date, nil
}

func parseLanguage(name string, arg string) (language, error) {
	lang, ok := languages[arg]
	if !ok {
		return language(""), &invalidArgumentError{
			name:   name,
			value:  fmt.Sprint(arg),
			reason: fmt.Sprintf("invalid language"),
		}
	}
	return lang, nil
}

type invalidArgumentError struct {
	name   string
	value  string
	reason string
}

func (e *invalidArgumentError) Error() string {
	return fmt.Sprintf("argument %s=%q is invalid: %s", e.name, e.value, e.reason)
}

func generate(language language, date time.Time, outputDir string, templatesDir string) error {
	fmt.Printf("~> generating %d/%d/%d in %s: %s (from %s)\n",
		date.Month(), date.Day(), date.Year(), string(language), outputDir, templatesDir)
	return nil
}

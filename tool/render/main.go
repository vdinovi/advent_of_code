package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vdinovi/advent_of_code_2023/tool/internal"
)

func main() {
	var (
		err          error
		date         time.Time
		templatesDir string
		outputDir    string
		language     internal.Language
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
	err = render(language, date, challengeDir, templatesDir)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n", err)
		os.Exit(1)
	}
}

const (
	year       = 2023
	month      = time.December
	endOfMonth = 31
)

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/vdinovi/advent_of_code/tool/internal"
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
	rm := flag.Bool("rm", false, "remove existing files")
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

	err = render(language, date, outputDir, templatesDir, *rm)
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

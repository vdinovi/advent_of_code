package main

import (
	"fmt"
	"time"

	"github.com/vdinovi/advent_of_code_2023/tool/internal"
)

func render(language internal.Language, date time.Time, outputDir string, templatesDir string) error {
	fmt.Printf("~> generating %d/%d/%d in %s: %s (from %s)\n",
		date.Month(), date.Day(), date.Year(), language.String(), outputDir, templatesDir)
	return nil
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/vdinovi/advent_of_code/tool/internal"
)

type invalidArgumentError struct {
	name   string
	value  string
	reason string
}

func (e *invalidArgumentError) Error() string {
	return fmt.Sprintf("argument %s=%q is invalid: %s", e.name, e.value, e.reason)
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

func parseLanguage(name string, arg string) (internal.Language, error) {
	lang, err := internal.GetLanguageByName(arg)
	if err != nil {
		return lang, &invalidArgumentError{
			name:   name,
			value:  fmt.Sprint(arg),
			reason: err.Error(),
		}
	}
	return lang, nil
}

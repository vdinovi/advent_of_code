package internal

import (
	"fmt"
	"strings"
	"unicode"
)

func GetLanguageByName(name string) (lang Language, err error) {
	name = makeValidLanguageName(name)
	for i, n := range LanguageNames {
		if n == name {
			return Languages[i], nil
		}
	}
	return lang, &unknownLanguageError{name: name}
}

type Language struct {
	Name         string
	LibName      string
	TemplateName string
}

func newLanguage(name, libName, templateName string) Language {
	return Language{
		Name:         makeValidLanguageName(name),
		LibName:      libName,
		TemplateName: templateName,
	}
}

func (l Language) String() string {
	return l.Name
}

var (
	Languages = [...]Language{
		newLanguage("Golang", "golang", "golang"),
		newLanguage("Python", "python", "python"),
	}
	LanguageNames []string
)

func init() {
	LanguageNames = make([]string, len(Languages))
	for i, lang := range Languages {
		LanguageNames[i] = lang.Name
	}
}

type unknownLanguageError struct {
	name string
}

func (e *unknownLanguageError) Error() string {
	return fmt.Sprintf("unknown language %s (choices: %s)",
		e.name, strings.Join(LanguageNames, ", "))
}

// sOmeLanGuagE -> Somelanguage
func makeValidLanguageName(s string) string {
	var i int
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		if i == 0 {
			r = unicode.ToUpper(r)
		} else {
			r = unicode.ToLower(r)
		}
		i += 1
		return r
	}, s)
}

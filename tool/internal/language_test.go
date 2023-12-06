package internal_test

import (
	"testing"

	"github.com/vdinovi/advent_of_code/tool/internal"
)

func TestGetLanguageByName(t *testing.T) {
	golang := internal.Languages[0]
	fnName := "GetLanguageByName"
	tests := []struct {
		name      string
		lang      internal.Language
		shouldErr bool
	}{
		{name: "Golang", lang: golang},
		{name: " goLanG ", lang: golang},
		{name: "Haskell", shouldErr: true},
	}
	for _, test := range tests {
		lang, err := internal.GetLanguageByName(test.name)
		if test.shouldErr {
			if err == nil {
				t.Errorf("expected %s(%q) to yield error but did not", fnName, test.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("unexpected error in %s(%q): %s", fnName, test.name, err)
			continue
		}
		if lang != test.lang {
			t.Errorf("expected %s(%q) to yield %s, but got %s",
				test.name, test.lang.String(), test.lang.String(), err)
		}
	}
}

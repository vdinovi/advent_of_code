package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/vdinovi/advent_of_code_2023/tool/internal"
)

type challenge struct {
	lang         internal.Language
	date         time.Time
	outputDir    string
	templatesDir string
}

const (
	templateExt = ".templ"
	dirMode     = 0755
	fileMode    = 0644
)

func render(
	language internal.Language,
	date time.Time,
	outputDir string,
	templatesDir string,
	remove bool) error {
	fmt.Printf("-> rendering %d/%d/%d in %s: %s (from %s)\n",
		date.Month(), date.Day(), date.Year(), language.String(), outputDir, templatesDir)
	ch := &challenge{
		lang:         language,
		date:         date,
		outputDir:    outputDir,
		templatesDir: templatesDir,
	}
	templates := filepath.Join(ch.templatesDir, ch.lang.TemplateName)
	fmt.Printf("-> stat %s\n", templates)
	if info, err := os.Stat(templates); err != nil {
		return err
	} else if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", templates)
	}
	target := filepath.Join(outputDir, ch.lang.TemplateName, fmt.Sprint(date.Day()))
	if remove {
		fmt.Printf("-> rm -r %s\n", target)
		if err := os.RemoveAll(target); err != nil {
			return err
		}
	}
	fmt.Printf("-> stat %s\n", target)
	if _, err := os.Stat(target); err == nil {
		return fmt.Errorf("target %s already exists", target)
	} else if !os.IsNotExist(err) {
		return err
	}
	fmt.Printf("-> mkdir %s\n", target)
	if err := os.MkdirAll(target, dirMode); err != nil {
		return err
	}
	abort := func() {
		fmt.Println("aborting...")
		fmt.Printf("-> rm -r %s\n", target)
		if err := os.RemoveAll(target); err != nil {
			panic(err)
		}
	}
	err := fs.WalkDir(os.DirFS(templates), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		dest := filepath.Join(target, path)
		if d.IsDir() {
			fmt.Printf("-> mkdir %s\n", dest)
			return os.Mkdir(dest, dirMode)
		}
		src := filepath.Join(templates, path)
		switch filepath.Ext(path) {
		case templateExt:
			dest = strings.TrimSuffix(dest, templateExt)
			return render_template(ch, src, dest)
		default:
			return copy_file(ch, src, dest)
		}
	})
	if err != nil {
		abort()
	}
	return err
}

type data struct {
	Day     int
	LibName string
}

func render_template(ch *challenge, source, dest string) (err error) {
	fmt.Printf("-> render %s %s\n", source, dest)

	templ := template.New(source)
	buf, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	templ, err = templ.Parse(string(buf))
	if err != nil {
		return err
	}

	fdest, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileMode)
	if err != nil {
		return err
	}
	defer fdest.Close()

	return templ.Execute(fdest, data{
		Day:     ch.date.Day(),
		LibName: ch.lang.LibName,
	})
}

func copy_file(ch *challenge, source, dest string) (err error) {
	fmt.Printf("-> cp %s %s\n", source, dest)
	fsource, err := os.Open(source)
	if err != nil {
		return err
	}
	defer fsource.Close()

	fdest, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileMode)
	if err != nil {
		return err
	}
	defer fdest.Close()

	_, err = io.Copy(fdest, fsource)
	return err
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Locale stores map from language ID to Language object
type Locale struct {
	Languages map[string]Language
}

// Language stores map of key to text body
type Language struct {
	Strings map[string]string
}

// GetLangString returns a text body based on language ID and key.
func (l Locale) GetLangString(lang string, key string, vargs ...interface{}) string {
	str := fmt.Sprintf(l.Languages[lang].Strings[key], vargs...)
	if str == "" {
		log.Printf("ERROR: undefined lang key: '%s'", key)
	}
	return str
}

func (app *App) loadLanguages() {
	files, err := ioutil.ReadDir("lang")
	if err != nil {
		log.Fatal(err)
	}

	var key string
	app.locale.Languages = make(map[string]Language)

	for _, f := range files {
		if f.IsDir() {
			key = f.Name()
			la := loadLanguageFromDir(filepath.Join("lang", key))
			app.locale.Languages[key] = la
		}
	}
}

func loadLanguageFromDir(dir string) Language {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var language Language
	var key string

	language = Language{
		Strings: make(map[string]string),
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		key = f.Name()
		str := loadLanguageStringFromFile(filepath.Join(dir, key))
		language.Strings[key] = str
	}

	return language
}

func loadLanguageStringFromFile(file string) string {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(contents)
}

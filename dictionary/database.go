package dictionary

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"regexp"
	"strings"
)

type dictDB struct {
	erdb *sql.DB
	redb *sql.DB
}

func GetDatabases() dictDB {
	erdb, err := sql.Open("sqlite3", "./dictionary/pkg/english-russian.sqlite3")
	if err != nil {
		return dictDB{}
	}
	redb, err := sql.Open("sqlite3", "./dictionary/pkg/russian-english.db")
	if err != nil {
		return dictDB{}
	}
	return dictDB{erdb: erdb, redb: redb}
}

func (db dictDB) translateRusToEn(word string) []string {
	translations := make([]string, 0)
	if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(word) == false {
		return translations
	}
	rows, err := db.redb.Query("select name from word where id in (select idTranslation from translation where idWord in (select id from word where word.name==$1))", word)
	if err != nil {
		//println(err)
		return translations
	}
	if rows == nil {
		return translations
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		rows.Scan(&name)
		translations = append(translations, name)
	}
	return translations
}

func (db dictDB) translateEnToRus(word string) string {
	if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(word) == false {
		return ""
	}
	table := "Caches" + strings.ToUpper(string(word[0]))
	translation := db.erdb.QueryRow("SELECT translationShort FROM "+table+" where word=$1", word)
	var translationShort string
	translation.Scan(&translationShort)
	return translationShort
}

func (db dictDB) GetRusToEnTranslation(word string) []string {
	return db.translateRusToEn(word)
}

func (db dictDB) GetEnToRusTranslation(word string) []string {
	translationShort := db.translateEnToRus(word)
	translationArr := make([]string, 0)
	if len(translationShort) == 0 {
		return translationArr
	}
	translationShort = regexp.MustCompile(`\([^()]*\)`).ReplaceAllString(translationShort, "")
	translations := regexp.MustCompile(`\d+\)`).Split(translationShort, -1)
	for _, item := range translations {
		parts := regexp.MustCompile(`[а-яё .]+`).FindAllString(item, -1)
		for _, p := range parts {
			p = regexp.MustCompile(`\s+[^а-яё]`).ReplaceAllString(p, "")
			if len(p) > 2 && string(p[0]) == " " {
				p = p[1:]
			}
			runes := []rune(p)
			if len(runes) > 0 {
				if runes[len(runes)-1] == ' ' {
					runes = runes[:len(runes)-1]
					p = p[:len(p)-1]
				}
				if len([]rune(p)) >= len(word) {
					translationArr = append(translationArr, p)
				}
			}
		}
	}
	return translationArr
}

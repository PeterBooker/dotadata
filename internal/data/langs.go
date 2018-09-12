package data

// Langs ...
type Langs []Lang

// Lang contains the language code and name
type Lang struct {
	Code, Name string
}

// Languages supported by Dota 2
var langs = &Langs{
	{"pt-br", "brazilian"},
	{"bg", "bulgarian"},
	{"cz", "czech"},
	{"da", "danish"},
	{"nl", "dutch"},
	{"en", "english"},
	{"fi", "finnish"},
	{"de", "german"},
	{"el", "greek"},
	{"hu", "hungarian"},
	{"it", "italian"},
	{"ja", "japanese"},
	{"ko", "korean"},
	{"no", "norwegian"},
	{"pl", "polish"},
	{"pt", "portuguese"},
	{"ro", "romainian"},
	{"ru", "russian"},
	{"es", "spanish"},
	{"th", "thai"},
	{"tr", "turkish"},
	{"uk", "ukranian"},
	{"zh-sg", "schinese"},
	{"zh-tw", "tchinese"},
}

// GetLangs returns a list of Dota 2 supported languages.
func GetLangs() *Langs {
	return langs
	/*
		return &Langs{
			{"cz", "czech"},
			{"en", "english"},
			{"de", "german"},
		}
	*/
}

// ValidateLang checks if the language is supported by Dota 2.
// Resorts to default if not supported.
func ValidateLang(lang string) string {
	for _, v := range *langs {
		if lang == v.Code || lang == v.Name {
			return v.Name
		}
	}

	return "english"
}

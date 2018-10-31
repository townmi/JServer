package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle i18n.Bundle

func init() {
	bundle = i18n.Bundle{DefaultLanguage: language.English}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	bundle.MustLoadMessageFile(PWD + "/language/app.en.toml")
	bundle.MustLoadMessageFile(PWD + "/language/app.zh.toml")
}

// GetLocalizer s
func GetLocalizer(lang string, accept string) *i18n.Localizer {
	return i18n.NewLocalizer(&bundle, lang, accept)
}

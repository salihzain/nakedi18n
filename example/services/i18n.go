package services

import "nakedi18n"

var NakedI18n *nakedi18n.NakedI18n

func init() {
	globalMessages := map[string]map[string]string{
		"en": {
			"welcome": "Welcome!",
		},
		"ar": {
			"welcome": "مرحبا",
		},
	}

	NakedI18n = nakedi18n.NewNakedI18n("en", []string{"en", "ar"}, true, globalMessages)
}

package accounting

import (
	"errors"
	"nakedi18n/example/services"
	"strings"
)

var (
	localMessages = map[string]map[string]string{
		"en": {
			"errInvalidUsername": "Invalid username: '%s'. The username must be a nonempty string containing no spaces",
			"errEmptyName":       "Invalid empty name. The name must be a nonempty string",
		},
		"ar": {
			"errInvalidUsername": "خطأ في اسم المستخدم: '%s'. يجب ان يكون اسم المستخدم خاليا من الفراغات",
			"errEmptyName":       "خطأ في الاسم. يرجى تسجيل اسم",
		},
	}

	t = services.NakedI18n.UseNakedI18n(localMessages, true)
)

type Account struct {
	Username string
	Name     string
}

func CreateAccount(username, name, lang string) (Account, error) {
	a := Account{}

	if strings.ContainsAny(username, " ") || username == "" {
		return a, errors.New(t(lang, "errInvalidUsername", username))
	}

	if name == "" {
		return a, errors.New(t(lang, "errEmptyName"))
	}

	a.Username = username
	a.Name = name

	return a, nil
}

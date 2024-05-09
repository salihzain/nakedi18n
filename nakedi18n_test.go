package nakedi18n

import (
	"reflect"
	"testing"
)

func TestNewNakedI18n(t *testing.T) {
	defaultLocale := "en"
	locales := []string{"en", "ar"}
	globalMessages := map[string]map[string]string{
		"en": {
			"welcome": "Welcome!",
		},
		"ar": {
			"welcome": "مرحبا",
		},
	}

	nakedI18n := NewNakedI18n(defaultLocale, locales, false, globalMessages)

	if nakedI18n.GetDefaultLocale() != defaultLocale {
		t.Fatalf("Expected default locale to be: '%s' but got '%s'", defaultLocale, nakedI18n.GetDefaultLocale())
	}

	if !reflect.DeepEqual(nakedI18n.GetLocales(), locales) {
		t.Fatalf("Expected locales to be: '%s' but got '%s'", locales, nakedI18n.GetLocales())
	}
}

func TestUseNakedI18n(t *testing.T) {
	defaultLocale := "en"
	locales := []string{"en", "ar"}
	globalMessages := map[string]map[string]string{
		"en": {
			"welcome":     "Welcome!",
			"errNotFound": "Not Found!",
			"success":     "Item has been created successfully",
		},
		"ar": {
			"welcome":     "مرحبا",
			"errNotFound": "تعذر الوجود",
		},
	}

	nakedI18n := NewNakedI18n(defaultLocale, locales, false, globalMessages)

	localize := nakedI18n.UseNakedI18n(
		map[string]map[string]string{
			"en": {
				"errNotFound":    "404 Not Found",
				"greeting":       "Greeting %s",
				"errInvalidData": "Invalid data provided",
			},
			"ar": {
				"errNotFound": "خطأ رقم 404 تعذر الوجود",
				"greeting":    "مرحبا بك %s",
			},
		},
		true)

	// Found in local messages
	expected := "Greeting Salih"
	actual := localize("en", "greeting", "Salih")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	expected = "مرحبا بك صالح"
	actual = localize("ar", "greeting", "صالح")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// Found & overshadowing global messages
	expected = "404 Not Found"
	actual = localize("en", "errNotFound")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// Local messages: Fallback on default language
	expected = "Invalid data provided"
	actual = localize("ar", "errInvalidData")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// Found in global messages
	expected = "Welcome!"
	actual = localize("en", "welcome")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// Found in global messages, fallback
	expected = "Item has been created successfully"
	actual = localize("ar", "success")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// no local messages passed
	localize = nakedI18n.UseNakedI18n(nil, true)

	expected = "Welcome!"
	actual = localize("en", "welcome")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// don't use global scope
	localize = nakedI18n.UseNakedI18n(
		map[string]map[string]string{},
		false)

	expected = "welcome"
	actual = localize("en", "welcome")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}

	// not found
	expected = "not_found"
	actual = localize("en", "not_found")
	if expected != actual {
		t.Fatalf("Expected: '%s', got: '%s'", expected, actual)
	}
}

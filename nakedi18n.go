package nakedi18n

import (
	"fmt"
	"log"
)

type NakedI18n struct {
	defaultLocale     string
	locales           []string
	globalMessages    map[string]map[string]string
	enableWarningLogs bool
}

func NewNakedI18n(defaultLocale string, locales []string, enableWarningLogs bool, globalMessages map[string]map[string]string) *NakedI18n {
	if defaultLocale == "" && enableWarningLogs {
		log.Println("WARNING: nakedi18n: defaultLocale is empty")
	}

	if locales == nil || len(locales) == 0 && enableWarningLogs {
		log.Println("WARNING: nakedi18n: locales are empty")
	}

	if globalMessages == nil {
		globalMessages = map[string]map[string]string{}
	}

	return &NakedI18n{
		defaultLocale:  defaultLocale,
		locales:        locales,
		globalMessages: globalMessages,
	}
}

func (n *NakedI18n) GetLocales() []string {
	return n.locales
}

func (n *NakedI18n) GetDefaultLocale() string {
	return n.defaultLocale
}

func (n *NakedI18n) UseNakedI18n(localMessages map[string]map[string]string, useGlobalScope bool) func(string, string, ...any) string {
	for _, l := range n.locales {
		if _, found := localMessages[l]; !found && n.enableWarningLogs {
			log.Printf("WARNING: nakedi18n: no messages found for locale: %s in local messages\n", l)
		}
	}

	if localMessages == nil {
		localMessages = map[string]map[string]string{}
	}

	return func(lang, key string, args ...any) string {
		tryLocal := localMessages[lang][key]

		if tryLocal != "" {
			return fmt.Sprintf(tryLocal, args...)
		}

		if n.enableWarningLogs {
			log.Printf("WARNING: nakedi18n: key: %s is not localized for lang: %s in local messages\n", key, lang)
		}

		if useGlobalScope {
			tryGlobal := n.globalMessages[lang][key]

			if tryGlobal != "" {
				return fmt.Sprintf(tryGlobal, args...)
			}

			if n.enableWarningLogs {
				log.Printf("WARNING: nakedi18n: key: %s is not localized for lang: %s in global messages\n", key, lang)
			}
		}

		// fallback on default locale
		if lang != n.defaultLocale {
			tryFallbackLocale := localMessages[n.defaultLocale][key]

			if tryFallbackLocale != "" {
				return fmt.Sprintf(tryFallbackLocale, args...)
			}

			if n.enableWarningLogs {
				log.Printf("WARNING: nakedi18n: key: %s is not localized for lang: %s in the fallback locale of local messages\n", key, lang)
			}

			if useGlobalScope {
				tryFallbackLocale := n.globalMessages[n.defaultLocale][key]

				if tryFallbackLocale != "" {
					return fmt.Sprintf(tryFallbackLocale, args...)
				}

				if n.enableWarningLogs {
					log.Printf("WARNING: nakedi18n: key: %s is not localized for lang: %s in the fallback locale of global messages\n", key, lang)
				}
			}
		}

		return key
	}
}

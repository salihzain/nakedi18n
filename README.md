# nakedi18n 
A simple, naive, painless, composable i18n golang package. 


### Why this simple? 
After coding in golang for almost 5 years, I have come to realize that I don't care about: 
1. Templating (i.e. `one`, `many`, `other`). 
2. Defining all messages in one place. 
3. Using `json` or `toml` files to define messages. 

I have come to realize that all of the above, while results in fantastic & robust error messages if done correctly,
it adds a lot of overhead that prevented me and other devs on my team from actually writing useful messages. 

On the other hand, it's been a breeze writing error messages in `JavaScript`, thanks to the amazing `vue-i18n`.
Hence, the inspiration to use currying to allow defining local messages from any package, while still having access to global messages. 

### Install
```sh
go get github.com/salihzain/nakedi18n
```

### Usage 

1. Setup your `NakedI18n` instance to be used project wide with your preferred config
   (look at the example project. In my case, I like to place it in services/i18n)
```go
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
```

2. Use from your local package
```go

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

....
	
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
```
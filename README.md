A simple GoLang Library that can be used to detect user current language on windows/darwin/linux

## Usage
1. Add this line to your go import:

```golang
	"github.com/wangyumu/default_lang"
```

2. `DetectLang` will return the user default language, two two lowercase letters like: en, zh, see [ISO 639](http://en.wikipedia.org/wiki/ISO_639) 

```golang
langs, err := default_lang.DetectLang()
println("langs:", langs)
```

3. `DetectLocale` will return the user default locale, there are two formats:
- [language designator]_[region designator] , like: zh_CN 
- [language designator]-[script designator]_[region designator], like: zh-Hans_CN 

```golang
locales, err := default_lang.DetectLocale()
println("locales:", locales)
if err != nil {
    matcher := language.NewMatcher([]language.Tag{language.English, language.Chinese})
    tag, _ := language.MatchStrings(matcher, locales...)
}
```

## How
1. windows via sys library 
2. *nix via environment variables
3. Darwin via `defaults read -g AppleLanguages` as GUI Application on Mac OS can not get LANG from os.Getenv()
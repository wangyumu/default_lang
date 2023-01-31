package default_lang

// DetectLang return user default languages, like zh, en, ja (ISO 639-1)
func DetectLang() ([]string, error) {
	locales, err := getDefaultLocales()
	if err != nil {
		return nil, err
	}
	for i, _ := range locales {
		locales[i] = locales[i][0:2] // first two lowercase letters languages (ISO 639-1)
	}
	return locales, nil
}

// DetectLocale return user default locales, like: zh_CN or zh-Hans_CN (darwin)
func DetectLocale() ([]string, error) {
	return getDefaultLocales()
}

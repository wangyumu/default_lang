//go:build windows && go1.9
// +build windows,go1.9

package default_lang

import (
	"golang.org/x/sys/windows"
)

func getDefaultLocales() (langs []string, err error) {
	return windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
}

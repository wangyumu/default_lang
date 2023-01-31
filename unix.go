//go:build !windows
// +build !windows

package default_lang

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getDefaultLocales() ([]string, error) {
	locales := []string{}
	for _, env := range []string{"LC_MESSAGES", "LC_ALL", "LANG"} {
		t := os.Getenv(env)
		if len(t) > 0 {
			locales = append(locales, t[:strings.LastIndex(t, ".")]) // en_US.UTF-8
		}
	}

	if len(locales) < 1 && runtime.GOOS == "darwin" { //for Mac OSX GUI application
		out, err := exec.Command("defaults", "read", "-g", "AppleLanguages").Output() // zh-Hans-CN
		if err != nil {
			return nil, err
		}
		f := []byte{}
		for _, b := range out {
			switch b {
			case '(', ')', '"', ' ', '\n':
				continue
			default:
				f = append(f, b)
			}
		}

		for _, t := range strings.Split(string(f), ",") {
			if len(t) > 0 {
				locales = append(locales, t)
			}
		}
	}
	return locales, nil
}

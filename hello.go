package hello

import "errors"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is required")
	}
	return "Hello, " + name + "!", nil
}

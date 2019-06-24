package main

import (
	"testing"
)

func TestExtractURL(t *testing.T) {
	message := "https://github.com/n-inja/goq\nテスト\nhttps://github.com/n-inja/gomniauth-traq"
	URLs := extractURL(message)
	expected := []string{"https://github.com/n-inja/goq", "https://github.com/n-inja/gomniauth-traq"}
	if len(URLs) != len(expected) {
		t.Fatalf("expecting %v got %v", expected, URLs)
	}
	for i := range URLs {
		if URLs[i] != expected[i] {
			t.Fatalf("expecting %v got %v", expected, URLs)
		}
	}
}

func TestGetOGP(t *testing.T) {
	URL := "https://github.com/n-inja/goq"
	expected := `n-inja/goq
> traQのOGP bot. Contribute to n-inja/goq development by creating an account on GitHub.


`
	message := getOGP(URL)
	if message != expected {
		t.Fatalf("expecting %v got %v", expected, message)
	}
}

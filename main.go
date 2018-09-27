package main

import (
	"github.com/clintonpuge/site-checker/app"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https:www.siteisdowndsad.com",
	}

	app.Checker(sites)
}

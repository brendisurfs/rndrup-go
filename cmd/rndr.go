package cmd

import (
	"log"
	"time"

	"github.com/mxschmitt/playwright-go"
)

// Send func: automates upload RNDR proc.
func Send(email, password string) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not launch playwright: %v", err)
	}

	// start browser: HEADLESS is a necessary option for later

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch firefox: %v", err)
	}
	// create a new page for the browser to go to
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	// visit site and wait until network is idle.
	if _, err = page.Goto("https://rndr.otoy.com", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	// pull out elementsfor input of password
	emailElement, err := page.QuerySelector("#session_email")
	if err != nil {
		log.Fatal("error with email")
	}
	passwordElement, err := page.QuerySelector("#session_password")
	if err != nil {
		log.Fatalf("could not get value: %v", err)
	}

	// fill in email and password
	log.Println("putting in email...")
	err = emailElement.Fill(email)
	if err != nil {
		log.Fatalf("error inputting email")
	}
	log.Println("email successful")

	log.Println("putting in password...")
	err = passwordElement.Fill(password)
	if err != nil {
		log.Fatalf("error inputting password")
	}
	log.Println("password input successful")

	// click on login
	submitButton, err := page.QuerySelector("input[name='commit']")
	if err != nil {
		log.Fatalf("error clicking submit button")
	}
	if err := submitButton.Click(); err != nil {
		log.Fatalf("could not click")
	}

	// check if login works

	// file upload: find upload area, upload to rndr.
	// -- this should spawn on desired number of threads, for multiple processes.
	// -- option is for varying internet speeds and ram size.
	time.Sleep(2 * time.Second)
}

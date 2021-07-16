package cmd

import (
	"fmt"
	"log"

	"github.com/mxschmitt/playwright-go"
)

// Send func: automates upload RNDR proc.
func Send() {
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

	pageName, err := page.QuerySelector("#session_password")
	log.Printf("%v", pageName)
	if err != nil {
		log.Fatalf("could not get value: %v", err)
	}

	fmt.Printf("lot of stuff by: %v", pageName)

	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)

		if err = browser.Close(); err != nil {
			log.Fatalf("could not close browser: %v", err)
		}

		if err = pw.Stop(); err != nil {
			log.Fatalf("could not stop playwright: %v", err)

		}
	}
}

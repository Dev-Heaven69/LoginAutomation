package helper

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func GmailLoginWOOTP(page playwright.Page, email string, password string) error {

	if _, err := page.Goto("https://accounts.google.com/AddSession?service=accountsettings&continue=https://myaccount.google.com/&ec=GAlAwAE&hl=en_GB&authuser=0"); err != nil {
		log.Printf("Error navigating to Google sign-in page: %v", err)
		return err
	}

	if err := page.Locator("input[type='email']").Type(email); err != nil {
		log.Printf("Error typing username: %v", err)
		return err
	}

	if err := page.Locator("#identifierNext > div > button > span").Click(); err != nil {
		log.Printf("Error clicking Next after username: %v", err)
		return err
	}

	if err := page.Locator("#password > div.aCsJod.oJeWuf > div > div.Xb9hP > input").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for password input: %v", err)
		return err
	}

	if err := page.Locator("#password > div.aCsJod.oJeWuf > div > div.Xb9hP > input").Type(password); err != nil {
		log.Printf("Error typing password: %v", err)
		return err
	}

	if err := page.Locator("#passwordNext > div > button > span").Click(); err != nil {
		log.Printf("Error clicking Next after password: %v", err)
		return err
	}
	
	return nil
}

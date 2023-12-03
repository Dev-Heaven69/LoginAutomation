package services

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func Login(page playwright.Page, smartleadEmail string, smartleadPassword string) error{
	smartLeadUser := smartleadEmail
	smartLeadPass := smartleadPassword

	if _, err := page.Goto("https://app.smartlead.ai/login?_gl=1*oz75vx*_gcl_aw*R0NMLjE2OTYzMzI4OTUuRUFJYUlRb2JDaE1JbDZmODN1VFpnUU1WM3lTREF4M3QtUTJjRUFBWUFTQUFFZ0tYbGZEX0J3RQ..*_gcl_au*MTg2MDQ3NjcwMS4xNjk2MzMyODk1"); err != nil {
		log.Printf("Error navigating to smartlead sign-in page: %v", err)
		return err
	}

	fmt.Println("LOGINING IN")

	if err := page.Locator("input[type='email']").Type(smartLeadUser); err != nil {
		log.Printf("Error while Typing email: %v", err)
		return err
	}

	if err := page.Locator("input[type='password']").Type(smartLeadPass); err != nil {
		log.Printf("Error while Typing password: %v", err)
		return err
	}

	if err := page.Locator("#q-app > div > div > main > section > div > form > button > span.q-btn__content.text-center.col.items-center.q-anchor--skip.justify-center.row > span").Click(); err != nil {
		log.Printf("Error while clicking next: %v", err)
		return err
	}

	fmt.Println("LOGGED IN")

	return nil
}
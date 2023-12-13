package services

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func ImagePermissions(page playwright.Page,useraccountidx int) error {
	fmt.Println("clicking profile button")
	if err := page.Locator("#yDmH0d > c-wiz > div > div:nth-child(2) > div > c-wiz > c-wiz > div > div.s7iwrf.gMPiLc.Kdcijb > div > div > header > div.g3lg0e > div > button > figure > img").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for profile button: %v", err)
		return err
	}

	if err := page.Locator("#yDmH0d > c-wiz > div > div:nth-child(2) > div > c-wiz > c-wiz > div > div.s7iwrf.gMPiLc.Kdcijb > div > div > header > div.g3lg0e > div > button > figure > img").Click(); err != nil {
		log.Printf("Error clicking for profile button: %v", err)
		return err
	}
	fmt.Println("clicked profile button")
	
	mainFrame := page.FrameLocator("#yDmH0d > iframe:nth-child(23)")

	fmt.Println("clicking visible to people button")
	if err := mainFrame.Locator("#yDmH0d > c-wiz > main > div > div.DHz5ad > div.cquyXc > div.RtYLze > div > div > button > div.VfPpkd-RLmnJb").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for visible to people button: %v", err)
		return err
	}

	if err := mainFrame.Locator("#yDmH0d > c-wiz > main > div > div.DHz5ad > div.cquyXc > div.RtYLze > div > div > button > div.VfPpkd-RLmnJb").Click(); err != nil {
		log.Printf("Error clicking for visible to people button: %v", err)
		return err
	}
	fmt.Println("clicked visible to people button")

	fmt.Println("clicking everyone button")
	if err := mainFrame.Locator("#ucc-1").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for everyone button: %v", err)
		return err
	}

	if err := mainFrame.Locator("#ucc-1").Click(); err != nil {
		log.Printf("Error clicking for everyone button: %v", err)
		return err
	}
	fmt.Println("clicked everyone button")

	fmt.Println("clicking save button")
	if err := mainFrame.Locator("#yDmH0d > c-wiz > main > div.B3TZB > div:nth-child(2) > button").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for save button: %v", err)
		return err
	}

	if err := mainFrame.Locator("#yDmH0d > c-wiz > main > div.B3TZB > div:nth-child(2) > button").Click(); err != nil {
		log.Printf("Error clicking for save button: %v", err)
		return err
	}
	fmt.Println("clicked save button")

	return nil
}

//#yDmH0d > c-wiz > main > div > div.DHz5ad > div.cquyXc > div.Tz1TEe.nfbOD
// A picture helps people recognize you and lets you know when you’re signed in to your account

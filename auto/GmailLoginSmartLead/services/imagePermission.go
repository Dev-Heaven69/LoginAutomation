package services

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func ImagePermissions(page playwright.Page,useraccountidx int) error {
	fmt.Println("clicking profile button")
	if err := page.Locator("#gb > div.gb_od.gb_id.gb_ud > div.gb_yd.gb_cb.gb_nd.gb_Ad > div.gb_Sd > div.gb_b.gb_v.gb_Zf.gb_H > div").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for profile button: %v", err)
		return err
	}

	if err := page.Locator("#gb > div.gb_od.gb_id.gb_ud > div.gb_yd.gb_cb.gb_nd.gb_Ad > div.gb_Sd > div.gb_b.gb_v.gb_Zf.gb_H > div").Click(); err != nil {
		log.Printf("Error clicking for profile button: %v", err)
		return err
	}
	fmt.Println("clicked profile button")

	fmt.Println("clicking edit button")
	frame := page.FrameLocator("#gb > div.gb_od.gb_id.gb_ud > div.gb_yd.gb_cb.gb_nd.gb_Ad > div:nth-child(3) > iframe")

	if err := frame.Locator("#yDmH0d > c-wiz > div.T4LgNb > div > div > div > div.sZ3gbf > div > div:nth-child(2) > c-wiz > div > div.XS2qof.Q3BXBb > img").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for edit button: %v", err)
		return err
	}

	if err := frame.Locator("#yDmH0d > c-wiz > div.T4LgNb > div > div > div > div.sZ3gbf > div > div:nth-child(2) > c-wiz > div > div.XS2qof.Q3BXBb > img").Click(); err != nil {
		log.Printf("Error clicking for edit button: %v", err)
		return err
	}
	fmt.Println("clicked edit button")

	mainFrame := page.FrameLocator("#yDmH0d > iframe:nth-child(25)")

	fmt.Println("clicking visible to people button")
	if err := mainFrame.Locator("#yDmH0d > c-wiz.zGzCU.ftAIlf.wiax5e.WjBoXb.rqSEqc.SSPGKf.TJKThb.gt34yd > main > div > div.DHz5ad > div.cquyXc > div.RtYLze > div > div > button > span.VfPpkd-kBDsod > svg").First().WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for visible to people button: %v", err)
		return err
	}

	if err := mainFrame.Locator("#yDmH0d > c-wiz.zGzCU.ftAIlf.wiax5e.WjBoXb.rqSEqc.SSPGKf.TJKThb.gt34yd > main > div > div.DHz5ad > div.cquyXc > div.RtYLze > div > div > button > span.VfPpkd-kBDsod > svg").First().Click(); err != nil {
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

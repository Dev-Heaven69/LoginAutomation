package services

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func CheckLessSecureApps(page playwright.Page, accountIdx int) error {

	lessSecureSelector := fmt.Sprintf("https://myaccount.google.com/u/%d/lesssecureapps", accountIdx)
	if _, err := page.Goto(lessSecureSelector); err != nil {
		log.Printf("Error navigating to Google sign-in page: %v", err)
		return err
	}

	if strin, err := page.Locator("#yDmH0d > c-wiz > div > div:nth-child(2) > div:nth-child(2) > c-wiz > div > div.VfPpkd-WsjYwc.VfPpkd-WsjYwc-OWXEXe-INsAgc.KC1dQ.Usd1Ac.AaN0Dd.F2KCCe.Z2xVec.E2bpG.injfOc > div > div > ul > li > div > div.kvjuQc.biRLo > div > button > span").IsHidden(playwright.LocatorIsHiddenOptions{}); strin {
		fmt.Println("cant add this account")
		return err
	}

	if err := page.Locator("#yDmH0d > c-wiz > div > div:nth-child(2) > div:nth-child(2) > c-wiz > div > div.VfPpkd-WsjYwc.VfPpkd-WsjYwc-OWXEXe-INsAgc.KC1dQ.Usd1Ac.AaN0Dd.F2KCCe.Z2xVec.E2bpG.injfOc > div > div > ul > li > div > div.kvjuQc.biRLo > div > button > span").Click(); err != nil {
		log.Printf("Error while clicking the less secure apps setting:  %v", err)
		return err
	}
	return nil
}

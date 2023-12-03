package gmailloginsmartlead

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/seew0/loginAutomation/auto/GmailLoginSmartLead/services"
	"github.com/seew0/loginAutomation/models"
)

func (gls *GmailLoginSmartlead) Switch(page playwright.Page, req models.GmailLoginSmartlead, accountsCount int, userAccountsCount int) error {
	if req.AddSmartlead {
		err := services.GmailLoginSmartLead(page, accountsCount)
		if err != nil {
			fmt.Println(err)
		}
	}

	if req.LessSecure {
		err := services.CheckLessSecureApps(page, userAccountsCount)
		if err != nil {
			fmt.Println(err)
		}
	}

	if req.ImagePermission {
		googleLoginSelector := fmt.Sprintf("https://myaccount.google.com/u/%d/", userAccountsCount)
		if _, err := page.Goto(googleLoginSelector); err != nil {
			log.Printf("Error navigating to smartlead sign-in page: %v", err)
			return err
		}
		time.Sleep(3 * time.Second)
		err := services.ImagePermissions(page, userAccountsCount)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

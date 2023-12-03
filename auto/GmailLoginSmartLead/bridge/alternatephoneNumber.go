package gmailloginsmartlead

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/seew0/loginAutomation/auto/GmailLoginSmartLead/services"
	"github.com/seew0/loginAutomation/auto/GmailLoginSmartLead/services/helper"
	"github.com/seew0/loginAutomation/models"
)

func (gls *GmailLoginSmartlead) AlternatePhoneNumber(page playwright.Page, username string, password string, req models.GmailLoginSmartlead, accountcount int, useraccountCount int) error {
	client := services.MakeClient()
	count := 0

	fmt.Printf("Logging in with %s\n", username)

	err := helper.GmailLoginWOOTP(page, username, password)
	if err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	if url := page.URL(); strings.HasPrefix(url, "https://accounts.google.com/speedbump/gaplustos") {
		if err := page.Locator("#confirm").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
			log.Printf("Error waiting for confirm: %v", err)
			return err
		}

		if err := page.Locator("#confirm").Click(); err != nil {
			log.Printf("Error clicking for confirm: %v", err)
			return err
		}
	}

	time.Sleep(5 * time.Second)
	if url := page.URL(); strings.HasPrefix(url, "https://myaccount.google.com/") {
		err = gls.Switch(page, req, accountcount, useraccountCount)
		if err != nil {
			return err
		}
		return nil
	}

	if err := page.Locator("#deviceAddress").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for phoneNumber: %v", err)
		return err
	}
	//ORDERED SET IN

	order, err := services.OrderActivation(client)
	if err != nil {
		log.Printf("Error while activation: %v", err)
		return err
	}

	phoneNumber := order.Phone
	// PHONE ACQUIRED

	if err := page.Locator("#deviceAddress").Type(phoneNumber); err != nil {
		log.Printf("Error typing phoneNumber: %v", err)
		return err
	}

	if err := page.Locator("#next-button").Click(); err != nil {
		log.Printf("Error clicking Next after phoneNumber: %v", err)
		return err
	}

	//CHECKING FOR ORDER

	checkOTP, err := client.CheckOrder(order.ID)
	if err != nil {
		log.Printf("Error while checking for otp: %v", err)
		return err
	}

	for len(checkOTP.SMS) == 0 {
		if count <= 40 {
			log.Printf("No Messages yet")
			checkOTP, err = client.CheckOrder(checkOTP.ID)
			time.Sleep(3 * time.Second)
			count = count + 1
		} else {
			fmt.Println("canceling order")
			services.CancelOrder(client, checkOTP.ID)
			fmt.Println("making new order")
			checkOTP, err = services.OrderActivation(client)

			phoneNumber = checkOTP.Phone

			err := helper.GmailLoginWOOTP(page, username, password)
			if err != nil {
				return err
			}

			if err := page.Locator("#deviceAddress").Clear(); err != nil {
				log.Printf("Error while clearing field: %v", err)
				return err
			}
			if err := page.Locator("#deviceAddress").Type(phoneNumber); err != nil {
				log.Printf("Error typing phoneNumber: %v", err)
				return err
			}

			if err := page.Locator("#next-button").Click(); err != nil {
				log.Printf("Error clicking Next after phoneNumber: %v", err)
				return err
			}

			count = 0
		}
	}

	otp := checkOTP.SMS[len(checkOTP.SMS)-1].Code

	if err := page.Locator("#smsUserPin").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for PhonePin: %v", err)
		return err
	}

	if err := page.Locator("#smsUserPin").Type(otp); err != nil {
		log.Printf("Error Typing PhonePin: %v", err)
		return err
	}

	if err := page.Locator("#next-button").Click(); err != nil {
		log.Printf("Error clicking Next after PhonePin: %v", err)
		return err
	}

	err = gls.Switch(page, req, accountcount, useraccountCount)
	if err != nil {
		return err
	}

	return nil
}

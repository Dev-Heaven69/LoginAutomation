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

type GmailLoginSmartlead struct {
}

func ProvideGmailLoginSmartlead() GmailLoginSmartlead {
	return GmailLoginSmartlead{}
}

func (gls *GmailLoginSmartlead) Manager(page playwright.Page, req models.GmailLoginSmartlead, accountsCount int, userAccountsCount int, username string, password string) error {
	count := 0
	client := services.MakeClient()

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
		err := gls.Switch(page, req, accountsCount, userAccountsCount)
		if err != nil {
			return err
		}
		return nil
	}

	if err := page.Locator("#phoneNumberId").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
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

	if err := page.Locator("#phoneNumberId").Type(phoneNumber); err != nil {
		log.Printf("Error typing phoneNumber: %v", err)
		return err
	}

	if err := page.Locator("#idvanyphonecollectNext > div > button > span").Click(); err != nil {
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

			if err := page.Locator("#yDmH0d > c-wiz > div > div.eKnrVb > div > div.Z6Ep7d > div > div.XOrBDc > div > div > button > span").Click(); err != nil {
				fmt.Println("Error while clicking try different button")
				return err
			}

			if err := page.Locator("#yDmH0d > c-wiz > div > div.eKnrVb > div > div.j663ec > div > form > span > section:nth-child(2) > div > div > section > div > div > div > ul > li:nth-child(1) > div > div.vxx8jf").WaitFor(); err != nil {
				fmt.Println("Error while waiting for use mobile button")
				return err
			}

			if err := page.Locator("#yDmH0d > c-wiz > div > div.eKnrVb > div > div.j663ec > div > form > span > section:nth-child(2) > div > div > section > div > div > div > ul > li:nth-child(1) > div > div.vxx8jf").Click(); err != nil {
				fmt.Println("Error while clicking for use mobile button")
				return err
			}

			if err := page.Locator("#phoneNumberId").Type(phoneNumber); err != nil {
				log.Printf("Error typing phoneNumber: %v", err)
				return err
			}

			if err := page.Locator("#idvanyphonecollectNext > div > button > span").Click(); err != nil {
				log.Printf("Error clicking Next after phoneNumber: %v", err)
				return err
			}

			count = 0
		}
	}

	otp := checkOTP.SMS[len(checkOTP.SMS)-1].Code

	if err := page.Locator("#idvAnyPhonePin").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error waiting for PhonePin: %v", err)
		return err
	}

	if err := page.Locator("#idvAnyPhonePin").Type(otp); err != nil {
		log.Printf("Error Typing PhonePin: %v", err)
		return err
	}

	if err := page.Locator("#idvanyphoneverifyNext > div > button > span").Click(); err != nil {
		log.Printf("Error clicking Next after PhonePin: %v", err)
		return err
	}

	if err := page.Locator("#confirm").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error clicking: %v", err)
		return err
	}

	if err := page.Locator("#confirm").Click(); err != nil {
		log.Printf("Error clicking: %v", err)
		return err
	}

	time.Sleep(3 * time.Second)

	err = gls.Switch(page, req, accountsCount, userAccountsCount)
	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func GmailLoginSmartLead(page playwright.Page, accountidx int) error {
	fmt.Println("routing to smartleads to add email")

	if _, err := page.Goto("https://app.smartlead.ai/login?_gl=1*oz75vx*_gcl_aw*R0NMLjE2OTYzMzI4OTUuRUFJYUlRb2JDaE1JbDZmODN1VFpnUU1WM3lTREF4M3QtUTJjRUFBWUFTQUFFZ0tYbGZEX0J3RQ..*_gcl_au*MTg2MDQ3NjcwMS4xNjk2MzMyODk1"); err != nil {
		log.Printf("Error navigating to smartlead sign-in page: %v", err)
		return err
	}

	fmt.Println("clicking email address tab")

	if err := page.Locator("#q-app > div > div.q-drawer-container > aside > div > div > a:nth-child(3)").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while waiting for email tag: %v", err)
		return err
	}

	if err := page.Locator("#q-app > div > div.q-drawer-container > aside > div > div > a:nth-child(3)").Click(); err != nil {
		log.Printf("Error while waiting for email tag: %v", err)
		return err
	}

	fmt.Println("clicked email address tab")

	fmt.Println("clicking add account button")

	if err := page.Locator("#projectPage > div > div.email-accounts-page-header > div.email-accounts-page-right-section >  button.q-btn.q-btn-item.non-selectable.no-outline.q-btn--unelevated.q-btn--rectangle.bg-primary.text-white.q-btn--actionable.q-focusable.q-hoverable.q-btn--no-uppercase.q-ml-md > span.q-btn__content.text-center.col.items-center.q-anchor--skip.justify-center.row > span").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while waiting for add email tag: %v", err)
		return err
	}

	if err := page.Locator("#projectPage > div > div.email-accounts-page-header > div.email-accounts-page-right-section > button.q-btn.q-btn-item.non-selectable.no-outline.q-btn--unelevated.q-btn--rectangle.bg-primary.text-white.q-btn--actionable.q-focusable.q-hoverable.q-btn--no-uppercase.q-ml-md > span.q-btn__content.text-center.col.items-center.q-anchor--skip.justify-center.row > span").Click(); err != nil {
		log.Printf("Error while waiting for add email tag: %v", err)
		return err
	}

	fmt.Println("clicked add account button")

	fmt.Println("clicking oauth button")

	if err := page.Locator("#q-portal--dialog--1 > div > div.q-dialog__inner.flex.no-pointer-events.q-dialog__inner--minimized.q-dialog__inner--standard.fixed-full.flex-center > div > div.app-modal-content > div > div:nth-child(2) > div > div:nth-child(1) > div").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while clicking email button: %v", err)
		return err
	}

	if err := page.Locator("#q-portal--dialog--1 > div > div.q-dialog__inner.flex.no-pointer-events.q-dialog__inner--minimized.q-dialog__inner--standard.fixed-full.flex-center > div > div.app-modal-content > div > div:nth-child(2) > div > div:nth-child(1) > div").Click(); err != nil {
		log.Printf("error while clicking: %v", err)
		return err
	}

	fmt.Println("clicked oauth button")

	fmt.Println("clicking add gmail account button")

	if err := page.Locator("#q-portal--dialog--2 > div > div.q-dialog__inner.flex.no-pointer-events.q-dialog__inner--minimized.q-dialog__inner--standard.fixed-full.flex-center > div > div.connect-action-footer > button > span.q-btn__content.text-center.col.items-center.q-anchor--skip.justify-center.row > span").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while waiting for oauth: %v", err)
		return err
	}

	if err := page.Locator("#q-portal--dialog--2 > div > div.q-dialog__inner.flex.no-pointer-events.q-dialog__inner--minimized.q-dialog__inner--standard.fixed-full.flex-center > div > div.connect-action-footer > button > span.q-btn__content.text-center.col.items-center.q-anchor--skip.justify-center.row > span").Click(); err != nil {
		log.Printf("Error while clicking oauth: %v", err)
		return err
	}

	fmt.Println("clicked add gmail account button")

	fmt.Println("clicking google panel")

	var accountSelector string
	fmt.Println(accountidx)

	// if accountidx == 1 {
	// 	accountSelector = "#view_container > div > div > div.pwWryf.bxPAYd > div > div.WEQkZc > div > form > span > section > div > div > div > div > ul > li.JDAKTe.ibdqA.W7Aapd.zpCp3.SmR8 > div"
	// } else {
	accountSelector = fmt.Sprintf("#view_container > div > div > div.pwWryf.bxPAYd > div > div.WEQkZc > div > form > span > section > div > div > div > div > ul > li:nth-child(%d) > div", accountidx)
	// }
	if err := page.Locator(accountSelector).WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while waiting for account: %v", err)
		return err
	}

	if err := page.Locator(accountSelector).Click(); err != nil {
		log.Printf("Error while clicking account: %v", err)
		return err
	}

	fmt.Println("clicked google panel")

	fmt.Println("clicking confirm button")

	if err := page.Locator("#submit_approve_access > div > button > div.VfPpkd-RLmnJb").WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		log.Printf("Error while waiting for confirmation button: %v", err)
		return err
	}

	if err := page.Locator("#submit_approve_access > div > button > div.VfPpkd-RLmnJb").Click(); err != nil {
		log.Printf("Error while clicking confirmation button: %v", err)
		return err
	}

	fmt.Println("clicked confirm button")

	time.Sleep(5 * time.Second)
	return nil
}

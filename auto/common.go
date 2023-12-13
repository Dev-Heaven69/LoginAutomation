package auto

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

type (
	PlaywrightAdapters interface {
		InitiateBrowser(pw *playwright.Playwright) (playwright.Browser, error)
		CreateContext(browser playwright.Browser) (playwright.BrowserContext, error)
	}
	FileAdapters interface {
		ReadCSV(CSVFilePath string) ([][]string, error)
	}
	PageAdapters interface {
		WaitFor(selector string, page playwright.Page) error
		Click(selector string, page playwright.Page) error
		Type(selector string, page playwright.Page,toType string) error
		GOTO(url string, page playwright.Page) error
	}
)

type Common struct {
	PlaywrightAdapters PlaywrightAdapters
	FileAdapters       FileAdapters
	PageAdapters       PageAdapters
}

func ProvideCommon(pla PlaywrightAdapters, fa FileAdapters, pa PageAdapters) Common {
	return Common{
		PlaywrightAdapters: pla,
		PageAdapters:       pa,
		FileAdapters:       fa,
	}
}

func (c Common) WaitFor(selector string, page playwright.Page) error {
	if err := page.Locator(selector).WaitFor(playwright.LocatorWaitForOptions{}); err != nil {
		return err
	}
	return nil
}

func (c Common) Click(selector string, page playwright.Page) error {
	if err := page.Locator(selector).Click(); err != nil {
		return err
	}
	return nil
}

func (c Common) Type(selector string, page playwright.Page, toType string) error {
	if err := page.Locator(selector).Type(toType); err != nil {
		return err
	}
	return nil
}

func (c Common) GOTO(url string, page playwright.Page) error {
	if _, err := page.Goto(url); err != nil {
		return err
	}
	return nil
}

func (c Common) InitiateBrowser(pw *playwright.Playwright) (playwright.Browser, error) {
	browserSource := "chromium"

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless:          playwright.Bool(false),
		Channel:           &browserSource,
		IgnoreDefaultArgs: []string{},
		Args: []string{
			"--start-maximized",
			"--incognito",
		},
	})
	if err != nil {
		log.Fatalf("Error launching Chromium: %v", err)
		return nil, err
	}
	return browser, err
}

func (c Common) CreateContext(browser playwright.Browser) (playwright.BrowserContext, error) {
	colorscheme := "dark"
	TimezoneId := "Asia/Kolkata"
	locale := "en-US"
	useragent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4595.0 Safari/537.36"

	ctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		ColorScheme: (*playwright.ColorScheme)(&colorscheme),
		Viewport:    nil,
		Locale:      &locale,
		UserAgent:   &useragent,
		TimezoneId:  &TimezoneId,
	})
	return ctx, err
}

func (c Common) ReadCSV(CSVFilePath string) ([][]string, error) {
	csvFile, err := os.Open(CSVFilePath)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
		return nil, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	accounts := make([][]string, 0)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		accounts = append(accounts, record)
	}

	return accounts, err
}

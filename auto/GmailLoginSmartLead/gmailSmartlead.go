package gmailloginmaster

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
	"github.com/seew0/loginAutomation/auto"
	gmailloginsmartlead "github.com/seew0/loginAutomation/auto/GmailLoginSmartLead/bridge"
	"github.com/seew0/loginAutomation/auto/GmailLoginSmartLead/services"
	"github.com/seew0/loginAutomation/models"
)

type GmailLoginMaster struct {
	common auto.Common
	gls    gmailloginsmartlead.GmailLoginSmartlead
}

func ProvideGmailLoginMaster(common auto.Common, gls gmailloginsmartlead.GmailLoginSmartlead) GmailLoginMaster {
	return GmailLoginMaster{common: common, gls: gls}
}

const (
	CSVFilePath           = "google_accounts.csv"
	MaxAccountsPerSession = 1
)

func (glm *GmailLoginMaster) GmailLoginManager(req models.GmailLoginSmartlead) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicf("error loading .env file: %v", err)
	}

	startIndex := 0
	accountscount := 1
	userAccountCount := 0

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Error launching Playwright: %v", err)
		return err
	}
	defer pw.Stop()
	// DONE

	// CREATING BROWSER AND CONTEXT
	browser, err := glm.common.InitiateBrowser(pw)
	if err != nil {
		log.Fatalf("Error launching Browser %v", err)
		return err
	}
	fmt.Println("created browser")
	defer browser.Close()

	ctx, err := glm.common.CreateContext(browser)
	if err != nil {
		log.Fatalf("Error creating context: %v", err)
		return err
	}
	fmt.Println("created context")
	defer ctx.Close()
	// DONE

	//CSV READING
	accounts, err := glm.common.ReadCSV(CSVFilePath)
	if err != nil {
		log.Fatalf("Error Failed to read CSV file: %v", err)
		return err
	}

	fmt.Println("csv file loaded")
	//DONE
	numAccounts := len(accounts)

	page, err := ctx.NewPage()
	if err != nil {
		log.Printf("Error creating page: %v", err)
		return err
	}

	fmt.Println("created page")
	
	if req.AddSmartlead {
		err := services.Login(page, string(req.SmartLeadEmail), req.SmartLeadPasswd)
		if err != nil {
			fmt.Println(err)
		}
	}

	for startIndex < numAccounts {
		endIndex := startIndex + MaxAccountsPerSession
		if endIndex > numAccounts {
			endIndex = numAccounts
		}


		time.Sleep(5 * time.Second) //BUFFER TIME

		for i := startIndex; i < endIndex; i++ {
			account := accounts[i]
			username := account[0]
			password := account[1]

			if strings.HasPrefix(password, "\"") && strings.HasSuffix(password, "\"") {
				password = strings.Trim(password, "\"")
			} //filter

			fmt.Println("username: ", username)
			fmt.Println("password: ", password)
			fmt.Println("accountcount ", accountscount)
			fmt.Println("useraccountcount ", userAccountCount)

			if req.AlternatePhoneNumber {
				glm.gls.AlternatePhoneNumber(page, username, password, req, accountscount, userAccountCount)
			} else {
				glm.gls.Manager(page, req, accountscount, userAccountCount, username, password)
			}

			accountscount++
			userAccountCount++
		}
		startIndex = endIndex

		time.Sleep(3 * time.Second)

		if startIndex%9 == 0 {
			fmt.Println("Switching")
			//closes  browser
			ctx.Close()

			ctx, err = glm.common.CreateContext(browser)
			if err != nil {
				log.Fatalf("Error creating context: %v", err)
				return err
			}
			fmt.Println("created context")
			defer ctx.Close()

			page, err = ctx.NewPage()
			if err != nil {
				log.Printf("Error creating page: %v", err)
				return err
			}

			if req.AddSmartlead {
		err := services.Login(page, string(req.SmartLeadEmail), req.SmartLeadPasswd)
		if err != nil {
			fmt.Println(err)
		}
	}
			accountscount = 1
			userAccountCount = 0
		}
		if startIndex <= numAccounts {
			fmt.Printf("Services completed for this account. Starting the next session...\n")
			time.Sleep(2 * time.Second)
		}
	}
	return nil
}

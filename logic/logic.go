package logic

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	gmailloginmaster "github.com/seew0/loginAutomation/auto/GmailLoginSmartLead"
	"github.com/seew0/loginAutomation/models"
	"github.com/seew0/loginAutomation/utils"
)

type Logic struct {
	glm gmailloginmaster.GmailLoginMaster
}

func ProvideLogic(glm gmailloginmaster.GmailLoginMaster) Logic {
	return Logic{glm}
}

func (l Logic) GetLoginSmartlead(req models.GmailLoginSmartlead, c *gin.Context, file *multipart.FileHeader) error {
	var filename string
	uploadPath := "./"

	if req.AccountType.IsValid() {
		switch req.AccountType {
		case models.AccountsTypeGmail:
			filename = "google_accounts.csv"
		case models.AccountsTypeOutlook:
			filename = "outlook_accounts.csv"
		}
	}

	filepath := filepath.Join(uploadPath, filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	if err := l.glm.GmailLoginManager(req); err != nil {
		return err
	}

	return nil
}

func (l Logic) UpdateFiveSimOptions(req models.FiveSimOptions) error {
	envFilePath := ".env"
	if req.Country == "" && req.Operator == "" {
		return fmt.Errorf("cant update request invalid use available countries or operator")
	}

	utils.UpdateEnvFile(envFilePath, "Country", fmt.Sprintf("'%s'",req.Country))

	utils.UpdateEnvFile(envFilePath, "Operator", fmt.Sprintf("'%s'",req.Operator))

	return nil
}

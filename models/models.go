package models

import (
	"mime/multipart"
	"net/mail"
)

type email string
type AccountsType string

const (
	AccountsTypeGmail   AccountsType = "gmail"
	AccountsTypeOutlook AccountsType = "outlook"
)

func (a AccountsType) IsValid() bool {
	switch a {
	case AccountsTypeGmail, AccountsTypeOutlook:
		return true
	}
	return false
}

type GmailLoginSmartlead struct {
	AccountType          AccountsType          `form:"accountType" binding:"required"`
	File                 *multipart.FileHeader `form:"file" binding:"required"`
	ImagePermission      bool                  `form:"imagePermissions"`
	AlternatePhoneNumber bool                  `form:"alternatePhoneNumber"`
	AddSmartlead         bool                  `form:"addSmartlead"`
	LessSecure           bool                  `form:"lessSecure"`
	SmartLeadEmail       email                 `form:"smartleadEmail"`
	SmartLeadPasswd      string                `form:"smartleadPassword"`
}

type countries string

type operators string


type FiveSimOptions struct{
	Country countries `json:"country"`
	Operator operators `json:"operator"`
}

func (e email) IsValid() bool {
	_, err := mail.ParseAddress(string(e))
	return err == nil
}

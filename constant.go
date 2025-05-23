package go_h2pay

import (
	"github.com/samber/lo"
	"strings"
)

// H2payDepositStatus represents the deposit status type
type H2payDepositStatus struct {
	Code string
	Name string
	Desc string
}

// String implements fmt.Stringer interface
func (s H2payDepositStatus) String() string {
	return s.Code
}

// Equal checks if the status code equals the given value
func (s H2payDepositStatus) Equal(value string) bool {
	return s.Code == value
}

// Predefined deposit status constants
var (
	H2payDepositStatusSuccess  = H2payDepositStatus{"000", "Success", "Success"}
	H2payDepositStatusFailed   = H2payDepositStatus{"001", "Failed", "Failed"}
	H2payDepositStatusApproved = H2payDepositStatus{"006", "Approved", "Approved"}
	H2payDepositStatusRejected = H2payDepositStatus{"007", "Rejected", "Rejected"}
	H2payDepositStatusCanceled = H2payDepositStatus{"008", "Canceled", "Canceled"}
	H2payDepositStatusPending  = H2payDepositStatus{"009", "Pending", "Pending"}
)

// StatusFromCode returns a status from its code
func StatusFromCode(code string) (H2payDepositStatus, bool) {
	switch strings.ToUpper(code) {
	case H2payDepositStatusSuccess.Code:
		return H2payDepositStatusSuccess, true
	case H2payDepositStatusFailed.Code:
		return H2payDepositStatusFailed, true
	case H2payDepositStatusApproved.Code:
		return H2payDepositStatusApproved, true
	case H2payDepositStatusRejected.Code:
		return H2payDepositStatusRejected, true
	case H2payDepositStatusCanceled.Code:
		return H2payDepositStatusCanceled, true
	case H2payDepositStatusPending.Code:
		return H2payDepositStatusPending, true
	default:
		return H2payDepositStatus{}, false
	}
}

//---------------------------------------

// H2payWithdrawStatus represents the withdrawal status type
type H2payWithdrawStatus struct {
	Code string
	Name string
	Desc string
}

// Value returns the status code (implements similar functionality to getValue())
func (s H2payWithdrawStatus) Value() string {
	return s.Code
}

// Equal checks if the status code equals the given value (implements eq())
func (s H2payWithdrawStatus) Equal(value string) bool {
	return s.Code == value
}

// String implements fmt.Stringer interface
func (s H2payWithdrawStatus) String() string {
	return s.Name
}

// Predefined withdrawal status constants
var (
	H2payWithdrawSuccess = H2payWithdrawStatus{"000", "Success", "Success"}
	H2payWithdrawFailed  = H2payWithdrawStatus{"001", "Failed", "Failed"}
)

// WithdrawStatusFromCode returns a status from its code
func WithdrawStatusFromCode(code string) (H2payWithdrawStatus, bool) {
	switch code {
	case H2payWithdrawSuccess.Code:
		return H2payWithdrawSuccess, true
	case H2payWithdrawFailed.Code:
		return H2payWithdrawFailed, true
	default:
		return H2payWithdrawStatus{}, false
	}
}

//==========================================

type H2payLanguage struct {
	Code string
	Desc string
}

var (
	H2payLanguageEnglish = H2payLanguage{"en-us", "English"}
	H2payLanguageChinese = H2payLanguage{"zh-cn", "Chinese Simplified"}
	H2payLanguageThai    = H2payLanguage{"th", "Thai"}
	H2payLanguageMalay   = H2payLanguage{"ms-my", "Malay (Malaysia)"}
	H2payLanguageVietnam = H2payLanguage{"vi-vn", "Vietnamese (Vietnam)"}
	H2payLanguageIndo    = H2payLanguage{"id-id", "Indonesian"}
	H2payLanguageBurmese = H2payLanguage{"bur", "Burmese"}
	H2payLanguageTagalog = H2payLanguage{"fil-ph", "Tagalog (Philippines)"}
	H2payLanguageHindi   = H2payLanguage{"hi-in", "Hindi (India)"}
	H2payLanguageKhmer   = H2payLanguage{"km-kh", "Khmer (Cambodia)"}
)

var LegalLanguageList = []H2payLanguage{
	H2payLanguageEnglish,
	H2payLanguageChinese,
	H2payLanguageThai,
	H2payLanguageMalay,
	H2payLanguageVietnam,
	H2payLanguageIndo,
	H2payLanguageBurmese,
	H2payLanguageTagalog,
	H2payLanguageHindi,
	H2payLanguageKhmer,
}

// 用来检查入参的正确性
func IsLanguageExist(code string) bool {
	_, ok := lo.Find(LegalLanguageList, func(i H2payLanguage) bool {
		return i.Code == code
	})
	return ok
}

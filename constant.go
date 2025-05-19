package go_h2pay

import "strings"

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

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
	Success  = H2payDepositStatus{"000", "Success", "Success"}
	Failed   = H2payDepositStatus{"001", "Failed", "Failed"}
	Approved = H2payDepositStatus{"006", "Approved", "Approved"}
	Rejected = H2payDepositStatus{"007", "Rejected", "Rejected"}
	Canceled = H2payDepositStatus{"008", "Canceled", "Canceled"}
	Pending  = H2payDepositStatus{"009", "Pending", "Pending"}
)

// StatusFromCode returns a status from its code
func StatusFromCode(code string) (H2payDepositStatus, bool) {
	switch strings.ToUpper(code) {
	case Success.Code:
		return Success, true
	case Failed.Code:
		return Failed, true
	case Approved.Code:
		return Approved, true
	case Rejected.Code:
		return Rejected, true
	case Canceled.Code:
		return Canceled, true
	case Pending.Code:
		return Pending, true
	default:
		return H2payDepositStatus{}, false
	}
}

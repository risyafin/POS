package transaction

import "errors"

var (
	errorAdminID  = errors.New("adminID not int ")
	errorBranchID = errors.New("branchID not int")
	errorUsername = errors.New("username not string")
)

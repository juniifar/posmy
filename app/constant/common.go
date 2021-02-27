package constant

// common
const (
	CreateBySystem = 0
	UpdateBySystem = 0

	Active   = 1
	Inactive = -1

	ErrQuerySetterNoDataFound = "<QuerySeter> no row found"
)

// otp
const (
	OTPTypeRegistration = "registration"
	OTPTypeLogin        = "login"
)

// session
const (
	SessionName = "Predict-Session"
)

// context
const (
	ContextUserID = "User-ID"
)

// employee type
const (
	EmployeeTypeGeneralManagerID = 1
	EmployeeTypeManagerID        = 2
	EmployeeTypeCashierID        = 3

	EmployeeTypeGeneralManagerLabel = "General Manager"
	EmployeeTypeManagerLabel        = "Manager"
	EmployeeTypeCashierLabel        = "Cashier"
)

package common

// ReasonCode represents reason code
type ReasonCode string

// General Error
const (
	ReasonGeneralError    ReasonCode = "-1"
	ReasonDBError         ReasonCode = "-2"
	ReasonCacheError      ReasonCode = "-3"
	ReasonAdapterError    ReasonCode = "-4"
	ReasonNotFound        ReasonCode = "-5"
	ReasonInvalidArgument ReasonCode = "-6"
)

// Application Error. From 1000
const (
	ReasonApplicationNotFound ReasonCode = "-1001"
	ReasonApplicationStatus   ReasonCode = "-1002"
)

// Transaction Error. From 2000
const (
	ReasonAmountNotAllow   ReasonCode = "-2001"
	ReasonTranTypeNotAllow ReasonCode = "-2011"
)

// User Error. From 3000
const (
	ReasonUserNotFound     ReasonCode = "-3001"
	ReasonUserNameNotValid ReasonCode = "-3011"
)

var reasonCodeValues = map[string]string{
	"-1": "General Error",
	"-2": "DB Error",
	"-3": "Cache Error",
	"-4": "Adapter Error",
	"-5": "Not Found",
	"-6": "Invalid Argument",
	// Application Error
	"-1001": "Application Not Found",
	"-1002": "Application Status Invalid",
	// Transaction Error
	"-2001": "Amount not allow",
	"-2011": "Transaction Type not allow",
	// User Error
	"-3001": "User not found",
	"-3011": "UserName not valid",
}

// Code represents reason code
func (rc ReasonCode) Code() string {
	return string(rc)
}

// Message represent reason message
func (rc ReasonCode) Message() string {
	if value, ok := reasonCodeValues[rc.Code()]; ok {
		return value
	}
	return ""
}

// ParseError parsing error code to message
func ParseError(err error) ReasonCode {
	if err == nil {
		return ReasonCode("")
	}
	return ReasonCode(err.Error())
}

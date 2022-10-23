package handlers

var bodyPatterns = map[string]string{
	"UserId":    "^(.+)$",
	"Amount":    "^(([1-9][0-9]*(\\.[0-9]{1,2})?)|(0\\.((0[1-9])|([1-9][0-9]?))))$",
	"ServiceId": "^(.+)$",
	"OrderId":   "^(.+)$",
}

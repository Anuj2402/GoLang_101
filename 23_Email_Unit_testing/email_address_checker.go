package main

import(
	"strings"
)

/*
	IsLinkedinEmployee func checks if the given email adddress ends with "linkedin.com"

*/


func IsLinkedinEmployee(adddress string ) bool {

	return strings.HasSuffix(adddress, "@linkedin.com")

}


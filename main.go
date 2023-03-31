package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {

	emailSplit := strings.Split(email, "@")

	emailInfo := strings.Split(emailSplit[1], ".")
	
	var domain string
	var tld string

	domain = emailInfo[0]

	if len(emailInfo) > 2 {
		tld = strings.Join(emailInfo[1:], ".")
	}
	hasil := fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
	return hasil

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("rangga.surya.p@students.esqbs.id"))
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}

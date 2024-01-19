package main

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/go-ldap/ldif"
)

func main() {
	var _ ldif.Entry
	_, _ = ldap.Dial("tcp", fmt.Sprintf("%s:%d", "ldap.example.com", 389))
}

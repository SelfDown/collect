package test

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"testing"
)

func Test_ldap(t *testing.T) {
	//替换example即可，可在Ldap中查询正确的链接写法
	//本文使用Windows Server 2012 Active Directory，打开管理工具 "Active Directory 域和信任关系"，左侧目录看到Ldap Url
	l, err := ldap.DialURL("ldap://172.26.0.20:389")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	_, err = l.SimpleBind(&ldap.SimpleBindRequest{
		Username: "cn=Manager,dc=weigao,dc=com",
		Password: "RQj9Ctjgc",
	})
	if err != nil {
		log.Fatalf("Failed to bind: %s\n", err)
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=weigao,dc=com", // The base dn to search
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=top)(cn=%s))", "chengyuan"), // The filter to apply
		[]string{"dn", "cn", "mail", "sn"},                      // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(sr.Entries))
	fmt.Println(len(sr.Entries[0].Attributes))
	fmt.Println(sr.Entries[0].Attributes[0].Name)
}

package public

import (
	"fmt"

	"github.com/eryajf/ldapool"
	ldap "github.com/go-ldap/ldap/v3"
)

// ========= test-ldap =============
const (
	LDAP_URL        = "ldap://localhost:389" // eryajf
	LDAP_BASE_DN    = "dc=eryajf,dc=net"
	LDAP_ADMIN_DN   = "cn=admin,dc=eryajf,dc=net"
	LDAP_ADMIN_PASS = "123456"
	LDAP_USER_DN    = "ou=People,dc=eryajf,dc=net"
	LDAP_GROUP_DN   = "ou=Group,dc=eryajf,dc=net"
)

// Init ldap conn
func InitCli() (l *ldap.Conn) {
	conn, err := ldapool.Open(ldapool.LdapConfig{
		Url:       LDAP_URL,
		BaseDN:    LDAP_BASE_DN,
		AdminDN:   LDAP_ADMIN_DN,
		AdminPass: LDAP_ADMIN_PASS,
		MaxOpen:   30,
		DsName:    "",
	})
	if err != nil {
		panic(fmt.Sprintf("get conn failed:%v\n", err))
	}

	return conn
}

// GetGroupDN
func GetGroupDN(group string) (dn string) {
	return fmt.Sprintf("cn=%s,%s", group, LDAP_GROUP_DN) // Example: cn=group,ou=Group,dc=eryajf,dc=net
}

// GetUserDN
func GetUserDN(user string) (dn string) {
	return fmt.Sprintf("uid=%s,%s", user, LDAP_USER_DN) // Example: uid=user,ou=People,dc=eryajf,dc=net
}

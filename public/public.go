package public

import (
	"fmt"
	"net"
	"time"

	ldap "github.com/go-ldap/ldap/v3"
)

// ========= test-ldap =============
const (
	LDAP_URL        = "ldap://localhost:390" // eryajf
	LDAP_BASE_DN    = "dc=eryajf,dc=net"
	LDAP_ADMIN_DN   = "cn=admin,dc=eryajf,dc=net"
	LDAP_ADMIN_PASS = "123465"
	LDAP_USER_DN    = "ou=People,dc=eryajf,dc=net"
	LDAP_GROUP_DN   = "ou=Group,dc=eryajf,dc=net"
)

// Init ldap conn
func InitCli() (l *ldap.Conn) {
	l, err := ldap.DialURL(LDAP_URL, ldap.DialWithDialer(&net.Dialer{Timeout: 5 * time.Second}))
	if err != nil {
		panic(err)
	}
	err = l.Bind(LDAP_ADMIN_DN, LDAP_ADMIN_PASS)
	if err != nil {
		panic(err)
	}
	return l
}

// GetGroupDN
func GetGroupDN(group string) (dn string) {
	return fmt.Sprintf("cn=%s,%s", group, LDAP_GROUP_DN) // Example: cn=group,ou=Group,dc=eryajf,dc=net
}

// GetUserDN
func GetUserDN(user string) (dn string) {
	return fmt.Sprintf("uid=%s,%s", user, LDAP_USER_DN) // Example: uid=user,ou=People,dc=eryajf,dc=net
}

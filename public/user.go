package public

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// User is a struct for user，The usage of some fields is defined according to the actual situation.
type User struct {
	CN               string `json:"cn"`               // 中文名全内容
	SN               string `json:"sn"`               // 名字拼音，可能是缩减的
	BusinessCategory string `json:"businessCategory"` // 业务类别，部门名字
	DepartmentNumber string `json:"departmentNumber"` // 部门编号，此处可以存放员工的职位
	Description      string `json:"description"`      // 描述
	DisplayName      string `json:"displayName"`      // 展示名字，可以是中文名字
	Mail             string `json:"mail"`             // 邮箱
	EmployeeNumber   string `json:"employeeNumber"`   // 员工工号
	GivenName        string `json:"givenName"`        // 给定名字，如果公司有花名，可以用这个字段
	PostalAddress    string `json:"postalAddress"`    // 家庭住址
	Mobile           string `json:"mobile"`           // 手机号
	UID              string `json:"uid"`              // 用户名
	Password         string `json:"password"`         // 密码
}

// AddUser add a user
func AddUser(user User) error {
	add := ldap.NewAddRequest(GetUserDN(user.UID), nil)

	add.Attribute("objectClass", []string{"inetOrgPerson"})
	add.Attribute("cn", []string{user.CN})
	add.Attribute("sn", []string{user.SN})
	add.Attribute("businessCategory", []string{user.BusinessCategory})
	add.Attribute("departmentNumber", []string{user.DepartmentNumber})
	add.Attribute("description", []string{user.Description})
	add.Attribute("displayName", []string{user.DisplayName})
	add.Attribute("mail", []string{user.Mail})
	add.Attribute("employeeNumber", []string{user.EmployeeNumber})
	add.Attribute("givenName", []string{user.GivenName})
	add.Attribute("postalAddress", []string{user.PostalAddress})
	add.Attribute("mobile", []string{user.Mobile})
	add.Attribute("uid", []string{user.UID})
	add.Attribute("userPassword", []string{user.Password})

	return InitCli().Add(add)
}

// UpdateUser update a user
func UpdateUser(user User) error {
	modify := ldap.NewModifyRequest(GetUserDN(user.UID), nil)

	modify.Replace("cn", []string{user.CN})
	modify.Replace("sn", []string{user.SN})
	modify.Replace("businessCategory", []string{user.BusinessCategory})
	modify.Replace("departmentNumber", []string{user.DepartmentNumber})
	modify.Replace("description", []string{user.Description})
	modify.Replace("displayName", []string{user.DisplayName})
	modify.Replace("mail", []string{user.Mail})
	modify.Replace("employeeNumber", []string{user.EmployeeNumber})
	modify.Replace("givenName", []string{user.GivenName})
	modify.Replace("postalAddress", []string{user.PostalAddress})
	modify.Replace("mobile", []string{user.Mobile})
	modify.Replace("uid", []string{user.UID})
	modify.Replace("userPassword", []string{user.Password})

	return InitCli().Modify(modify)
}

func GetAllUser() (users []*User, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		LDAP_USER_DN,                                                // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=organizationalPerson))", // This is Filter for LDAP query
		[]string{},                              // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)
	var sr *ldap.SearchResult
	// Search through ldap built-in search
	sr, err = InitCli().Search(searchRequest)
	if err != nil {
		return nil, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		for _, user := range sr.Entries {
			users = append(users, &User{
				CN:               user.GetAttributeValue("cn"),
				SN:               user.GetAttributeValue("sn"),
				BusinessCategory: user.GetAttributeValue("businessCategory"),
				DepartmentNumber: user.GetAttributeValue("departmentNumber"),
				Description:      user.GetAttributeValue("description"),
				DisplayName:      user.GetAttributeValue("displayName"),
				Mail:             user.GetAttributeValue("mail"),
				EmployeeNumber:   user.GetAttributeValue("employeeNumber"),
				GivenName:        user.GetAttributeValue("givenName"),
				PostalAddress:    user.GetAttributeValue("postalAddress"),
				Mobile:           user.GetAttributeValue("mobile"),
				UID:              user.GetAttributeValue("uid"),
				Password:         "******",
			})
		}
	}
	return
}

func GetUserByUID(uid string) (rst *User, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		LDAP_USER_DN,                                                // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", uid), // The request to filter user queries is defined here, and other filtering parameters can be changed.
		[]string{}, // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)
	var sr *ldap.SearchResult
	// Search through ldap built-in search
	sr, err = InitCli().Search(searchRequest)
	if err != nil {
		return nil, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		user := sr.Entries[0]
		rst = &User{
			CN:               user.GetAttributeValue("cn"),
			SN:               user.GetAttributeValue("sn"),
			BusinessCategory: user.GetAttributeValue("businessCategory"),
			DepartmentNumber: user.GetAttributeValue("departmentNumber"),
			Description:      user.GetAttributeValue("description"),
			DisplayName:      user.GetAttributeValue("displayName"),
			Mail:             user.GetAttributeValue("mail"),
			EmployeeNumber:   user.GetAttributeValue("employeeNumber"),
			GivenName:        user.GetAttributeValue("givenName"),
			PostalAddress:    user.GetAttributeValue("postalAddress"),
			Mobile:           user.GetAttributeValue("mobile"),
			UID:              user.GetAttributeValue("uid"),
			Password:         "******",
		}
	}
	return rst, nil
}

// DelUser delete a user
func DelUser(uid string) error {
	del := ldap.NewDelRequest(GetUserDN(uid), nil)
	return InitCli().Del(del)
}

// CheckUser Based on the user name and password, it is detected whether the user is normally available.
// Users often say that they have failed to log in to a certain system. This method can be used to check what is wrong.
func CheckUser(username, password string) error {
	udn := GetUserDN(username)
	if username == "admin" {
		udn = LDAP_ADMIN_DN
	}
	return InitCli().Bind(udn, password)
}

// see https://github.com/go-ldap/ldap/pull/54
// UpdateUserDN update user DN
func UpdateUserDN(olduid, newuid string) error {
	modify := ldap.NewModifyDNRequest(
		GetUserDN(olduid),
		fmt.Sprintf("uid=%s", newuid),
		true,
		"")
	return InitCli().ModifyDN(modify)
}

// ModifyUserPassword
// User uid, old password, and new password need to be passed.
// The old password is not a mandatory item. If it is left blank, the program will overwrite the old password and return it. Usually this should be sent to the modified students by mail.
func ModifyUserPassword(uid, oldpasswd, newpasswd string) (string, error) {
	passModify := ldap.NewPasswordModifyRequest(
		GetUserDN(uid),
		oldpasswd,
		newpasswd)
	result, err := InitCli().PasswordModify(passModify)
	if err != nil {
		return "", err
	}
	return result.GeneratedPassword, nil
}

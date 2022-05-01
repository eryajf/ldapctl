package public

import (
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

// Group struct, The usage of some fields is defined according to the actual situation.
type Group struct {
	CN       string   `json:"cn"`       // 是组的拼音名称
	Desc     string   `json:"desc"`     // 是组的描述,换句话说就是组的中文叫法
	Member   []string `json:"member"`   // 是组的成员
	ObjClass []string `json:"objClass"` // 是组的类型
}

func GetAllGroup() (groups []*Group, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		LDAP_GROUP_DN,                                               // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		"(&(objectClass=groupofnames))", // This is Filter for LDAP query
		[]string{},                      // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
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
		for _, v := range sr.Entries {
			groups = append(groups, &Group{
				CN:       v.GetAttributeValue("cn"),
				Desc:     v.GetAttributeValue("description"),
				Member:   v.GetAttributeValues("member"),
				ObjClass: v.GetAttributeValues("objectClass"),
			})
		}
	}
	return
}

func GetGroupMenber(group string) (menbers []string, err error) {
	// Construct query request
	searchRequest := ldap.NewSearchRequest(
		LDAP_GROUP_DN,                                               // This is basedn, we will start searching from this node.
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, // Here several parameters are respectively scope, derefAliases, sizeLimit, timeLimit,  typesOnly
		fmt.Sprintf("(&(objectClass=groupofnames)(cn=%s))", group), // This is Filter for LDAP query
		[]string{}, // Here are the attributes returned by the query, provided as an array. If empty, all attributes are returned
		nil,
	)
	// fmt.Sprintf("(&(objectClass=groupofnames)(cn=%s))", group)
	var sr *ldap.SearchResult
	// Search through ldap built-in search
	sr, err = InitCli().Search(searchRequest)
	if err != nil {
		return nil, err
	}
	// Refers to the entry that returns data. If it is greater than 0, the interface returns normally.
	if len(sr.Entries) > 0 {
		menbers = sr.Entries[0].GetAttributeValues("member")
	}
	return
}

// AddGroup
func AddGroup(group Group) error {
	add := ldap.NewAddRequest(GetGroupDN(group.CN), nil)

	add.Attribute("objectClass", []string{"groupOfNames", "top"}) // If groupOfNAmes is defined, member must be specified, otherwise the error is reported as follows：object class 'groupOfNames' requires attribute 'member'
	add.Attribute("cn", []string{group.CN})
	add.Attribute("description", []string{group.Desc})
	add.Attribute("member", []string{LDAP_ADMIN_DN}) // Therefore, when creating a group here, admin is added to it by default, so as not to report the above error without personnel during creation.

	return InitCli().Add(add)
}

// UpdateGroup
func UpdateGroup(group Group) error {
	modify := ldap.NewModifyRequest(GetGroupDN(group.CN), nil)
	modify.Replace("description", []string{group.Desc})
	return InitCli().Modify(modify)
}

// DelGroup
func DelGroup(group string) error {
	del := ldap.NewDelRequest(GetGroupDN(group), nil)
	return InitCli().Del(del)
}

// AddUserToGroup
func AddUserToGroup(user, group string) error {
	addUserToGroup := ldap.NewModifyRequest(GetGroupDN(group), nil)
	addUserToGroup.Add("member", []string{GetUserDN(user)})
	return InitCli().Modify(addUserToGroup)
}

// RemoveUserFromGroup
func RemoveUserFromGroup(user, group string) error {
	removeUserFromGroup := ldap.NewModifyRequest(GetGroupDN(group), nil)
	removeUserFromGroup.Delete("member", []string{GetUserDN(user)})
	return InitCli().Modify(removeUserFromGroup)
}

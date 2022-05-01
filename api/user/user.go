package user

import (
	"fmt"
	"ldapctl/public"

	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
)

// AddUserCmd
//   ./ldapctl user add --cn testuser1 --sn 测试用户1  --employeeNumber 001 --mail testuer1@eryajf.net --mobile 13888888881 --uid testuser1 --userPassword testuser1
var AddUserCmd = &cobra.Command{
	Use:   "add",
	Short: "add user",
	Long:  `Add user actions`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		sn, _ := cmd.Flags().GetString("sn")
		businessCategory, _ := cmd.Flags().GetString("businessCategory")
		departmentNumber, _ := cmd.Flags().GetString("departmentNumber")
		description, _ := cmd.Flags().GetString("desc")
		displayName, _ := cmd.Flags().GetString("displayName")
		mail, _ := cmd.Flags().GetString("mail")
		employeeNumber, _ := cmd.Flags().GetString("employeeNumber")
		givenName, _ := cmd.Flags().GetString("givenName")
		postalAddress, _ := cmd.Flags().GetString("postalAddress")
		mobile, _ := cmd.Flags().GetString("mobile")
		userPassword, _ := cmd.Flags().GetString("userPassword")

		uid, _ := cmd.Flags().GetString("uid")
		if businessCategory == "" {
			businessCategory = "Undefined"
		}
		if departmentNumber == "" {
			departmentNumber = "Undefined"
		}
		if description == "" {
			description = "Undefined"
		}
		if displayName == "" {
			displayName = "Undefined"
		}
		if givenName == "" {
			givenName = "Undefined"
		}
		if postalAddress == "" {
			postalAddress = "Undefined"
		}

		err := public.AddUser(public.User{
			CN:               cn,
			SN:               sn,
			BusinessCategory: businessCategory,
			DepartmentNumber: departmentNumber,
			Description:      description,
			DisplayName:      displayName,
			Mail:             mail,
			EmployeeNumber:   employeeNumber,
			GivenName:        givenName,
			PostalAddress:    postalAddress,
			Mobile:           mobile,
			UID:              uid,
			Password:         userPassword,
		})
		if err != nil {
			fmt.Printf("add user failed, err: %v\n", err)
		} else {
			fmt.Println("add user success")
		}
	},
}

// UpdateUserCmd
// 	./ldapctl user update -u testuser1 --displayName testtest
var UpdateUserCmd = &cobra.Command{
	Use:   "update",
	Short: "update user",
	Long:  `Update user actions`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		sn, _ := cmd.Flags().GetString("sn")
		businessCategory, _ := cmd.Flags().GetString("businessCategory")
		departmentNumber, _ := cmd.Flags().GetString("departmentNumber")
		description, _ := cmd.Flags().GetString("description")
		displayName, _ := cmd.Flags().GetString("displayName")
		mail, _ := cmd.Flags().GetString("mail")
		employeeNumber, _ := cmd.Flags().GetString("employeeNumber")
		givenName, _ := cmd.Flags().GetString("givenName")
		postalAddress, _ := cmd.Flags().GetString("postalAddress")
		mobile, _ := cmd.Flags().GetString("mobild")
		uid, _ := cmd.Flags().GetString("uid")

		user, err := public.GetUserByUID(uid)
		if err != nil {
			fmt.Printf("get user failed, err: %v\n", err)
		}
		if cn == "" {
			cn = user.CN
		}
		if sn == "" {
			sn = user.SN
		}
		if businessCategory == "" {
			businessCategory = user.BusinessCategory
		}
		if departmentNumber == "" {
			departmentNumber = user.DepartmentNumber
		}
		if description == "" {
			description = user.Description
		}
		if displayName == "" {
			displayName = user.DisplayName
		}
		if mail == "" {
			mail = user.Mail
		}
		if employeeNumber == "" {
			employeeNumber = user.EmployeeNumber
		}
		if givenName == "" {
			givenName = user.GivenName
		}
		if postalAddress == "" {
			postalAddress = user.PostalAddress
		}
		if mobile == "" {
			mobile = user.Mobile
		}

		err = public.UpdateUser(public.User{
			CN:               cn,
			SN:               sn,
			BusinessCategory: businessCategory,
			DepartmentNumber: departmentNumber,
			Description:      description,
			DisplayName:      displayName,
			Mail:             mail,
			EmployeeNumber:   employeeNumber,
			GivenName:        givenName,
			PostalAddress:    postalAddress,
			Mobile:           mobile,
			UID:              uid,
		})
		if err != nil {
			fmt.Printf("add user failed, err: %v\n", err)
		} else {
			fmt.Println("update user success")
		}
	},
}

// GetAllUsersCmd
// 		./ldapctl user getall
var GetAllUsersCmd = &cobra.Command{
	Use:   "getall",
	Short: "list all users",
	Long:  `list all users, No parameters are required`,
	Run: func(cmd *cobra.Command, args []string) {
		users, err := public.GetAllUser()
		if err != nil {
			fmt.Printf("get all user failed: %v\n", err)
			return
		}
		tb, err := gotable.Create("cn", "sn", "businessCategory", "departmentNumber", "description", "displayName", "mail", "employeeNumber", "givenName", "postalAddress", "mobile", "uid", "userPassword")
		if err != nil {
			fmt.Printf("create table failed: %v\n", err)
			return
		}
		for _, user := range users {
			row := make(map[string]string)
			row["cn"] = user.CN
			row["sn"] = user.SN
			row["businessCategory"] = user.BusinessCategory
			row["departmentNumber"] = user.DepartmentNumber
			row["description"] = user.Description
			row["displayName"] = user.DisplayName
			row["mail"] = user.Mail
			row["employeeNumber"] = user.EmployeeNumber
			row["givenName"] = user.GivenName
			row["postalAddress"] = user.PostalAddress
			row["mobile"] = user.Mobile
			row["uid"] = user.UID
			row["userPassword"] = user.Password
			tb.AddRow(row)
		}
		fmt.Println(tb)
	},
}

// GetUserByUIDCmd
//  ./ldapctl user get -u testuser1
var GetUserByUIDCmd = &cobra.Command{
	Use:   "get",
	Short: "get user by uid",
	Long:  `get user by uid`,
	Run: func(cmd *cobra.Command, args []string) {
		uid, _ := cmd.Flags().GetString("uid")
		user, err := public.GetUserByUID(uid)
		if err != nil {
			fmt.Printf("get all user failed: %v\n", err)
			return
		}
		tb, err := gotable.Create("cn", "sn", "businessCategory", "departmentNumber", "description", "displayName", "mail", "employeeNumber", "givenName", "postalAddress", "mobile", "uid", "userPassword")
		if err != nil {
			fmt.Printf("create table failed: %v\n", err)
			return
		}
		row := make(map[string]string)
		row["cn"] = user.CN
		row["sn"] = user.SN
		row["businessCategory"] = user.BusinessCategory
		row["departmentNumber"] = user.DepartmentNumber
		row["description"] = user.Description
		row["displayName"] = user.DisplayName
		row["mail"] = user.Mail
		row["employeeNumber"] = user.EmployeeNumber
		row["givenName"] = user.GivenName
		row["postalAddress"] = user.PostalAddress
		row["mobile"] = user.Mobile
		row["uid"] = user.UID
		row["userPassword"] = user.Password
		tb.AddRow(row)
		fmt.Println(tb)
	},
}

// DelUserByUIDCmd
// 	./ldapctl user del -u testuser1
var DelUserByUIDCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete user by uid",
	Long:  `delete user by uid`,
	Run: func(cmd *cobra.Command, args []string) {
		uid, _ := cmd.Flags().GetString("uid")
		err := public.DelUser(uid)
		if err != nil {
			fmt.Printf("delete user %s failed: %v\n", uid, err)
		} else {
			fmt.Println("delete user success")
		}
	},
}

// CheckUserPassCmd
// 	./ldapctl user checkuser -u testuser1 -p testuser1
var CheckUserPassCmd = &cobra.Command{
	Use:   "checkuser",
	Short: "Check whether the user password is correct",
	Long:  `Check whether the user password is correct`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("uid")
		pwd, _ := cmd.Flags().GetString("pwd")
		err := public.CheckUser(username, pwd)
		if err != nil {
			fmt.Printf("check user %s password failed: %v\n", username, err)
		} else {
			fmt.Println("check user password success")
		}
	},
}

// UpdataUserDNCmd
// 	./ldapctl user updateuserdn -o testuser1 -n user1
var UpdataUserDNCmd = &cobra.Command{
	Use:   "updateuserdn",
	Short: "update user uid",
	Long:  `update user uid`,
	Run: func(cmd *cobra.Command, args []string) {
		olduid, _ := cmd.Flags().GetString("olduid")
		newuid, _ := cmd.Flags().GetString("newuid")
		err := public.UpdateUserDN(olduid, newuid)
		if err != nil {
			fmt.Printf("check user %s password failed: %v\n", olduid, err)
		}
	},
}

// ChangeUserPwdCmd
// 	./ldapctl user changeuserpwd -u testuser1 -o testuser1 -n user1
var ChangeUserPwdCmd = &cobra.Command{
	Use:   "changeuserpwd",
	Short: "update user uid",
	Long:  `update user uid`,
	Run: func(cmd *cobra.Command, args []string) {
		uid, _ := cmd.Flags().GetString("uid")
		oldpwd, _ := cmd.Flags().GetString("oldpwd")
		newpwd, _ := cmd.Flags().GetString("newpwd")
		pwd, err := public.ModifyUserPassword(uid, oldpwd, newpwd)
		if err != nil {
			fmt.Printf("check user %s password failed: %v\n", uid, err)
		} else {
			fmt.Println("change user password success")
			fmt.Printf("new password is %s\n", pwd)
		}
	},
}

/*
Copyright © 2022 eryajf Linuxlql@163.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/eryajf/ldapctl/api/user"

	"github.com/spf13/cobra"
)

// userCmd represents the jenkins command
var userCmd = &cobra.Command{
	Use:  "user",
	Long: `User-related operations`,
}

func init() {
	rootCmd.AddCommand(userCmd)

	// get all users
	userCmd.AddCommand(user.GetAllUsersCmd)

	// add user
	userCmd.AddCommand(user.AddUserCmd)
	aset := user.AddUserCmd.Flags()
	aset.StringP("cn", "c", "", "user cn, It is usually a user name, adjusted according to the actual situation of your own use.")
	aset.StringP("sn", "s", "", "user sn, Usually a surname, adjusted according to the actual situation of your own use")
	aset.StringP("businessCategory", "b", "", "user businessCategory")
	aset.StringP("departmentNumber", "d", "", "user departmentNumber")
	aset.StringP("desc", "", "", "user description")
	aset.StringP("displayName", "", "", "user displayName")
	aset.StringP("mail", "m", "", "user mail")
	aset.StringP("employeeNumber", "e", "", "user employeeNumber")
	aset.StringP("givenName", "g", "", "user givenName")
	aset.StringP("postalAddress", "a", "", "user postalAddress")
	aset.StringP("mobile", "", "", "user mobile")
	aset.StringP("uid", "u", "", "user uid,Usually consistent with sn, can be adjusted according to the actual situation")
	aset.StringP("userPassword", "p", "", "userPassword")
	user.AddUserCmd.MarkFlagRequired("cn")
	user.AddUserCmd.MarkFlagRequired("sn")
	user.AddUserCmd.MarkFlagRequired("mail")
	user.AddUserCmd.MarkFlagRequired("employeeNumber")
	user.AddUserCmd.MarkFlagRequired("mobile")
	user.AddUserCmd.MarkFlagRequired("uid")
	user.AddUserCmd.MarkFlagRequired("userPassword")

	// update user, Only information other than the user's uid can be updated. If you want to update the user's uid, you need to use the UpdateUserDN method.
	userCmd.AddCommand(user.UpdateUserCmd)
	bset := user.UpdateUserCmd.Flags()
	bset.StringP("uid", "u", "", "user uid, Uid is the unique ID of the user when updating user information. uid cannot be updated directly here.")
	bset.StringP("cn", "c", "", "user cn, It is usually a user name, adjusted according to the actual situation of your own use.")
	bset.StringP("sn", "s", "", "user sn, Usually a surname, adjusted according to the actual situation of your own use")
	bset.StringP("businessCategory", "b", "", "user businessCategory")
	bset.StringP("departmentNumber", "d", "", "user departmentNumber")
	bset.StringP("description", "", "", "user description")
	bset.StringP("displayName", "", "", "user displayName")
	bset.StringP("mail", "m", "", "user mail")
	bset.StringP("employeeNumber", "e", "", "user employeeNumber")
	bset.StringP("givenName", "g", "", "user givenName")
	bset.StringP("postalAddress", "a", "", "user postalAddress")
	bset.StringP("mobile", "", "", "user mobile")
	user.UpdateUserCmd.MarkFlagRequired("uid")

	// get user by uid
	userCmd.AddCommand(user.GetUserByUIDCmd)
	cset := user.GetUserByUIDCmd.Flags()
	cset.StringP("uid", "u", "", "user uid")
	user.GetUserByUIDCmd.MarkFlagRequired("uid")

	// delete user by uid
	userCmd.AddCommand(user.DelUserByUIDCmd)
	dset := user.DelUserByUIDCmd.Flags()
	dset.StringP("uid", "u", "", "user uid")
	user.DelUserByUIDCmd.MarkFlagRequired("uid")

	// Check whether the user password is correct
	userCmd.AddCommand(user.CheckUserPassCmd)
	eset := user.CheckUserPassCmd.Flags()
	eset.StringP("uid", "u", "", "user uid")
	eset.StringP("pwd", "p", "", "user password")
	user.CheckUserPassCmd.MarkFlagRequired("uid")
	user.CheckUserPassCmd.MarkFlagRequired("pwd")

	// update user uid
	userCmd.AddCommand(user.UpdataUserDNCmd)
	fset := user.UpdataUserDNCmd.Flags()
	fset.StringP("olduid", "o", "", "user old uid")
	fset.StringP("newuid", "n", "", "user new uid")
	user.UpdataUserDNCmd.MarkFlagRequired("olduid")
	user.UpdataUserDNCmd.MarkFlagRequired("newuid")

	// change user password
	userCmd.AddCommand(user.ChangeUserPwdCmd)
	gset := user.ChangeUserPwdCmd.Flags()
	gset.StringP("uid", "u", "", "user uid")
	gset.StringP("oldpwd", "o", "", "user old password")
	gset.StringP("newpwd", "n", "", "user new password")
	user.ChangeUserPwdCmd.MarkFlagRequired("uid")

}

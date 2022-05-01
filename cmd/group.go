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
	"ldapctl/api/group"

	"github.com/spf13/cobra"
)

// groupCmd represents the jenkins command
var groupCmd = &cobra.Command{
	Use:  "group",
	Long: `Group-related operations`,
}

func init() {
	rootCmd.AddCommand(groupCmd)

	// get all groups
	groupCmd.AddCommand(group.GetAllGroupsCmd)

	// get all group menbers
	groupCmd.AddCommand(group.GetGroupMenbersCmd)
	set := group.GetGroupMenbersCmd.Flags()
	set.StringP("cn", "c", "", "group cn")
	group.GetGroupMenbersCmd.MarkFlagRequired("group")

	// add group
	groupCmd.AddCommand(group.AddGroupCmd)
	aset := group.AddGroupCmd.Flags()
	aset.StringP("cn", "c", "", "group cn, It is usually a group name")
	aset.StringP("desc", "d", "", "group description")
	group.AddGroupCmd.MarkFlagRequired("cn")
	group.AddGroupCmd.MarkFlagRequired("desc")

	// update group
	groupCmd.AddCommand(group.UpdateGroupCmd)
	bset := group.UpdateGroupCmd.Flags()
	bset.StringP("cn", "c", "", "group cn, When updating the group, cn filters the group as the unique ID of the group.")
	bset.StringP("desc", "d", "", "group description")
	group.UpdateGroupCmd.MarkFlagRequired("cn")
	group.UpdateGroupCmd.MarkFlagRequired("desc")

	// delete group
	groupCmd.AddCommand(group.DeleteGroupCmd)
	cset := group.DeleteGroupCmd.Flags()
	cset.StringP("cn", "c", "", "group cn, When deleting the group, cn filters the group as the unique ID of the group.")
	group.UpdateGroupCmd.MarkFlagRequired("cn")

	// add user to group
	groupCmd.AddCommand(group.AddUserToGroupCmd)
	dset := group.AddUserToGroupCmd.Flags()
	dset.StringP("cn", "c", "", "group cn")
	dset.StringP("user", "u", "", "user uid")
	group.AddUserToGroupCmd.MarkFlagRequired("cn")
	group.AddUserToGroupCmd.MarkFlagRequired("user")

	// remove user to group
	groupCmd.AddCommand(group.RemoveUserFromGroupCmd)
	eset := group.RemoveUserFromGroupCmd.Flags()
	eset.StringP("cn", "c", "", "group cn")
	eset.StringP("user", "u", "", "user uid")
	group.RemoveUserFromGroupCmd.MarkFlagRequired("cn")
	group.RemoveUserFromGroupCmd.MarkFlagRequired("user")
}

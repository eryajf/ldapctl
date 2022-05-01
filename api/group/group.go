package group

import (
	"fmt"
	"ldapctl/public"

	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
)

// GetAllGroupsCmd
// 		./ldapctl group getall
var GetAllGroupsCmd = &cobra.Command{
	Use:   "getall",
	Short: "list all groups",
	Long:  `list all groups, No parameters are required`,
	Run: func(cmd *cobra.Command, args []string) {
		groups, err := public.GetAllGroup()
		if err != nil {
			fmt.Printf("get all group failed: %v\n", err)
			return
		}

		tb, err := gotable.Create("cn", "description")
		if err != nil {
			fmt.Printf("create table failed: %v\n", err)
			return
		}
		for _, group := range groups {
			row := make(map[string]string)
			row["cn"] = group.CN
			row["description"] = group.Desc
			tb.AddRow(row)
		}
		fmt.Println(tb)
	},
}

// GetGroupMenbersCmd
// 		./ldapctl group get -c groupname
var GetGroupMenbersCmd = &cobra.Command{
	Use:   "get",
	Short: "get group menbers",
	Long:  `get group menbers`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		members, err := public.GetGroupMenber(cn)
		if err != nil {
			fmt.Printf("get all group failed: %v\n", err)
			return
		}
		for _, member := range members {
			fmt.Println(member)
		}
	},
}

// AddGroupCmd
// 	./ldapctl group add -c testg -d "测试分组"
var AddGroupCmd = &cobra.Command{
	Use:   "add",
	Short: "add group",
	Long:  `add group`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		desc, _ := cmd.Flags().GetString("desc")
		err := public.AddGroup(public.Group{
			CN:   cn,
			Desc: desc,
		})
		if err != nil {
			fmt.Printf("add group failed: %v\n", err)
			return
		} else {
			fmt.Println("add group success")
		}
	},
}

// UpdateGroupCmd
// 	./ldapctl group update -c testg -d "测试分组1"
var UpdateGroupCmd = &cobra.Command{
	Use:   "update",
	Short: "update group",
	Long:  `update group`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		desc, _ := cmd.Flags().GetString("desc")
		err := public.UpdateGroup(public.Group{
			CN:   cn,
			Desc: desc,
		})
		if err != nil {
			fmt.Printf("update group failed: %v\n", err)
			return
		} else {
			fmt.Println("update group success")
		}
	},
}

// DeleteGroupCmd
// 	./ldapctl group update -c testg -d "测试分组1"
var DeleteGroupCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete group",
	Long:  `delete group`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		err := public.DelGroup(cn)
		if err != nil {
			fmt.Printf("delete group failed: %v\n", err)
			return
		} else {
			fmt.Println("delete group success")
		}
	},
}

// AddUserToGroupCmd
// 		./ldapctl group getall
var AddUserToGroupCmd = &cobra.Command{
	Use:   "adduser",
	Short: "add user to group",
	Long:  `add user to group`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		user, _ := cmd.Flags().GetString("user")
		err := public.AddUserToGroup(user, cn)
		if err != nil {
			fmt.Printf("add user to group: %v\n", err)
			return
		} else {
			fmt.Println("add user to group success")
		}
	},
}

// RemoveUserFromGroupCmd
// 		./ldapctl group getall
var RemoveUserFromGroupCmd = &cobra.Command{
	Use:   "removeuser",
	Short: "remove user to group",
	Long:  `remove user to group`,
	Run: func(cmd *cobra.Command, args []string) {
		cn, _ := cmd.Flags().GetString("cn")
		user, _ := cmd.Flags().GetString("user")
		err := public.RemoveUserFromGroup(user, cn)
		if err != nil {
			fmt.Printf("remove user to group: %v\n", err)
			return
		} else {
			fmt.Println("remove user to group success")
		}
	},
}

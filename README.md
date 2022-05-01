
**ldapctl** is a openLDAP CLI based on [ldap](https://github.com/go-ldap/ldap) library. 🚀

It only takes two steps to start the ldapctl experience：

1. use docs/start-ldap-eryajf.sh script Start a openLDAP instance locally through docker。
2. Run the make build compilation project, and then you can test it.。

If you want to manage your ldap, you can directly modify the configuration information in public/public.go, then compile the project and put it into use.

At present, it provides simple management of users and groups, which is mainly used to learn ldap library. There may be some deficiencies. Welcome to communicate.

- user
  -  add: add user
  -  changeuserpwd: update user uid
  -  checkuser: Check whether the user password is correct
  -  delete: delete user by uid
  -  get: get user by uid
  -  getall: list all users
  -  update: update user
  -  updateuserdn: update user uid

- group
  - add: add group
  - adduser: add user to group
  - delete: delete group
  - get: get group menbers
  - getall: list all groups
  - removeuser: remove user to group
  - update: update group


If you have other functions you want to add, please submit them in issue.
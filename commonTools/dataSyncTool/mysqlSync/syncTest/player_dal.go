//package test
package main

import (
	"fmt"
)

var (
	con_player_tableUserName = "sys_user"
)

func init() {
	registerSyncObj(con_player_tableUserName)
}

func insert(obj *player) {
	command := fmt.Sprintf("INSERT INTO `%s` (`UserID`,`UserName`) VALUES ('%v','%v') ", con_player_tableUserName, obj.UserID, obj.UserName)
	save(con_player_tableUserName, command)
}

func update(obj *player) {
	command := fmt.Sprintf("UPDATE `%s` SET  `UserName` = '%v' WHERE `UserID` = '%v';", con_player_tableUserName, obj.UserName, obj.UserID)
	save(con_player_tableUserName, command)
}

func clear(obj *player) {
	command := fmt.Sprintf("DELETE FROM %s where UserID = '%v';", con_player_tableUserName, obj.UserID)
	save(con_player_tableUserName, command)
}

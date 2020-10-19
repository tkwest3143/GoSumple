package db

import (
	"github/GoSumple/gin_sumple/go/data"
)

//UserRelationInsert userRelations情報を挿入します
func UserRelationInsert(insData data.UserRelation) {
	d := GormConnect()
	d.Create(&insData)
	defer d.Close()
}

//GetFriendList 引数に指定されたユーザがをもとにフレンドリストを取得します。
func GetFriendList(userID string) []string {
	d := GormConnect()
	selData := []data.UserRelation{}
	d.Find(&selData, "user_id1=? or user_id2 = ?", userID, userID)
	defer d.Close()
	friendList := []string{}
	for _, b := range selData {
		if userID != b.UserID1 {
			friendList = append(friendList, b.UserID1)
		} else {
			friendList = append(friendList, b.UserID2)
		}
	}
	return friendList
}

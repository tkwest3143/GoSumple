package db

import (
	"github/GoSumple/chat-application/src/go/data"

	"github.com/jinzhu/gorm"
)

//UserRelationDBInit UserRelationsテーブルの初期化を行います
func UserRelationDBInit(d *gorm.DB) {
	if !d.HasTable(&data.UserRelation{}) {
		d.CreateTable(&data.UserRelation{})
	}
}

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
	d.Find(&selData, "user_id1=?", userID)
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

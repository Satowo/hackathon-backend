package dao

import (
	"errors"
	"hackathon-backend/model"
	"log"
)

func ChannelSearchDao() ([]model.Channel, error) {
	rows, err := db.Query("SELECT * FROM channel")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	channels := make([]model.Channel, 0)
	for rows.Next() {
		var u model.Channel
		if err := rows.Scan(&u.ChannelId, &u.ChannelName); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		channels = append(channels, u)
	}

	return channels, nil
}

func ChannelJoinDao(Id string, channelId string, userId string) (model.UserInfo, error) {
	var userInfo model.UserInfo

	//重複レコードがないか確認
	duplicateCheckStmt, err := db.Prepare("SELECT COUNT(*) FROM channelMember WHERE userId = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (duplicateCheckStmt), %v\n", err)
		return userInfo, err
	}
	defer duplicateCheckStmt.Close()

	var count int
	err = duplicateCheckStmt.QueryRow(userId).Scan(&count)
	if err != nil {
		log.Printf("fail: duplicateCheckStmt.QueryRow, %v\n", err)
		return userInfo, err
	}

	// 重複レコードが存在する場合にエラーを返す
	if count > 0 {
		return userInfo, errors.New("duplicate record found")
	}

	// データがテーブルの構造に一致しているか確認
	stmt, err := db.Prepare("INSERT INTO channelMember (Id, channelId, userId) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return userInfo, err
	}
	defer stmt.Close()

	// データをデータベースに挿入
	_, err = stmt.Exec(Id, channelId, userId)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return userInfo, err
	}

	var email string
	err = db.QueryRow(`SELECT email FROM appUser WHERE userId = ?`, userId).Scan(&email)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return userInfo, err
	}

	userInfo, err = UserInfoDao(email)
	if err != nil {
		log.Printf("fail: UserInfoDao(channelJoin), %v\n", err)
		return userInfo, err
	}

	return userInfo, nil
}

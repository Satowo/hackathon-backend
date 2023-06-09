package dao

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"hackathon-backend/model"
	"log"
)

func ChannelMemberSearchDao(channelId string) ([]model.AppUser, error) {
	//取得したクエリパラメータとchannel_idが一致するカラムのuser_idを取得
	rows, err := db.Query("SELECT appUser.userId, userName, email, password FROM appUser INNER JOIN channelMember ON appUser.userId = channelMember.userId WHERE channelId = ?", channelId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	members := make([]model.AppUser, 0)
	for rows.Next() {
		var u model.AppUser
		if err := rows.Scan(&u.UserId, &u.UserName, &u.Email, &u.Password); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		members = append(members, u)
	}
	return members, nil
}

func ChannelJoinDao(Id string, channelId string, userId string) (model.UserInfo, error) {
	var userInfo model.UserInfo

	//重複レコードがないか確認
	duplicateCheckStmt, err := db.Prepare("SELECT COUNT(*) FROM channelMember WHERE channelId = ? AND userId = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (duplicateCheckStmt), %v\n", err)
		return userInfo, err
	}
	defer duplicateCheckStmt.Close()

	var count int
	err = duplicateCheckStmt.QueryRow(channelId, userId).Scan(&count)
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

func ChannelLeaveDao(Id string, channelId string, userId string) (model.UserInfo, error) {
	var userInfo model.UserInfo

	//重複レコードがあるか確認
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

	// 重複レコードがない場合にエラーを返す
	if count == 0 {
		return userInfo, errors.New("duplicate record found")
	}

	// データがテーブルの構造に一致しているか確認
	stmt, err := db.Prepare("DELETE FROM channelMember WHERE userId = ? AND channelId = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return userInfo, err
	}
	defer stmt.Close()

	// データをデータベースに反映
	_, err = stmt.Exec(userId, channelId)
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

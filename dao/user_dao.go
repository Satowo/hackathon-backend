package dao

import (
	"errors"
	"hackathon-backend/model"
	"log"
)

func UserInfoDao(email string) (model.UserInfo, error) {
	var userInfo model.UserInfo
	err := db.QueryRow(`SELECT userId, userName FROM appUser WHERE email = ?`, email).Scan(&userInfo.UserId, &userInfo.UserName)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return userInfo, err
	}

	// 最初のクエリを実行して複数の channelId を取得
	rows, err := db.Query("SELECT channelId FROM channelMember INNER JOIN appUser ON channelMember.userId = appUser.userId WHERE appUser.email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var channels []model.Channel
	for rows.Next() {
		var u model.Channel
		err := rows.Scan(&u.ChannelId)
		if err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)
			return userInfo, err
		}

		// 2番目のクエリを実行して channelId に対応する channelName を取得
		err = db.QueryRow("SELECT channelName FROM channel WHERE channelId = ?", u.ChannelId).Scan(&u.ChannelName)
		if err != nil {
			log.Printf("fail: db.QueryRow, %v\n", err)
			return userInfo, err
		}

		channels = append(channels, u)
	}
	userInfo.Channels = channels
	userInfo.Email = email

	return userInfo, nil
}

func UserRegisterDao(userId string, userName string, email string, password string) error {
	duplicateCheckStmt, err := db.Prepare("SELECT COUNT(*) FROM appUser WHERE userId = ? OR userName = ? OR email = ? OR password = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (duplicateCheckStmt), %v\n", err)
		return err
	}
	defer duplicateCheckStmt.Close()

	var count int
	err = duplicateCheckStmt.QueryRow(userId, userName, email, password).Scan(&count)
	if err != nil {
		log.Printf("fail: duplicateCheckStmt.QueryRow, %v\n", err)
		return err
	}

	// 重複レコードが存在する場合にエラーを返す
	if count > 0 {
		return errors.New("duplicate record found")
	}

	// データがテーブルの構造に一致しているか確認
	stmt, err := db.Prepare("INSERT INTO appUser (userId, userName, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return err
	}
	defer stmt.Close()

	// データをデータベースに挿入
	_, err = stmt.Exec(userId, userName, email, password)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return err
	}

	return nil
}

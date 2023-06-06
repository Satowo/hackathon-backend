package dao

import (
	"errors"
	"hackathon-backend/model"
	"log"
)

func UserSearchDao() ([]model.AppUser, error) {
	rows, err := db.Query("SELECT * FROM appUser")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	users := make([]model.AppUser, 0)
	for rows.Next() {
		var u model.AppUser
		if err := rows.Scan(&u.UserId, &u.UserName, &u.Email, &u.Password); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
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

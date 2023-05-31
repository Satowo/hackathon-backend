package dao

import (
	"hackathon-backend/model"
	"log"
)

func UserSearchDao() ([]model.User, error) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	users := make([]model.User, 0)
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
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

func DataBaseClose() error {
	//DBを閉じる
	err := db.Close()
	return err
}

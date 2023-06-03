package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"hackathon-backend/model"
	"log"
)

func ChannelMemberSearchDao(channelId string) ([]model.AppUser, error) {
	//取得したクエリパラメータとchannel_idが一致するカラムのuser_idを取得
	rows, err := db.Query("SELECT app_user.user_id, user_name, email, password FROM app_user INNER JOIN channel_member ON app_user.user_id = channel_member.user_id WHERE channel_id = ?", channelId)
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

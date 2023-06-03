package dao

import (
	"hackathon-backend/model"
	"log"
)

func MessageSearchDao() ([]model.Message, error) {
	rows, err := db.Query("SELECT * FROM message")
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	messages := make([]model.Message, 0)
	for rows.Next() {
		var u model.Message
		if err := rows.Scan(&u.MessageId, &u.UserId, &u.ChannelId, &u.MessageContent); err != nil {
			log.Printf("fail: rows.Scan, %v\n", err)

			if err := rows.Close(); err != nil { // 500を返して終了するが、その前にrowsのClose処理が必要
				log.Printf("fail: rows.Close(), %v\n", err)
			}
			return nil, err
		}
		messages = append(messages, u)
	}

	return messages, nil
}

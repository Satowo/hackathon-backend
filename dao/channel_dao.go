package dao

import (
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

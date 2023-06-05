package dao

import (
	"hackathon-backend/model"
	"log"
)

func MessageSearchDao(channelId string) ([]model.Message, error) {
	rows, err := db.Query("SELECT messageId, appUser.userId, userName, channelId, messageContent, edited FROM message INNER JOIN appUser ON message.userId = appUser.userId WHERE channelId = ?", channelId)
	if err != nil {
		log.Printf("fail: db.Query, %v\n", err)
		return nil, err
	}

	messages := make([]model.Message, 0)
	for rows.Next() {
		var u model.Message
		if err := rows.Scan(&u.MessageId, &u.UserId, &u.UserName, &u.ChannelId, &u.MessageContent, &u.Edited); err != nil {
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

func MessageRegisterDao(messageId string, userId string, channelId string, messageContent string) ([]model.Message, error) {
	// データがテーブルの構造に一致しているか確認
	stmt, err := db.Prepare("INSERT INTO message (messageId, userId, channelId, messageContent) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	// データをデータベースに挿入
	_, err = stmt.Exec(messageId, userId, channelId, messageContent)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return nil, err
	}

	messages, err := MessageSearchDao(channelId)
	if err != nil {
		log.Printf("fail: MessageSearchDao, %v\n", err)
		return nil, err
	}

	return messages, nil
}

func MessageEditDao(messageId string, channelId string, messageContent string) ([]model.Message, error) {
	// データを編集する枠を作る
	stmt, err := db.Prepare("UPDATE message SET messageContent = ?, edited = 1 WHERE messageId = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	// データをデータベースに挿入
	_, err = stmt.Exec(messageContent, messageId)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return nil, err
	}

	messages, err := MessageSearchDao(channelId)
	if err != nil {
		log.Printf("fail: MessageSearchDao, %v\n", err)
		return nil, err
	}

	return messages, nil
}

func MessageDeleteDao(messageId string, channelId string) ([]model.Message, error) {
	// データを編集する枠を作る
	stmt, err := db.Prepare("DELETE FROM message WHERE messageId = ?")
	if err != nil {
		log.Printf("fail: db.Prepare (Stmt), %v\n", err)
		return nil, err
	}
	defer stmt.Close()

	// データをデータベースに挿入
	_, err = stmt.Exec(messageId)
	if err != nil {
		log.Printf("fail: stmt.Exec, %v\n", err)
		return nil, err
	}

	messages, err := MessageSearchDao(channelId)
	if err != nil {
		log.Printf("fail: MessageSearchDao, %v\n", err)
		return nil, err
	}

	return messages, nil
}

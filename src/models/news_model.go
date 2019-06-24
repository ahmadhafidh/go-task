package models

import (
	"database/sql"
	"entities"
)

type NewsModel struct {
	Db *sql.DB
}

func (newsModel NewsModel) FindAll() (news []entities.News, err error) {
	rows, err := newsModel.Db.Query("select * from tb_news")
	if err != nil {
		return nil, err
	} else {
		var newsnew []entities.News
		for rows.Next() {
			var id int64
			var author string
			var body string
			var created string
			err2 := rows.Scan(&id, &author, &body, &created)
			if err2 != nil {
				return nil, err2
			} else {
				news := entities.News{
					Id:      id,
					Author:  author,
					Body:    body,
					Created: created,
				}
				newsnew = append(newsnew, news)
			}
		}
		return newsnew, nil
	}
}

// func (newsModel NewsModel) Search(keyword string) (news []entities.News, err error) {
// 	rows, err := newsModel.Db.Query("select * from tb_news where name like?", "%"+keyword+"%")
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		var newsnew []entities.News
// 		for rows.Next() {
// 			var id int64
// 			var author string
// 			var body string
// 			var created string
// 			err2 := rows.Scan(&id, &author, &body, &created)
// 			if err2 != nil {
// 				return nil, err2
// 			} else {
// 				news := entities.News{
// 					Id:      id,
// 					Author:  author,
// 					Body:    body,
// 					Created: created,
// 				}
// 				newsnew = append(newsnew, news)
// 			}
// 		}
// 		return newsnew, nil
// 	}
// }

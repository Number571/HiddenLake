package db

import (
	"github.com/number571/gopeer"
	"github.com/number571/hiddenlake/models"
	"github.com/number571/hiddenlake/settings"
)

func GetAllEmails(user *models.User) []models.Email {
	var (
		emails []models.Email
		incoming bool 
		lasttime string 
		hashname string
		public string 
		receiver string
		message string 
		salt string 
		hash string 
		sign string
		nonce uint64
	)
	id := GetUserId(user.Username)
	if id < 0 {
		return nil
	}
	rows, err := settings.DB.Query(
		"SELECT Incoming, LastTime, SenderHash, Sender, Receiver, Message, Salt, Hash, Sign, Nonce FROM Email WHERE IdUser=$1 AND Temporary=0 ORDER BY Id DESC",
		id,
	)
	if err != nil {
		panic("query 'getallemails' failed")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&incoming, &lasttime, &hashname, &public, &receiver, &message, &salt, &hash, &sign, &nonce)
		if err != nil {
			break
		}
		emails = append(emails, models.Email{
			Info: models.EmailInfo{
				Incoming: incoming,
				Time: lasttime,
			},
			Email: models.EmailType{
				Head: models.EmailHead{
					Sender: models.EmailSender{
						Public: public,
						Hashname: hashname,
					},
					Receiver: receiver,
				},
				Body: models.EmailBody{
					Data: string(gopeer.DecryptAES(user.Auth.Pasw, gopeer.Base64Decode(message))),
					Desc: models.EmailDesc{
						Rand: salt,
						Hash: hash,
						Sign: sign,
						Nonce: nonce,
						Difficulty: settings.DIFFICULTY,
					},
				},
			},
		})
	}
	return emails
}
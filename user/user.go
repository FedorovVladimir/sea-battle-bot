package user

import "sea-battle-bot/field"

type User struct {
	chatId int
	name   string
	field  *field.Field
}

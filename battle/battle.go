package battle

import (
	"sea-battle-bot/user"
)

type battles []*battle

type battle struct {
	user1 *user.User
	user2 *user.User
	step  *user.User
}

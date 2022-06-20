package usrsvc

import (
	"usersapi/usrdata"
)

var users = make(map[string]usrdata.User)

func Save(user usrdata.User) {
	users[user.Id] = user
}

func FindAll() []usrdata.User {
	usersSlice := make([]usrdata.User, 0, len(users))
	for _, currUser := range users {
		usersSlice = append(usersSlice, currUser)
	}
	return usersSlice
}

func FindOne(id string) usrdata.User {
	return users[id]
}

func Delete(id string) {
	delete(users, id)
}

package userssvc

import "github.com/shpboris/usersdata"

var users = make(map[string]usersdata.User)

func Save(user usersdata.User) {
	users[user.Id] = user
}

func FindAll() []usersdata.User {
	usersSlice := make([]usersdata.User, 0, len(users))
	for _, currUser := range users {
		usersSlice = append(usersSlice, currUser)
	}
	return usersSlice
}

func FindOne(id string) usersdata.User {
	return users[id]
}

func Delete(id string) {
	delete(users, id)
}

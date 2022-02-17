package data

type userProfile struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Subscribers  int    `json:"subscribers"`
	Subscription int    `json:"subscription"`
	Followed     bool   `json:"followed"`
	Avatar       string `json:"avatar"`
}

type allProfiles struct {
	AllUsersProfiles []userProfile `json:"allUsersProfiles"`
	CountUsers       int           `json:"countUsers"`
}

var UsersStore allProfiles = allProfiles{
	AllUsersProfiles: []userProfile{
		{
			Id:           5,
			Name:         "Пользователь 1",
			Subscribers:  1111,
			Subscription: 111,
			Followed:     true,
			Avatar:       "",
		},
	},
	CountUsers: 1,
}

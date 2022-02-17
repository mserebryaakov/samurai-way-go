package data

type messageData struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type dialogData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DialogPage struct {
	MessageData        []messageData `json:"messageData"`
	CurrentTextMessage string        `json:"currentTextMesage"`
	DialogData         []dialogData  `json:"dialogData"`
}

var DialogsStore DialogPage = DialogPage{
	MessageData: []messageData{
		{
			Id:      5,
			Message: "Сообщение 1",
		},
		{
			Id:      6,
			Message: "Сообщение 2",
		},
	},
	CurrentTextMessage: "",
	DialogData: []dialogData{
		{
			Id:   5,
			Name: "Пользователь 1",
		},
		{
			Id:   6,
			Name: "Пользователь 2",
		},
	},
}

package data

type messageData struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type message []messageData

type dialogData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type dialogs []dialogData

type DialogPage struct {
	MessageData        message `json:"messageData"`
	CurrentTextMessage string  `json:"currentTextMessage"`
	DialogData         dialogs `json:"dialogData"`
}

var Store DialogPage = DialogPage{
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

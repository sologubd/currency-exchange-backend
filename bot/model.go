package bot

type Chat struct {
	ID int `json:"id"`
}

type Message struct {
	ID   int    `json:"message_id"`
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}

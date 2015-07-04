package telegram

type ApiResult struct {
	Ok     bool   `json:"ok"`
	Result string `json:"result"`
}

type UpdateResult struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result"`
}

type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type User struct {
	Id        int
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type GroupChat struct {
	Id    int
	Title string
}

type PhotoSize struct {
	FileId        string `json:"file_id"`
	Width, Height int
	FileSize      int `json:"file_size"`
}

type Audio struct {
	FileId   string `json:"file_id"`
	Duration int
	MimeType string
	FileSize int
}

type Document struct {
	FileId   string `json:"file_id"`
	Thumb    *PhotoSize
	FileName string
	MimeType string
	FileSize int
}

type Sticker struct {
	FileId        string `json:"file_id"`
	Width, Height int
	Thumb         *PhotoSize
	FileSize      int
}

type Video struct {
	FileId        string `json:"file_id"`
	Width, Height int
	Duration      int
	Thumb         *PhotoSize
	MimeType      string
	FileSize      int
	Caption       string
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      string `json:"user_id"`
}

type Location struct {
	Longitude float64
	Latitude  float64
}

type UserProfilePhotos struct {
	TotalCount int
	Photos     [][]PhotoSize
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]string
	ResizeKeyboard  bool
	OneTimeKeyboard bool
	Selective       bool
}

type ReplyKeyboardHide struct {
	HideKeyboard bool
	Selective    bool
}

type ForceReply struct {
	ForceReply bool
	Selective  bool
}

type Message struct {
	MessageId           int `json:"message_id"`
	From                *User
	Date                int
	Chat                *interface{}
	ForwardFrom         *User    `json:"forward_from"`
	ForwardDate         int      `json:"forward_date"`
	ReplyToMessage      *Message `json:"reply_to_message"`
	Text                string
	Audio               *Audio
	Document            *Document
	Photo               []PhotoSize
	Sticker             *Sticker
	Video               *Video
	Contact             *Contact
	Location            *Location
	NewChatParticipant  *User       `json:"new_chat_participant"`
	LeftChatParticipant *User       `json:"left_chat_participant"`
	NewChatTitle        string      `json:"new_chat_title"`
	NewChatPhoto        []PhotoSize `json:"new_chat_photo"`
	DeleteChatPhoto     bool        `json:"delete_chat_photo"`
	GroupChatCreated    bool        `json:"group_chat_crated"`
}

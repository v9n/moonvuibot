package telegram

type UpdateResult struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
}

type GroupChat struct {
	Id    int
	Title string
}

type PhotoSize struct {
	FileId        string
	Width, Height int
	FileSize      int
}

type Audio struct {
	FileId   string
	Duration int
	MimeType string
	FileSize int
}

type Document struct {
	FileId   string
	Thumb    *PhotoSize
	FileName string
	MimeType string
	FileSize int
}

type Sticker struct {
	FileId        string
	Width, Height int
	Thumb         *PhotoSize
	FileSize      int
}

type Video struct {
	FileId        string
	Width, Height int
	Duration      int
	Thumb         *PhotoSize
	MimeType      string
	FileSize      int
	Caption       string
}

type Contact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	UserId      string
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
	MessageId           int
	From                *User
	Date                int
	Chat                *interface{}
	ForwardFrom         *User
	ForwardDate         int
	ReplyToMessage      *Message
	Text                string
	Audio               *Audio
	Document            *Document
	Photo               []PhotoSize
	Sticker             *Sticker
	Video               *Video
	Contact             *Contact
	Location            *Location
	NewChatParticipant  *User
	LeftChatParticipant *User
	NewChatTitle        string
	NewChatPhoto        []PhotoSize
	DeleteChatPhoto     bool
	GroupChatCreated    bool
}

package domain

import "time"

type Tweet interface {
	GetText() string
	GetUser() string
	GetDate() *time.Time
	GetId() int

	Printable() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
	Id   int
}

type ImageTweet struct {
	TextTweet
	Url string
}

type QuoteTweet struct {
	TextTweet
	Quoted Tweet
}

func (t TextTweet) GetText() string {
	return t.Text
}

func (t TextTweet) GetUser() string {
	return t.User
}

func (t TextTweet) GetDate() *time.Time {
	return t.Date
}

func (t TextTweet) GetId() int {
	return t.Id
}

func (t ImageTweet) GetUrl() string {
	return t.Url
}
func NewTextTweet(user string, text string) *TextTweet {
	var date = time.Now()
	var creado = TextTweet{User: user, Text: text, Date: &date}
	return &creado
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	var date = time.Now()
	var creado = ImageTweet{TextTweet: TextTweet{User: user, Text: text, Date: &date}, Url: url}
	return &creado
}

func NewQuoteTweet(user string, text string, quoted Tweet) *QuoteTweet {
	var date = time.Now()
	var creado = QuoteTweet{TextTweet: TextTweet{User: user, Text: text, Date: &date}, Quoted: quoted}
	return &creado
}

func (tweet *TextTweet) Printable() string {
	printable := ("@" + tweet.User + ": " + tweet.Text)

	return printable
}

func (tweet *ImageTweet) Printable() string {
	printable := ("@" + tweet.User + ": " + tweet.Text + " " + tweet.Url)

	return printable
}

func (tweet *QuoteTweet) Printable() string {
	printable := ("@" + tweet.User + ": " + tweet.Text + ` "` + tweet.Quoted.Printable() + `"`)

	return printable
}

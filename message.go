package chainmail

import (
	"time"
	"net/mail"
)

type Message struct{
	*mail.Message
	Received time.Time
	Date time.Time
	From []mail.Address
	Subject string
	Sender mail.Address
	To []mail.Address
	Cc []mail.Address
	Bcc []mail.Address
}

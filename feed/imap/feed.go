package imap

import (
  "net/mail"
  "github.com/joshwillik/chainmail"
  "github.com/emersion/go-imap"
  "github.com/emersion/go-imap/client"
)

type Feed struct{
  Address string
  Username string
  Password string
  Mailboxes []string
  client *client.Client
  messages chan chainmail.Message
}

func (f Feed) Init() error {
  client, err := client.DialTLS(f.Address, nil)
  if err != nil {
    return err
  }
  f.client = client
  if len(f.Mailboxes) == 0 {
    boxes := make(chan *imap.MailboxInfo, 10)
    done := make(chan error, 1)
    go func(){
      done <- f.client.List("", "*", boxes)
    }()
    for box := range boxes {
      f.Mailboxes = append(f.Mailboxes, box.Name)
    }
    if err := <-done; err != nil {
      return err
    }
  }
  return nil
}

func (f Feed) Open(m chan chainmail.Message) error {
  f.messages = m
  for _, name := range f.Mailboxes {
    if err := f.syncBox(name); err != nil {
      return err
    }
  }
  return nil
}

func (f Feed) syncBox(name string) error {
  _, err := f.client.Select(name, true)
  if err != nil {
    return err
  }
  messages := make(chan *imap.Message, 10)
  seq := new(imap.SeqSet)
  //seq.AddRange(1, box.Messages) // only on plentiful internet
  seq.AddRange(1, 10)
  done := make(chan error, 1)
  go func(){
    done <- f.client.Fetch(seq, []imap.FetchItem{imap.FetchEnvelope}, messages)
  }()
  for message := range messages {
    f.messages <- parseMessage(message)
  }
  if err := <-done; err != nil {
    return err
  }
  return nil
}

func (f Feed) Close(){
  f.client.Logout()
  close(f.messages)
}

func parseMessage(m *imap.Message) chainmail.Message {
  e := m.Envelope
  return chainmail.Message{
    MessageId: e.MessageId,
    InReplyTo: e.InReplyTo,
    Received: m.InternalDate,
    Date: e.Date,
    From: toAddresses(e.From),
    Subject: e.Subject,
    Sender: toAddresses(e.Sender)[0],
    To: toAddresses(e.To),
    Cc: toAddresses(e.Cc),
    Bcc: toAddresses(e.Bcc),
  }
}

func toAddresses(in []*imap.Address) []mail.Address {
  out := []mail.Address{}
  for _, address := range in {
    out = append(out, mail.Address{
      Name: address.PersonalName,
      Address: address.MailboxName + "@" + address.HostName,
    })
  }
  return out
}

package imap

import (
  "fmt"
  "os"
  "testing"
  "github.com/joshwillik/chainmail"
)

func TestFeed(t *testing.T){
  server := os.Getenv("IMAP_SERVER")
  username := os.Getenv("IMAP_USERNAME")
  password := os.Getenv("IMAP_PASSWORD")
  feed := Feed{
    Address: server,
    Username: username,
    Password: password,
    Mailboxes: []string{"INBOX"},
  }
  fmt.Println("Init feed")
  if err := feed.Init(); err != nil {
    t.Error(err)
    t.FailNow()
  }
  messages := make(chan chainmail.Message, 10)
  done := make(chan error, 1)
  go func(){
    fmt.Println("Opening feed")
    done <- feed.Open(messages)
    fmt.Println("Opened feed")
  }()
  fmt.Println("Iterate messages")
  // TODO josh: validate that it actually reads messages
  if err := <-done; err != nil {
    t.Error(err)
    t.FailNow()
  }
}

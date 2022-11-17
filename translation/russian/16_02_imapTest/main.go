// IMAP connect to GMAIL
package main

import (
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	log.Println("Connecting to server...")

	// Подключение к серверу
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Отложенное отключение
	defer c.Logout()

	// Аутотентификация
	if err := c.Login("gotestsmtp@gmail.com", "PASS"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// Все ящики
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Только ВХОДЯЩИЕ
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("Flags for INBOX:", mbox.Flags)

	// Получение последнего сообщения
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(mbox.Messages, mbox.Messages)

	// Получеие всех данных
	attrs := []string{"BODY[]"}

	messages := make(chan *imap.Message, 1)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, attrs, messages)
	}()

	log.Println("Last message:")
	msg := <-messages
	r := msg.GetBody("BODY[]")
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	log.Println("Date:", header.Get("Date"))
	log.Println("From:", header.Get("From"))
	log.Println("To:", header.Get("To"))
	log.Println("Subject:", header.Get("Subject"))

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

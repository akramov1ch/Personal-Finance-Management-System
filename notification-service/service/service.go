package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	notifactionpb "notifaction-service/protos/notifaction"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
)

type NotificationServiceServer struct {
	notifactionpb.UnimplementedNotificationServiceServer
	Consumer *kafka.Consumer
	clients  map[*websocket.Conn]bool // WebSocket client connections
	mu       sync.Mutex               // Mutex for concurrent access to clients map
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewNotificationServer(cons *kafka.Consumer) *NotificationServiceServer {
	ns := &NotificationServiceServer{
		Consumer: cons,
		clients:  make(map[*websocket.Conn]bool),
	}
	return ns
}

func (s *NotificationServiceServer) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket connection closed:", err)
			s.mu.Lock()
			delete(s.clients, conn)
			s.mu.Unlock()
			break
		}
	}
}

func (s *NotificationServiceServer) ListenForKafkaMessages() {
	for {
		msg, err := s.Consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Olingan Kafka xabari: %s\n", string(msg.Value))

			var emailMsg notifactionpb.NotificationReq
			err := json.Unmarshal(msg.Value, &emailMsg)
			if err != nil {
				log.Printf("Xabarni parsing qilishda xatolik: %v", err)
				continue
			}

			err = SendEmail(emailMsg.UserEmail, emailMsg.Subject, emailMsg.Message)
			if err != nil {
				log.Printf("Email yuborishda xatolik: %v", err)
			}

			s.broadcastWebSocketMessage(fmt.Sprintf("Email: %s\nSubject: %s\nMessage: %s", emailMsg.UserEmail, emailMsg.Subject, emailMsg.Message))
		} else {
			log.Printf("Xabar o'qishda xatolik: %v\n", err)
		}
	}
}

func (s *NotificationServiceServer) broadcastWebSocketMessage(message string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Xabarni WebSocket orqali yuborishda xatolik:", err)
			client.Close()
			delete(s.clients, client)
		}
	}
}

func SendEmail(to, subject, body string) error {
	from := "dilshoddilmurodov112@gmail.com"
	pass := "xmxu rdhp pmdf pezk"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}

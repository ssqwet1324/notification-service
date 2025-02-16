package service

import (
	"context"
	"fmt"
	"log"
	"service_notification/internal/email"
	"service_notification/internal/repository"
	"service_notification/internal/telegram"
)

type Service struct {
	Tg    telegram.Telegram
	Email email.Email
	Repo  repository.Repository
}

func NewService(tg telegram.Telegram, email email.Email, repo repository.Repository) *Service {
	return &Service{tg, email, repo}
}

func (s *Service) SendNotifications(ctx context.Context, userID int, message string) error {
	userInfo, err := s.Repo.GetUserInfo(ctx, userID)
	if err != nil {
		return fmt.Errorf("GetNotificationByID: %w", err)
	}

	if userInfo.EmailRecipient != "" {
		err := s.Email.Send(userInfo.EmailRecipient, "New notification", message)
		if err != nil {
			log.Println("Error sending email", err)
		} else {
			log.Printf("Notification sent to %s", userInfo.EmailRecipient)
		}
	}

	if userInfo.TelegramUserID != 0 {
		err = s.Tg.Send(int64(userInfo.TelegramUserID), message)
		if err != nil {
			log.Printf("Error sending notification to telegram")
		} else {
			log.Printf("Notification sent successfully %d", userInfo.TelegramUserID)
		}
	}
	return nil
}

package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"service_notification/internal/config"
)

// Repository отвечает за работу с БД
type Repository struct {
	DB *pgxpool.Pool
}

// UserInfoDBO — структура для хранения данных о пользователе
type UserInfoDBO struct {
	EmailRecipient string
	TelegramUserID int
}

func NewRepository(cfg *config.Config) *Repository {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DbName, cfg.DbPassword, cfg.DbHost, cfg.DbPort, "postgres")
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	return &Repository{dbPool}
}

// GetUserInfo получает информацию о пользователе по его ID
func (r *Repository) GetUserInfo(ctx context.Context, userID int) (UserInfoDBO, error) {
	var userInfo UserInfoDBO

	err := r.DB.QueryRow(
		ctx,
		`SELECT email_recipient, telegram_user_id FROM UserInfo WHERE id = $1`,
		userID,
	).Scan(&userInfo.EmailRecipient, &userInfo.TelegramUserID)

	if err != nil {
		log.Println("Ошибка получения информации о пользователе:", err)
		return UserInfoDBO{}, fmt.Errorf("ошибка получения информации о пользователе: %w", err)
	}

	return userInfo, nil
}

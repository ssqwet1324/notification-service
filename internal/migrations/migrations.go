package migrations

import (
	"context"
	"fmt"
	"service_notification/internal/config"
	"service_notification/internal/repository"
)

// Migration отвечает за выполнение миграций
type Migration struct {
	repo *repository.Repository
}

// NewMigration создаёт новый объект для миграций
func NewMigration(repo *repository.Repository) *Migration {
	return &Migration{repo: repo}
}

// InitTables создаёт таблицу, если её нет
func (m *Migration) InitTables(ctx context.Context, cfg *config.Config) error {
	query := `CREATE TABLE IF NOT EXISTS UserInfo (
		id SERIAL PRIMARY KEY, 
		email_recipient VARCHAR(255), 
		telegram_user_id BIGINT
	);`

	_, err := m.repo.DB.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating table UserInfo: %w", err)
	}

	// Вставка тестовых данных корректным способом
	query = `INSERT INTO UserInfo (email_recipient, telegram_user_id) VALUES ($1, $2);`
	_, err = m.repo.DB.Exec(ctx, query, cfg.TestUserMail, cfg.TelegramChatID)
	if err != nil {
		return fmt.Errorf("error inserting into UserInfo: %w", err)
	}

	return nil
}

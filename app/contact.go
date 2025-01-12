package app

import (
	"gorm.io/gorm"
	"sync"
	"taha_tahvieh_tg_bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	faqPort "taha_tahvieh_tg_bot/internal/faq/port"
	settingsPort "taha_tahvieh_tg_bot/internal/settings/port"
)

type App interface {
	DB() *gorm.DB
	Config() config.Config
	Bot() *tgbotapi.BotAPI
	AppState(id int64) *UserState
	ResetUserState(id int64)
	DeleteUserState(id int64)
	SettingsService() settingsPort.Service
	FaqService() faqPort.Service
}

type appState struct {
	userStates map[int64]*UserState
	mutex      *sync.Mutex
}

type UserState struct {
	Conversation string
	Step         int
	Data         map[string]string
	Active       bool
}

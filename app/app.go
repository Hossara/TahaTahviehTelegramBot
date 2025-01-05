package app

import (
	"context"
	"gorm.io/gorm"
	"sync"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/internal/settings"
	"taha_tahvieh_tg_bot/pkg/adapters/storage"
	"taha_tahvieh_tg_bot/pkg/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	settingsPort "taha_tahvieh_tg_bot/internal/settings/port"
)

type app struct {
	cfg             config.Config
	ctx             context.Context
	db              *gorm.DB
	bot             *tgbotapi.BotAPI
	state           *appState
	settingsService settingsPort.Service
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Bot() *tgbotapi.BotAPI {
	return a.bot
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) setAppState() {
	a.state = &appState{
		userStates: make(map[int64]*UserState),
		mutex:      &sync.Mutex{},
	}
}

func (a *app) AppState(id int64) *UserState {
	a.state.mutex.Lock()
	defer a.state.mutex.Unlock()

	if state, exists := a.state.userStates[id]; exists {
		return state
	}

	state := &UserState{Step: 0, Data: make(map[string]string), Active: false}
	a.state.userStates[id] = state
	return state
}

func (a *app) DeleteUserState(id int64) {
	a.state.mutex.Lock()
	defer a.state.mutex.Unlock()
	delete(a.state.userStates, id)
}

func (a *app) ResetUserState(id int64) {
	a.state.mutex.Lock()
	defer a.state.mutex.Unlock()

	if _, exists := a.state.userStates[id]; exists {
		delete(a.state.userStates, id)
	}

	state := &UserState{Step: 0, Data: make(map[string]string), Active: false}
	a.state.userStates[id] = state
}

func (a *app) SettingsService() settingsPort.Service {
	if a.settingsService == nil {
		if a.db != nil {
			a.settingsService = settings.NewService(a.ctx, storage.NewSettingRepo(a.db))

			if err := a.settingsService.RunMigrations(); err != nil {
				panic("failed to run migrations")
			}

			return a.settingsService
		}

		return nil
	}

	return a.settingsService
}

func (a *app) setBot() error {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   a.cfg.DB.Host,
		Port:   a.cfg.DB.Port,
		User:   a.cfg.DB.User,
		Pass:   a.cfg.DB.Pass,
		Name:   a.cfg.DB.Name,
		Schema: a.cfg.DB.Schema,
	})

	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func NewApp(ctx context.Context, cfg config.Config, bot *tgbotapi.BotAPI) (App, error) {
	a := &app{cfg: cfg, bot: bot, ctx: ctx}

	a.setAppState()

	if err := a.setBot(); err != nil {
		return nil, err
	}

	return a, nil
}

func MustNewApp(ctx context.Context, cfg config.Config, bot *tgbotapi.BotAPI) App {
	a, err := NewApp(ctx, cfg, bot)
	if err != nil {
		panic(err)
	}
	return a
}

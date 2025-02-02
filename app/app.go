package app

import (
	"context"
	"gorm.io/gorm"
	"sync"
	"taha_tahvieh_tg_bot/config"
	"taha_tahvieh_tg_bot/internal/faq"
	"taha_tahvieh_tg_bot/internal/product"
	"taha_tahvieh_tg_bot/internal/settings"
	"taha_tahvieh_tg_bot/internal/storage"
	"taha_tahvieh_tg_bot/pkg/adapters/database"
	"taha_tahvieh_tg_bot/pkg/minio"
	"taha_tahvieh_tg_bot/pkg/postgres"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	faqPort "taha_tahvieh_tg_bot/internal/faq/port"
	productPort "taha_tahvieh_tg_bot/internal/product/port"
	settingsPort "taha_tahvieh_tg_bot/internal/settings/port"
	storagePort "taha_tahvieh_tg_bot/internal/storage/port"
	storageAdapter "taha_tahvieh_tg_bot/pkg/adapters/storage"
)

type app struct {
	cfg   config.Config
	ctx   context.Context
	db    *gorm.DB
	bot   *tgbotapi.BotAPI
	state *appState

	faqService      faqPort.Service
	productService  productPort.Service
	settingsService settingsPort.Service
	storageService  storagePort.Service
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

func (a *app) ProductService() productPort.Service {
	if a.productService == nil {
		a.productService = product.NewService(
			a.ctx, database.NewProductRepo(a.db), database.NewProductMetaRepo(a.db),
		)

		if err := a.productService.RunProductMigrations(); err != nil {
			panic("failed to run migrations for product service!")
		}

		if err := a.productService.RunBrandMigrations(); err != nil {
			panic("failed to run migrations for product brand service!")
		}

		if err := a.productService.RunProductTypeMigrations(); err != nil {
			panic("failed to run migrations for product type service!")
		}

		return a.productService
	}

	return a.productService
}

func (a *app) StorageService() storagePort.Service {
	return a.storageService
}

func (a *app) setStorageService() {
	if a.storageService == nil {
		minioConfig := a.cfg.Minio

		client := storageAdapter.NewStorageRepo(minio.Config{
			Endpoint:        minioConfig.Endpoint,
			AccessKeyID:     minioConfig.AccessKeyID,
			SecretAccessKey: minioConfig.SecretAccessKey,
			SSL:             minioConfig.SSL,
		})

		a.storageService = storage.NewService(a.ctx, database.NewStorageRepo(a.db), client)

		if err := a.storageService.RunMigrations(); err != nil {
			panic("failed to run migrations for storage service!")
		}

		err := a.storageService.InitBucket("products")

		if err != nil {
			panic(fmt.Errorf("failed to initialize storage bucket: %s", err))
		}
	}
}

func (a *app) SettingsService() settingsPort.Service {
	if a.settingsService == nil {
		a.settingsService = settings.NewService(a.ctx, database.NewSettingRepo(a.db))

		if err := a.settingsService.RunMigrations(); err != nil {
			panic("failed to run migrations for faq setting!")
		}

		return a.settingsService
	}

	return a.settingsService
}

func (a *app) FaqService() faqPort.Service {
	if a.faqService == nil {
		a.faqService = faq.NewService(a.ctx, database.NewFaqRepo(a.db))

		if err := a.faqService.RunMigrations(); err != nil {
			panic("failed to run migrations for faq service!")
		}

		return a.faqService
	}

	return a.faqService
}

func (a *app) setDB() error {
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

	if err := a.setDB(); err != nil {
		return nil, err
	}

	a.setAppState()
	a.setStorageService()

	return a, nil
}

func MustNewApp(ctx context.Context, cfg config.Config, bot *tgbotapi.BotAPI) App {
	a, err := NewApp(ctx, cfg, bot)
	if err != nil {
		panic(err)
	}
	return a
}

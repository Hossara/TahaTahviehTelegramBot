package mapper

import (
	"encoding/json"
	"taha_tahvieh_tg_bot/internal/settings/domain"
	"taha_tahvieh_tg_bot/pkg/adapters/storage/models"
	"time"
)

func ToDomainSetting(m *models.Setting) *domain.Setting {
	if m == nil {
		return nil
	}

	var content = domain.Content{}

	err := json.Unmarshal([]byte(m.Content), &content)

	if err != nil {
		return nil
	}

	return &domain.Setting{
		SettingID: domain.SettingID(m.ID),
		Title:     m.Title,
		Content:   content,
	}
}

func ToModelSetting(m *domain.Setting) *models.Setting {
	if m == nil {
		return nil
	}

	content, err := json.Marshal(m.Content)
	if err != nil {
		return nil
	}

	return &models.Setting{
		ID:        uint8(m.SettingID),
		Title:     m.Title,
		Content:   string(content),
		UpdatedAt: time.Now(),
	}
}

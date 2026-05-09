package service

import (
	"encoding/json"
	"strconv"
	"time"

	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
)

const configCacheKey = "zyqhome:config:public"

type ConfigService struct {
	repo *repository.ConfigRepo
}

func NewConfigService(repo *repository.ConfigRepo) *ConfigService {
	return &ConfigService{repo: repo}
}

type PublicConfig struct {
	SiteTitle         string `json:"site_title"`
	SiteDescription   string `json:"site_description"`
	BackgroundImage   string `json:"background_image"`
	BackgroundColor   string `json:"background_color"`
	DefaultSearch     string `json:"default_search"`
	SearchPlaceholder string `json:"search_placeholder"`
	FooterText        string `json:"footer_text"`
	ShowCenterIcons   bool   `json:"show_center_icons"`
	ShowRightMenu     bool   `json:"show_right_menu"`
	IconColumns       int    `json:"icon_columns"`
}

func defaultPublicConfig() *PublicConfig {
	return &PublicConfig{
		SiteTitle:         "ZyqHome",
		SiteDescription:   "我的导航",
		BackgroundImage:   "",
		BackgroundColor:   "#f0f2f5",
		DefaultSearch:     "google",
		SearchPlaceholder: "搜索...",
		FooterText:        "",
		ShowCenterIcons:   true,
		ShowRightMenu:     true,
		IconColumns:       5,
	}
}

func (s *ConfigService) GetPublicConfig() (*PublicConfig, error) {
	cached, err := database.RDB.Get(database.Ctx, configCacheKey).Result()
	if err == nil && cached != "" {
		var cfg PublicConfig
		if json.Unmarshal([]byte(cached), &cfg) == nil {
			return &cfg, nil
		}
	}

	configs, err := s.repo.FindAll()
	if err != nil {
		return defaultPublicConfig(), nil
	}

	cfg := defaultPublicConfig()
	for _, c := range configs {
		applyConfig(cfg, c.ConfigKey, c.ConfigValue)
	}

	if data, err := json.Marshal(cfg); err == nil {
		database.RDB.Set(database.Ctx, configCacheKey, string(data), time.Hour)
	}

	return cfg, nil
}

func (s *ConfigService) GetAllConfigs() ([]model.SiteConfig, error) {
	return s.repo.FindAll()
}

func (s *ConfigService) UpdateConfigs(configs map[string]string) error {
	for key, value := range configs {
		if err := s.repo.UpsertByKey(key, value); err != nil {
			return err
		}
	}
	database.RDB.Del(database.Ctx, configCacheKey)
	return nil
}

func applyConfig(cfg *PublicConfig, key, value string) {
	switch key {
	case "site_title":
		cfg.SiteTitle = value
	case "site_description":
		cfg.SiteDescription = value
	case "background_image":
		cfg.BackgroundImage = value
	case "background_color":
		cfg.BackgroundColor = value
	case "default_search":
		cfg.DefaultSearch = value
	case "search_placeholder":
		cfg.SearchPlaceholder = value
	case "footer_text":
		cfg.FooterText = value
	case "show_center_icons":
		cfg.ShowCenterIcons, _ = strconv.ParseBool(value)
	case "show_right_menu":
		cfg.ShowRightMenu, _ = strconv.ParseBool(value)
	case "icon_columns":
		cfg.IconColumns, _ = strconv.Atoi(value)
	}
}

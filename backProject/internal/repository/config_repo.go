package repository

import (
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
)

type ConfigRepo struct{}

func NewConfigRepo() *ConfigRepo {
	return &ConfigRepo{}
}

func (r *ConfigRepo) FindAll() ([]model.SiteConfig, error) {
	var configs []model.SiteConfig
	err := database.DB.Find(&configs).Error
	return configs, err
}

func (r *ConfigRepo) FindByKey(key string) (*model.SiteConfig, error) {
	var config model.SiteConfig
	err := database.DB.Where("config_key = ?", key).First(&config).Error
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *ConfigRepo) Upsert(config *model.SiteConfig) error {
	var existing model.SiteConfig
	result := database.DB.Where("config_key = ?", config.ConfigKey).First(&existing)
	if result.Error != nil {
		return database.DB.Create(config).Error
	}
	config.ID = existing.ID
	return database.DB.Save(config).Error
}

func (r *ConfigRepo) UpsertByKey(key, value string) error {
	return database.DB.Where("config_key = ?", key).Assign(model.SiteConfig{ConfigValue: value}).FirstOrCreate(&model.SiteConfig{}).Error
}

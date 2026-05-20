package cache

import "github.com/Citrus0974/ProxyProject/internal/models"

type Service interface {
	Get(key string) (*models.Entry, bool)
	Set(key string, entry *models.Entry)
}

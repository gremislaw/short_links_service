package repository

import (
	"runtime"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type cacheItem struct {
	value  string
	expiry time.Time
}

func (item *cacheItem) IsExpired() bool {
	return time.Now().After(item.expiry)
}

func cleanupCache(mapa1, mapa2 *sync.Map) {
	mapa1.Range(func(key, value interface{}) bool {
		// Преобразуем key и value в нужные типы
		v := value.(cacheItem)
		if v.IsExpired() {
			mapa1.Delete(key)
			mapa2.Delete(v.value)
			logrus.Infof("%v -> %v expired and deleted from memory: ", key, v)
		}
		return true
	})
}

func clearMapIfMemoryExceeded(mapa1, mapa2 *sync.Map) {
	// Получаем текущие данные о памяти
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Проверяем, если используемая память больше половины от всей доступной
	// Если используемая память больше 50% от доступной
	if memStats.Alloc > memStats.Sys/2 {
		// Если память превышает порог, очищаем мапу
		mapa1.Range(func(key, value interface{}) bool {
			v := value.(cacheItem).value
			mapa1.Delete(key)
			mapa2.Delete(v)
			logrus.Infof("%v -> %v deleted from memory: ", key, v)
			return true
		})
		logrus.Info("Cache cleared due to high memory usage.")
	}
}

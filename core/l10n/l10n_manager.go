package l10n

import (
	"encoding/json"
	"sync"
	"crypto/md5"
	"fmt"
)

// L10nManager 管理本地化静态文本
type L10nManager struct {
	mu       sync.Mutex
	textMap  map[string]string // key -> raw text
}

var instance *L10nManager
var once sync.Once

func GetManager() *L10nManager {
	once.Do(func() {
		instance = &L10nManager{
			textMap: make(map[string]string),
		}
	})
	return instance
}

// AddText 收集一个 text，如果没有指定 key，根据内容生成一个
func (m *L10nManager) AddText(key string, text string) string {
	if text == "" {
		return ""
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()

	if key == "" {
		hash := md5.Sum([]byte(text))
		key = fmt.Sprintf("TEXT_%x", hash)
	}

	m.textMap[key] = text
	return key
}

// Export 导出所有收集到的多语言 key-value 映射为 JSON 格式
func (m *L10nManager) Export() ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	return json.MarshalIndent(m.textMap, "", "  ")
}

// Clear 清空收集记录
func (m *L10nManager) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.textMap = make(map[string]string)
}
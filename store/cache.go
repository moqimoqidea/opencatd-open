package store

import (
	"log"
	"math/rand"

	"github.com/Sakurasan/to"
	"github.com/patrickmn/go-cache"
)

var (
	KeysCache *cache.Cache
	AuthCache *cache.Cache
)

func init() {
	KeysCache = cache.New(cache.NoExpiration, cache.NoExpiration)
	AuthCache = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func LoadKeysCache() {
	keys, err := GetAllKeys()
	if err != nil {
		log.Println(err)
		return
	}
	for _, key := range keys {
		KeysCache.Set(key.Key, true, cache.NoExpiration)
	}
}

func FromKeyCacheRandomItem() string {
	items := KeysCache.Items()
	idx := rand.Intn(len(items))
	item := items[to.String(idx)]
	return item.Object.(string)
}

func LoadAuthCache() {
	users, err := GetAllUsers()
	if err != nil {
		log.Println(err)
		return
	}
	for _, user := range users {
		AuthCache.Set(user.Token, true, cache.NoExpiration)
	}
}

func IsExistAuthCache(auth string) bool {
	items := AuthCache.Items()
	_, ok := items[auth]
	return ok
}

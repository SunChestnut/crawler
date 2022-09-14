package persist

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

type ItemSaverService struct {
	Client *elasticsearch.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item %v saved.", item)
	if err != nil {
		log.Printf("Error saving item %v: %v", item, err)
		return err
	}
	*result = "ok"
	return nil
}

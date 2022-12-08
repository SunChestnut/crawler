package client

import (
	"context"
	"crawler/grpcsupport"
	"crawler/model"
	"crawler/pb"
	"log"
)

func StartItemSaverClient(address string) (chan model.Item, error) {
	log.Println("💫ItemSaver client is running...")

	grpcClient := grpcsupport.NewItemSaverClient(address)
	ctx := context.Background()
	out := make(chan model.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("[ItemSaver] got item #%d: %v", itemCount, item)
			itemCount++

			// TODO
			pbItem := engineItemToPbItem(item)

			_, err := grpcClient.Save(ctx, &pb.ItemSaverRequest{Item: pbItem})
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v\n", item, err)
			}
		}
	}()
	return out, nil
}

func engineItemToPbItem(item model.Item) *pb.Item {
	switch p := item.Payload.(type) {
	case pb.Profile:
		return &pb.Item{
			Url:     item.Url,
			Id:      item.Id,
			Type:    item.Type,
			Profile: &p,
		}
	}
	return nil
}

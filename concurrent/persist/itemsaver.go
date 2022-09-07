package persist

import "log"

func ItemSaver() chan any {
	out := make(chan any)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("【Item Saver】Got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}

package model

import "crawler/concurrent/engine"

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}

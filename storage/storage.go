package storage

type Service interface {
	Add(string) (*Item, error)     // Add new url into storage
	Visit(string) (string, error)  // Visit id once
	GetInfo(string) (*Item, error) // Get info by id
	Close() error                  // Close storage connection
}

type Item struct {
	ID        string `json:"id" redis:"id"`               // ID, unique key
	URL       string `json:"url" redis:"url"`             // Full url of this item
	CreatedAt int64  `json:"createdat" redis:"createdat"` // Creation timestamp
	Visits    int64  `json:"visits" redis:"visits"`       // Visit count
}

package entity

type Commerce struct {
	ID      string
	Name    string
	Address string
}

type Branch struct {
	ID         string
	CommerceID string
	Name       string
	Address    string
}

package data

// Unit ...
type Unit struct {
	ID        string        `json:"id"`
	Localized LocalizedUnit `json:"localized"`
}

// LocalizedUnit ...
type LocalizedUnit struct {
	Name string
	Hype string
	Lore string
}

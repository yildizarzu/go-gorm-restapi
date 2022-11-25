package models

// swagger:model
type Ticket struct {
	ID         int64  `json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Desc       string `gorm:"not null" json:"desc"`
	Allocation int    `gorm:"not null" json:"allocation"`
}

// swagger:model
type Ticket_Purchase struct {
	Quantity int    `json:"quantity"`
	UserID   string `json:"user_id"`
}

// swagger:model
type TicketCreate struct {
	Name       string `gorm:"not null" json:"name"`
	Desc       string `gorm:"not null" json:"desc"`
	Allocation int    `gorm:"not null" json:"allocation"`
}

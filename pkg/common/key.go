package common

import "time"

// Key represents a hardened key used for one-time encrypted
// communication with a user.
type Key struct {
	ID        ID        `gorm:"primaryKey" json:"-"` // The key identifier.
	CreatedAt time.Time `gorm:"" json:"createdAt"`   // When the key was created.
	UpdatedAt time.Time `gorm:"" json:"updatedAt"`   // The last time the key was updated.
	UserID    ID        `gorm:"index" json:"-"`      // The user identifier.

	Name      string `gorm:"" json:"name"`            // A name for the key or other information.
	PublicKey string `gorm:"unique" json:"publicKey"` // Hexadecimal encoded key.
	Used      bool   `gorm:"" json:"used"`            // Has the key been returned in a /lookup request by another user.
}

// TableName for managing the GORM generated table name.
func (Key) TableName() string {
	return "keys"
}

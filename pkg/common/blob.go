package common

import "time"

// Blob is the representation of an entry in Depot that is
// associated with a user.
type Blob struct {
	ID        ID        `gorm:"primaryKey" json:"-"` // The blob identifier.
	CreatedAt time.Time `gorm:"" json:"createdAt"`   // When the blob was first uploaded.
	UpdatedAt time.Time `gorm:"" json:"updatedAt"`   // The last time the blob was updated.
	UserID    ID        `gorm:"index" json:"-"`      // The user identifier.

	Data     string `gorm:"" json:"data"`      // Encoded data.
	Metadata string `gorm:"" json:"metadata"`  // Information about the data or encoding.
	Name     string `gorm:"index" json:"name"` // A name for the blob.
}

// TableName for managing the GORM generated table name.
func (Blob) TableName() string {
	return "blobs"
}

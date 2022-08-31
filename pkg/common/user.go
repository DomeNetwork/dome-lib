package common

import "time"

// User is the internal account for DOME.  All interactions and usage
// of the platform require valid users.
type User struct {
	ID        ID        `gorm:"primaryKey" json:"-"` // The UUIDv4 identifier for the user.
	CreatedAt time.Time `gorm:"" json:"createdAt"`   // When the the user was created.
	UpdatedAt time.Time `gorm:"" json:"updatedAt"`   // The last time the user was updated.

	Domain    string `gorm:"index" json:"domain"`          // The domain of the user.
	Email     string `gorm:"index;unique" json:"email"`    // The verified email of the user.
	Password  string `gorm:"" json:"password"`             // The stored hash of the password.
	PublicKey string `gorm:"" json:"publicKey"`            // Hexadecimal encoded root identity key.
	Username  string `gorm:"index;unique" json:"username"` // The handle or username.
}

// TableName for managing the GORM generated table name.
func (User) TableName() string {
	return "users"
}

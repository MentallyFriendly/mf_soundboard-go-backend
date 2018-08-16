package db

import "time"

// Sound type
type Sound struct {
	ID        uint       `json:"id"`
	Name      *string    `json:"name" sql:"not null"`
	Path      *string    `json:"path"`
	Letter    *string    `json:"letter"`
	Emoji     *string    `json:"emoji" sql:"default:'U+1F525'"`
	Groups    []Group    `json:"groups" gorm:"many2many:sound_groups"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// Group type
type Group struct {
	ID        uint       `json:"id"`
	Name      *string    `json:"name" sql:"not null"`
	Sounds    []Sound    `json:"sounds" gorm:"many2many:sound_groups"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

package database

import (
	"github.com/Kokosalah45/share-my-notes/internal/database/models"
)



type UserCRUD interface {
	CreateUser(user *models.User) error
	GetUser(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type NotesCRUD interface {
	CreateNote(note *models.Note) error
	GetNoteByID(hexID string) (*models.Note, error)
	ListNodes(userID int) ([]*models.Note, error)
	UpdateNoteByID(hexID string, note *models.Note) error
	DeleteNoteByID(hexID string) error
	GetSharedNotes(userID, friendID int ) ([]*models.Note, error)
}

type dbConfig struct {
	Username string
	DatabaseName string
	Host string
	Password string
}



type storer interface {
	InitDB(dbConfig *dbConfig)
	getDBConnection() (*Idb, error)
}

type Idb interface {
	UserCRUD
	NotesCRUD
}

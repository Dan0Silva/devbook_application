package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id         string    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uuid.UUID `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      int32     `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

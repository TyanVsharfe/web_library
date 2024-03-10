package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	Register(ctx context.Context, login, hashPassword string) (string, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}

type BookRepository interface {
	CreateBook(ctx context.Context, book model.Book) (int, error)
	DeleteBook(ctx context.Context, bookId int) error
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, error)
	UpdateBook(ctx context.Context, book model.Book, bookId int) (model.Book, error)
}

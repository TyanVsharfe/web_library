package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	Register(ctx context.Context, login, hashPassword string) (string, error)
}

type BookRepository interface {
	CreateBook(ctx context.Context, book model.Book) (int, error)
	DeleteBook(ctx context.Context, bookId int) error
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, []int, error)
	UpdateBook(ctx context.Context, book model.Book, bookId int) (model.Book, error)

	GetBooksByCondition(ctx context.Context, num int, bookGenre string) ([]model.Book, []int, error)
}

type FavouritesRepository interface {
	CreateFavourite(ctx context.Context, login string, bookId int) (int, error)
	GetFavourites(ctx context.Context, login string) ([]model.Book, []int, error)
	DeleteFavourite(ctx context.Context, favouriteId int) error
}

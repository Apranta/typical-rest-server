package service

import (
	"github.com/typical-go/typical-rest-server/app/repository"
	"go.uber.org/dig"
)

// BookService contain logic for Book Controller
type BookService interface {
	repository.BookRepo
}

// BookServiceImpl is implementation of BookService
type BookServiceImpl struct {
	dig.In
	repository.BookRepo
}

// NewBookService return new instance of BookService
func NewBookService(impl BookServiceImpl) BookService {
	return &impl
}

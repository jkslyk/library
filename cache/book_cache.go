package cache

import "github.com/jkslyk/library/internal/domain"

type BookCache interface {
	Set(key string, value *domain.Book)
	Get(key string) *domain.Book
}

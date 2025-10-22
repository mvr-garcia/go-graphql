package graph

import "github.com/mvr-garcia/go-graphql/internal/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryRepo domain.CategoryRepository
	CourseRepo   domain.CourseRepository
}

func NewResolver(categoryRepo domain.CategoryRepository, courseRepo domain.CourseRepository) *Resolver {
	return &Resolver{
		CategoryRepo: categoryRepo,
		CourseRepo:   courseRepo,
	}
}

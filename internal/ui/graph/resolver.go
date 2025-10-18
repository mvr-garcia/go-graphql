package graph

import "github.com/mvr-garcia/go-graphql/internal/domain"

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

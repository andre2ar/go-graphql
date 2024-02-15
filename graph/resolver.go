package graph

import "github.com/andre2ar/go-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryRepository *database.Category
	CourseRepository   *database.Course
}

package domain

type CategoryRepository interface {
	FindAll() ([]Category, error)
	FindByID(id string) (Category, error)
	Create(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(id string) error
}

type CourseRepository interface {
	FindAll() ([]Course, error)
	FindByID(id string) (Course, error)
	Create(course Course) (Course, error)
	Update(course Course) (Course, error)
	Delete(id string) error
}

package infra

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mvr-garcia/go-graphql/internal/domain"
)

type CourseAdapter struct {
	db *sql.DB
}

func NewCourseAdapter(db *sql.DB) domain.CourseRepository {
	return &CourseAdapter{db: db}
}

func (ca *CourseAdapter) FindAll() ([]domain.Course, error) {
	rows, err := ca.db.Query("SELECT id, title, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		var course domain.Course
		if err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func (ca *CourseAdapter) FindByID(id string) (domain.Course, error) {
	var course domain.Course
	err := ca.db.QueryRow("SELECT id, title, description, category_id FROM courses WHERE id = ?", id).Scan(&course.ID, &course.Title, &course.Description, &course.CategoryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Course{}, domain.ErrNotFound
		}
		return domain.Course{}, err
	}
	return course, nil
}

func (ca *CourseAdapter) Create(course domain.Course) (domain.Course, error) {
	id := uuid.New().String()
	_, err := ca.db.Exec("INSERT INTO courses (id, title, description, category_id) VALUES (?, ?, ?, ?)", id, course.Title, course.Description, course.CategoryID)
	if err != nil {
		return domain.Course{}, err
	}
	return domain.Course{ID: id, Title: course.Title, Description: course.Description, CategoryID: course.CategoryID}, nil
}

func (ca *CourseAdapter) Update(course domain.Course) (domain.Course, error) {
	_, err := ca.db.Exec("UPDATE courses SET title = ?, description = ?, category_id = ? WHERE id = ?", course.Title, course.Description, course.CategoryID, course.ID)
	if err != nil {
		return domain.Course{}, err
	}
	return course, nil
}

func (ca *CourseAdapter) Delete(id string) error {
	_, err := ca.db.Exec("DELETE FROM courses WHERE id = ?", id)
	return err
}

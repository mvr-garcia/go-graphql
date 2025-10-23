package infra

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mvr-garcia/go-graphql/internal/domain"
)

type CategoryAdapter struct {
	db *sql.DB
}

func NewCategoryAdapter(db *sql.DB) domain.CategoryRepository {
	return &CategoryAdapter{db: db}
}

func (ca *CategoryAdapter) FindAll() ([]domain.Category, error) {
	rows, err := ca.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (ca *CategoryAdapter) FindByID(id string) (domain.Category, error) {
	var category domain.Category
	err := ca.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, domain.ErrNotFound
		}
		return domain.Category{}, err
	}
	return category, nil
}

func (ca *CategoryAdapter) FindByCourseID(courseID string) (domain.Category, error) {
	stmt := `
		SELECT c.id, c.name, c.description
		FROM categories c
		JOIN courses co ON co.category_id = c.id
		WHERE co.id = ?
	`

	var category domain.Category
	err := ca.db.QueryRow(stmt, courseID).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Category{}, domain.ErrNotFound
		}
		return domain.Category{}, err
	}

	return category, nil
}

func (ca *CategoryAdapter) Create(category domain.Category) (domain.Category, error) {
	id := uuid.New().String()
	_, err := ca.db.Exec("INSERT INTO categories (id, name, description) VALUES (?, ?, ?)", id, category.Name, category.Description)
	if err != nil {
		return domain.Category{}, err
	}
	return domain.Category{ID: id, Name: category.Name, Description: category.Description}, nil
}

func (ca *CategoryAdapter) Update(category domain.Category) (domain.Category, error) {
	_, err := ca.db.Exec("UPDATE categories SET name = ?, description = ? WHERE id = ?", category.Name, category.Description, category.ID)
	if err != nil {
		return domain.Category{}, err
	}
	return category, nil
}

func (ca *CategoryAdapter) Delete(id string) error {
	_, err := ca.db.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}

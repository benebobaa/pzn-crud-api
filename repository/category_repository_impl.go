package repository

import (
	"context"
	"database/sql"
	"errors"
	"pzn-crud-restapi/helper"
	"pzn-crud-restapi/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (reporepository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
func (reporepository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}
func (reporepository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "DELETE from category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

}
func (reporepository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"

	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		return category, nil
	} else {
		return category, errors.New("id not found")
	}
}
func (reporepository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

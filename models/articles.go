package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Joziaz/Go-DB/db"
)

type Article struct {
	ID          int
	Title       string
	Description string
}

func NewArticle() Article {
	return Article{}
}

type ArticleModel struct {
	db *sql.DB
}

// NewArticleModel return a ArticleModel
func NewArticleModel() ArticleModel {
	return ArticleModel{
		db: db.GetConnection(),
	}
}

// En una Api tendria que pasarle el contexto de la request

// GetAll Return all de rows of the model
func (model *ArticleModel) GetAll() []Article {

	var articles []Article

	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelfunc()

	sql := "Select * from Articles;"

	rows, err := model.db.QueryContext(ctx, sql)
	if err != nil {
		log.Printf("Error when execute select query %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var article Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Description,
		)
		if err != nil {
			log.Print(err)
		}

		articles = append(articles, article)
	}
	return articles
}

// GetOne return the article when the id match
func (model *ArticleModel) GetOne(id int) (Article, error) {
	var article Article

	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelfunc()

	sql := fmt.Sprintf("Select * from Articles where id = %d;", id)

	row := model.db.QueryRowContext(ctx, sql)

	err := row.Scan(&article.ID, &article.Title, &article.Description)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}

// Create insert a Article in the database
func (model *ArticleModel) Create(article Article) error {

	sql := `
		Insert into articles (title, description)
			Values ($1, $2)
	`
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelfunc()

	_, err := model.db.ExecContext(ctx, sql, article.Title, article.Description)
	if err != nil {
		return err
	}

	return nil
}

// Update recieve a the id an delete de article from
// the database
func (model *ArticleModel) Update(article Article) error {

	sql := `
		Update articles
			set title = $1, description = $2
			where id = $3;
	`
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelfunc()

	result, err := model.db.ExecContext(ctx, sql, article.Title, article.Description, article.ID)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		log.Fatalf("Expected singel row affected, got %d rows affected", rows)
	}
	return nil
}

// Delete recieve a the id an delete de article from
// the database
func (model *ArticleModel) Delete(id int) error {

	sql := `
		delete from articles 
			where id = $1;
	`
	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelfunc()

	result, err := model.db.ExecContext(ctx, sql, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		log.Fatalf("Expected singel row affected, got %d rows affected", rows)
	}

	return nil
}

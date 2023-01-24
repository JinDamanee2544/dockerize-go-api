package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {

	if db == nil {
		return c.JSON(http.StatusInternalServerError, "db is nil")
	}

	rows, err := db.Query("SELECT * FROM posts")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	posts := []Post{}
	for rows.Next() {
		p := Post{}
		err := rows.Scan(&p.ID, &p.Title)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		posts = append(posts, p)
	}

	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {
	p := Post{}
	err := c.Bind(&p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if db == nil {
		return c.JSON(http.StatusInternalServerError, "db is nil")
	}

	row := db.QueryRow("INSERT INTO posts (title) VALUES ($1) RETURNING id", p.Title)

	err = row.Scan(&p.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, p)
}

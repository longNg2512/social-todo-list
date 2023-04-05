package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"social-todo-list/todos"
)

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	// POST /v1/items (create new a item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// (PUT || PATCH) v1/items/:id (update item by id)
	// DELETE /v1/items/:id (delete item by id)

	v1 := r.Group("/v1")

	{
		items := v1.Group("/items")
		{
			items.POST("", todos.CreateItem(db))
			items.GET("/:id", todos.GetItem(db))
			items.GET("", todos.ListItem(db))
			items.PATCH("/:id", todos.UpdateItem(db))
			items.DELETE("/:id", todos.DeleteItem(db))
			items.GET("/paging", todos.Pagination(db))
		}
	}

	r.Run(":4060")
}

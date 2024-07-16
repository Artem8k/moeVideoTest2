package adCompanyRoutes

import (
	"context"
	"net/http"
	"strconv"
	"test/database/models"
	"test/database/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AddAdCompanyRoutes(rg *gin.RouterGroup, db *sqlx.DB) {
	adv := rg.Group("/adCompany")
	repo := repositories.NewAdCompanyRepository(db)

	// получить adCompany по id
	adv.GET("/get/:id", func(c *gin.Context) {

		// Создаю контекст
		ctx := context.Background()
		timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Достаю id из запроса
		idStr := c.Param("id")

		// Парсинг строки в число
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Вызываю метод репозитория
		resp, err := repo.Get(id, timeOutCtx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Отправляю ответ
		c.JSON(http.StatusOK, resp)
	})

	// Добавить adCompany
	adv.POST("/add", func(c *gin.Context) {
		// Создаю контекст
		ctx := context.Background()
		timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Достаю тело запроса
		var jsonReq models.AdCompany
		if err := c.ShouldBindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		jsonReq.Create_date = time.Now()

		// Вызываю метод репозитория
		err := repo.Add(jsonReq, timeOutCtx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Отправляю ответ
		c.JSON(http.StatusCreated, gin.H{"message": "Created"})
	})

	// Обновить adCompany
	adv.PUT("/update", func(c *gin.Context) {
		// Создаю контекст
		ctx := context.Background()
		timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Достаю тело запроса
		var jsonReq models.AdCompanyUpdate
		if err := c.ShouldBindJSON(&jsonReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Вызываю метод репозитория
		err := repo.Update(jsonReq, timeOutCtx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Отправляю ответ
		c.JSON(http.StatusCreated, gin.H{"message": "Updated"})
	})

	// Удалить adCompany
	adv.DELETE("/delete/:id", func(c *gin.Context) {
		// Создаю контекст
		ctx := context.Background()
		timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Достаю id из запроса
		idStr := c.Param("id")

		// Парсинг строки в число
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Вызываю метод репозитория
		err = repo.Delete(id, timeOutCtx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Отправляю ответ
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	})
}

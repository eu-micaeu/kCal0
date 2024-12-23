package handlers

import (
	"database/sql"
	"fmt"

	"time"

	"github.com/gin-gonic/gin"
)

type Food struct {

	Food_ID int `json:"food_id"`

	Meal_ID int `json:"meal_id"`

	FoodName string `json:"food_name"`

	Calories int `json:"calories"`

	Quantity int `json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
	
}

// Função para criar um novo alimento
func (f *Food) CriarAlimento(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		_, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var food Food

		if err := c.BindJSON(&food); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao criar alimento"})

			fmt.Println(err)

			return

		}

		_, err = db.Exec("INSERT INTO foods (meal_id, food_name, calories, quantity, created_at) VALUES ($1, $2, $3, $4, $5)", food.Meal_ID, food.FoodName, food.Calories, food.Quantity, time.Now())

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao criar alimento"})

			fmt.Println(err)

			return

		}

		c.JSON(200, gin.H{"message": "Alimento criado com sucesso"})

	}

}

// Função para listar todos os alimentos de uma refeição
func (f *Food) ListarAlimentosDeUmaRefeicao(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		_, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var foods []Food

		rows, err := db.Query("SELECT * FROM foods WHERE meal_id = $1", c.Param("meal_id"))

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao listar alimentos"})

			fmt.Println(err)

			return

		}

		for rows.Next() {

			var food Food

			rows.Scan(&food.Food_ID, &food.Meal_ID, &food.FoodName, &food.Calories, &food.Quantity, &food.CreatedAt)

			foods = append(foods, food)

		}

		c.JSON(200, gin.H{"foods": foods})

	}

}

// Função para listar um alimento
func (f *Food) ListarAlimento(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		_, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var food Food

		row := db.QueryRow("SELECT * FROM foods WHERE food_id = $1", c.Param("id"))

		err = row.Scan(&food.Food_ID, &food.Meal_ID, &food.FoodName, &food.Calories, &food.Quantity, &food.CreatedAt)

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao listar alimento"})

			fmt.Println(err)

			return

		}

		c.JSON(200, gin.H{"food": food})

	}

}

// Função para atualizar um alimento
func (f *Food) AtualizarAlimento(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		_, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		var food Food

		if err := c.BindJSON(&food); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao atualizar alimento"})

			fmt.Println(err)

			return

		}

		_, err = db.Exec("UPDATE foods SET food_name = $1, calories = $2, quantity = $3 WHERE food_id = $4", food.FoodName, food.Calories, food.Quantity, c.Param("food_id"))

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao atualizar alimento"})

			fmt.Println(err)

			return

		}

		c.JSON(200, gin.H{"message": "Alimento atualizado com sucesso"})

	}

}

// Função para deletar um alimento
func (f *Food) DeletarAlimento(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		_, err := ValidarOToken(token)

		if err != nil {

			c.JSON(401, gin.H{"message": "Token inválido"})

			return

		}

		_, err = db.Exec("DELETE FROM foods WHERE food_id = $1", c.Param("food_id"))

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao deletar alimento"})

			fmt.Println(err)

			return

		}

		c.JSON(200, gin.H{"message": "Alimento deletado com sucesso"})

	}

}
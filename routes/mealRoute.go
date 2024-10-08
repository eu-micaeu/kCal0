package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/eu-micaeu/kCal0/handlers"
)

// Função para criar as rotas de refeição
func MealRoutes(r *gin.Engine, db *sql.DB) {
	
	mealHandler := handlers.Meal{}

	r.POST("/createMeal", mealHandler.CriarRefeicao(db))

	r.GET("/listMenuMeals/:menu_id", mealHandler.ListarRefeicoesDeUmMenu(db))

	r.GET("/meal/:meal_id", mealHandler.CarregarRefeicao(db))

	r.GET("/calculateMealCaloriesAndQuantity/:meal_id", mealHandler.CalcularTotalDeCaloriasEQuantidadeDaRefeicao(db))

	r.DELETE("/deleteMeal/:meal_id", mealHandler.DeletarRefeicao(db))

}

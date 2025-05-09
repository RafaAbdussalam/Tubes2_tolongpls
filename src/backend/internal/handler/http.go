// internal/handler/http.go
package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"little_alchemy_backend/internal/model"
	"little_alchemy_backend/internal/repo"
	"little_alchemy_backend/internal/tree"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(repo *repo.RecipeRepository) *gin.Engine {

	// Create router
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/api/tree", func(context *gin.Context) {

		// Element query
		element := context.Query("element")
		if element == "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "element parameter is required"})
			return
		}

		// Mode query
		modeStr := context.DefaultQuery("mode", string(model.BFS))
		mode := model.Traversal(modeStr)

		// Amount query
		amountStr := context.DefaultQuery("amount", "1")
		amount, err := strconv.Atoi(amountStr)
		if err != nil || amount <= 0 {
			amount = 1
		}

		// Initialize builder
		builder, err := tree.NewBuilder(repo, mode)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Build tree
		recipeTree, err := builder.BuildTree(element, amount)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Set Recipe Count
		model.SetRecipeCount(recipeTree)

		// Send tree
		fmt.Println(recipeTree.String())
		context.JSON(http.StatusOK, recipeTree)

	})

	return router
}

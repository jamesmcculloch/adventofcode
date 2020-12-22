package main

import (
	"testing"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	foodLabels, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	foods, allergenToFoods := loadFoods(foodLabels)

	candidateIngredientsForAllergen := findCandidateIngredientsForAllergen(foods, allergenToFoods)

	ingredientAllergenMapping := make(map[string]string)
	findIngredientAllergenMapping(candidateIngredientsForAllergen, ingredientAllergenMapping)

	assert.Equal(t, 5, occurencesOfIngredientsWithNoAllergen(foods, ingredientAllergenMapping))
}

func TestPart2(t *testing.T) {
	foodLabels, err := utils.LoadStringsFromFile("input.sample")
	if err != nil {
		panic(err)
	}

	foods, allergenToFoods := loadFoods(foodLabels)

	candidateIngredientsForAllergen := findCandidateIngredientsForAllergen(foods, allergenToFoods)

	ingredientAllergenMapping := make(map[string]string)
	findIngredientAllergenMapping(candidateIngredientsForAllergen, ingredientAllergenMapping)

	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", findIngredientsSortedByAllergen(ingredientAllergenMapping))
}

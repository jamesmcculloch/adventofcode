package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
)

type food struct {
	allergens   map[string]bool
	ingredients map[string]bool
}

func loadFoods(foodLabels []string) ([]*food, map[string][]int) {
	foods := make([]*food, len(foodLabels))
	allergenToFoods := map[string][]int{}

	for index, foodLabel := range foodLabels {
		parts := strings.Split(foodLabel, "(")
		ingredients := strings.Split(strings.TrimSpace(parts[0]), " ")

		allergenDetails := strings.TrimPrefix(parts[1], "contains ")
		allergenDetails = strings.TrimSuffix(allergenDetails, ")")
		allergens := strings.Split(allergenDetails, ", ")

		newFood := &food{
			ingredients: make(map[string]bool),
			allergens:   make(map[string]bool),
		}

		for _, allergen := range allergens {
			newFood.allergens[allergen] = true
			allergenToFoods[allergen] = append(allergenToFoods[allergen], index)
		}

		for _, ingredient := range ingredients {
			newFood.ingredients[ingredient] = true
		}

		foods[index] = newFood
	}
	return foods, allergenToFoods
}

func findCandidateIngredientsForAllergen(foods []*food, allergenToFoods map[string][]int) map[string]map[string]bool {
	candidateIngredientsForAllergen := make(map[string]map[string]bool)

	for allergen, potentialFoods := range allergenToFoods {
		candidateIngredientsForAllergen[allergen] = make(map[string]bool)
		ingredientCount := make(map[string]int)
		for _, index := range potentialFoods {
			food := foods[index]
			for ingredient := range food.ingredients {
				ingredientCount[ingredient]++
			}
		}
		for ingredient, count := range ingredientCount {
			if count == len(potentialFoods) {
				candidateIngredientsForAllergen[allergen][ingredient] = true
			}
		}
	}

	return candidateIngredientsForAllergen
}

func findIngredientAllergenMapping(candidateIngredientsForAllergen map[string]map[string]bool, ingredientAllergenMapping map[string]string) {
	if len(candidateIngredientsForAllergen) == 0 {
		return
	}

	for allergen, ingredients := range candidateIngredientsForAllergen {
		for ingredient := range ingredientAllergenMapping {
			delete(ingredients, ingredient)
		}

		if len(ingredients) == 1 {
			for ingredient := range ingredients {
				ingredientAllergenMapping[ingredient] = allergen
			}

			delete(candidateIngredientsForAllergen, allergen)

			findIngredientAllergenMapping(candidateIngredientsForAllergen, ingredientAllergenMapping)
		}
	}
}

func occurencesOfIngredientsWithNoAllergen(foods []*food, ingredientAllergenMapping map[string]string) int {
	occurences := 0

	for _, food := range foods {
		for ingredient := range food.ingredients {
			if _, hasAllergen := ingredientAllergenMapping[ingredient]; !hasAllergen {
				occurences++
			}
		}
	}

	return occurences
}

func findIngredientsSortedByAllergen(ingredientAllergenMapping map[string]string) string {
	allergens := []string{}
	allegrenIngredientMapping := make(map[string]string)

	for ingredient, allergen := range ingredientAllergenMapping {
		allergens = append(allergens, allergen)
		allegrenIngredientMapping[allergen] = ingredient
	}

	sort.Strings(allergens)

	sortedIngredients := make([]string, len(allergens))

	for index, allergen := range allergens {
		sortedIngredients[index] = allegrenIngredientMapping[allergen]
	}

	return strings.Join(sortedIngredients, ",")
}

func main() {
	foodLabels, err := utils.LoadStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	foods, allergenToFoods := loadFoods(foodLabels)

	candidateIngredientsForAllergen := findCandidateIngredientsForAllergen(foods, allergenToFoods)

	ingredientAllergenMapping := make(map[string]string)
	findIngredientAllergenMapping(candidateIngredientsForAllergen, ingredientAllergenMapping)

	fmt.Printf("part 1: %d\n", occurencesOfIngredientsWithNoAllergen(foods, ingredientAllergenMapping))
	fmt.Printf("part 2: %s\n", findIngredientsSortedByAllergen(ingredientAllergenMapping))
}

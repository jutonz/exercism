package lasagna

// PreparationTime tells how long it will take to prepare the entire lasagna,
// if it takes timePerLayer minutes to prepare each layer
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers) * timePerLayer
}

const noodlesPerLayerInGrams = 50
const saucePerNoodleLayerInLiters = 0.2

// Quantities takes a slice of layers and returns how much noodles (in grams)
// and sauce (in liters) you'll need in order to complete the full recipe.
func Quantities(layers []string) (noodlesInGrams int, sauceInLiters float64) {
	noodlesInGrams = 0
	sauceInLiters = 0

	for _, layer := range layers {
		if layer == "sauce" {
			sauceInLiters += saucePerNoodleLayerInLiters
		} else if layer == "noodles" {
			noodlesInGrams += noodlesPerLayerInGrams
		}
	}

	return noodlesInGrams, sauceInLiters
}

// AddSecretIngredient takes two recipes, and replaces the final ingredient of
// the second recipe with the final ingredient of the first.
func AddSecretIngredient(layersWithSecretIngredient, myLayers []string) {
	secretIngredient := layersWithSecretIngredient[len(layersWithSecretIngredient) - 1]
	myLayers[len(myLayers) - 1] = secretIngredient
}

// ScaleRecipe takes the quantities needed to make two portions of lasagna, and
// returns what quantities are needed to produce the desired number of
// portions.
func ScaleRecipe(quantitiesForTwoPortions []float64, desiredPortions int) []float64 {
	desiredQuantities := make([]float64, len(quantitiesForTwoPortions))

	for i, quantitiyForTwo := range quantitiesForTwoPortions {
		quantityForOne := quantitiyForTwo / 2
		desiredQuantities[i] = quantityForOne * float64(desiredPortions)
	}

	return desiredQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

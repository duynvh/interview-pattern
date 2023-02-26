package main

import "fmt"

func findRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	// Use a dependents map for easier lookup
    dependents := make(map[string][]string)
    outdegrees := make(map[string]int)
    recipesMap := make(map[string]struct{})

    for i := 0; i < len(recipes); i++ {
        for _, ingredient := range ingredients[i] {
            dependents[ingredient] = append(dependents[ingredient], recipes[i])
            outdegrees[recipes[i]]++
        }

        recipesMap[recipes[i]] = struct{}{}
    }

    var queue []string
    var res []string
    var curr string
    queue = append(queue, supplies...)

    for len(queue) != 0 {
        curr, queue = queue[0], queue[1:]
        if _, ok := recipesMap[curr]; ok {
            res = append(res, curr)
        }

        for _, dependent := range dependents[curr] {
            outdegrees[dependent]--

            if outdegrees[dependent] == 0 {
                queue = append(queue, dependent)
            }
        }
    }

    return res
}

func main() {
	recipes := []string{"soup", "spaghetti"}
	ingredients := [][]string{
		{"meat", "yogurt", "sauce"}, 
        {"spaghetti", "pasta", "salt", "sauce"}, 
	}
    supplies := []string{"meat", "spaghetti", "pasta", "yogurt", "sauce"}

	// for i, prerequisite := range prerequisites {
	// 	fmt.Printf("%d.\tPrerequisites: %s\n", i+1, strings.Replace(fmt.Sprint(prerequisite), " ", ", ", -1))
	// 	fmt.Printf("\tTotal number of course, n: %d\n", n[i])
	// 	fmt.Printf("\tValid courses order: %s\n", strings.Replace(fmt.Sprint(findOrder(n[i], prerequisite)), " ", ", ", -1))
	// 	fmt.Printf("%s\n", strings.Repeat("-", 100))
	// }

    fmt.Println(findRecipes(recipes, ingredients, supplies))
}
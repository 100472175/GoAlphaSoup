package main

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

var directions = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func main() {
	// Check the number of arguments
	if len(os.Args) != 5 {
		fmt.Println("Usage: alphabet-soup <filename> <words> <multithread [Y|N]> <output [Y|N]>")
		return
	}

	if os.Args[3] != "Y" && os.Args[3] != "N" {
		fmt.Println("The third argument must be 1 or 0")
		return
	}

	if os.Args[4] != "Y" && os.Args[4] != "N" {
		fmt.Println("The fourth argument must be 1 or 0")
		return
	}

	time_start := time.Now()
	// Read the input file with the grid
	grid, err := ReadGrid(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tiempo de lectura de la sopa de letras:", time.Since(time_start))

	time_start = time.Now()
	// Read the input words to find
	words, err := ReadWords(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tiempo de lectura de las palabras:", time.Since(time_start))

	if os.Args[3] == "Y" {
		// Find the words in the grid and parallelize the search for each word
		time_start := time.Now()
		wg := sync.WaitGroup{}
		for _, word := range words {
			wg.Add(1)
			go func(word string, wg *sync.WaitGroup) {
				if os.Args[4] == "Y" {
					result := FindWord(grid, word)
					if len(result) > 0 {
						fmt.Println("La palabra", word, "se encuentra en las posiciones:", result)
					} else {
						fmt.Println("La palabra", word, "no se encuentra en la sopa de letras")
					}
				} else {
					FindWord(grid, word)
				}
				wg.Done()
			}(word, &wg)
		}
		wg.Wait()
		fmt.Fprint(os.Stderr, "Tiempo de ejecución: ", time.Since(time_start), "\n")

	} else {
		// Find the words in the grid and search for each word sequentially
		time_start := time.Now()
		for _, word := range words {
			if os.Args[4] == "Y" {
				result := FindWord(grid, word)
				if len(result) > 0 {
					fmt.Println("La palabra", word, "se encuentra en las posiciones:", result)
				} else {
					fmt.Println("La palabra", word, "no se encuentra en la sopa de letras")
				}
			} else {
				FindWord(grid, word)
			}
		}
		fmt.Fprint(os.Stderr, "Tiempo de ejecución: ", time.Since(time_start), "\n")

	}

}

func ContinuaPalabra(grid [][]rune, word string, i int, j int, result [][2]int, direction [2]int) ([][2]int, error) {
	if len(word) == 0 {
		return result, nil
	}
	res := result
	new_x := i
	new_y := j
	for _, letter := range word {
		new_x += direction[0]
		new_y += direction[1]
		if new_x >= 0 && new_x < len(grid) && new_y >= 0 && new_y < len(grid[0]) && grid[new_x][new_y] == letter {
			res = append(res, [2]int{new_x, new_y})
		} else {
			return res, errors.New("no se ha encontrado la palabra")
		}
	}

	if len(res) == 1 {
		res = res[:0]
	} else if len(res) == len(word)+2 {
		return res, nil
	}
	return res, errors.New("no se ha encontrado la palabra")
}

func FindWord(grid [][]rune, word string) [][2]int {
	var result_loc [][2]int
	for i, row := range grid {
		for j, cell := range row {
			if cell == rune(word[0]) {
				result_loc = append(result_loc, [2]int{i, j})
				for _, direction := range directions {
					new_x := i + direction[0]
					new_y := j + direction[1]
					if new_x >= 0 && new_x < len(grid) && new_y >= 0 && new_y < len(grid[0]) && grid[new_x][new_y] == rune(word[1]) {
						result_loc = append(result_loc, [2]int{new_x, new_y})
						result, _ := ContinuaPalabra(grid, word[2:], new_x, new_y, result_loc, [2]int{direction[0], direction[1]})
						if len(result) == len(word) {
							return result
						} else {
							result_loc = result_loc[:1]
						}
					}
				}
				result_loc = result_loc[:0]
			} else {
				continue
			}
		}
	}
	return result_loc
}

func ReadGrid(filename string) ([][]rune, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the grid
	var grid [][]rune
	for {
		var row []rune
		for {
			var cell rune
			_, err := fmt.Fscanf(file, "%c", &cell)
			if err != nil {
				break
			}
			if cell == '\n' {
				break
			} else {
				row = append(row, cell)
			}
		}
		if len(row) == 0 {
			break
		}
		grid = append(grid, row)
	}
	return grid, nil
}

func ReadWords(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the words
	var words []string
	for {
		var word string
		_, err := fmt.Fscanf(file, "%s", &word)
		if err != nil {
			break
		}
		words = append(words, word)
	}
	return words, nil
}

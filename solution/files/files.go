package files

import (
	"bufio"
	"os"
	"strconv"
)

func GetIntArray(filepath string) ([]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	array := []int{}
	for scanner.Scan() {
		newElementString := scanner.Text()
		newElement, err := strconv.Atoi(newElementString)
		if err != nil {
			return []int{}, err
		}
		array = append(array, newElement)
	}

	if err := scanner.Err(); err != nil {
		return []int{}, err
	}

	return array, nil
}

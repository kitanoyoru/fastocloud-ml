package utils

import (
	"github.com/sbinet/npyio/npz"
)

func LoadNPZFile(path string) (data [][]float32) {
	file, err := npz.Open(path)
	if err != nil {
		panic(err) // TODO: Handle error
	}
	defer file.Close()

	for _, key := range file.Keys() {
		arr := []float32{}
		err = file.Read(key, &arr)
		if err != nil {
			panic(err) // TODO: Handle error
		}
		data = append(data, arr)
	}

	return
}

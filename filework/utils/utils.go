package utils

import (
	"fmt"
	"os"
)

const counter int = 1000
const fileExtension string = ".json"

/*
// Return file
func getFile(name string, max int) *File {
	// Combine name with extension
	// If file not exist
	//// create file
	//// return file

	// If file exist
	//// count = 0, get precount
	//// Combine name with extension and precount
	//// If file not exist
	////// create file
	////// return file
	//// If file exist
}

*/

// CreateNewFile Создает новый файл
func CreateNewFile(dirJSON string, fileName string) *os.File {
	dir, err := os.Open(dirJSON)
	if err != nil {
		fmt.Println("Error when open directory" + dirJSON)
		os.Exit(1)
	}
	defer dir.Close()

	// Получение инфо файлов в каталоге
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error when read directory" + dirJSON)
		os.Exit(1)
	}

	// Проверка на наличие файла в каталоге
	var isAbsent bool = true
	for _, fi := range fileInfos {
		if fi.Name() == (fileName + fileExtension) {
			isAbsent = false
			break
		}
	}

	// Если отсутствует, то создаем
	if isAbsent {
		file, err := os.Create(dirJSON + fileName + fileExtension)
		if err != nil {
			fmt.Println("Error when creating file" + fileName + fileExtension)
			os.Exit(1)
		}
		return file
	}

	// Если присутствует, перебираем имена
	// и создаем с отсутствующим именем
	isAbsent = true
	for i := 0; i < counter; i++ {

		nameCombine := fileName + "_" + fmt.Sprintf("%.3d", i) + fileExtension

		for _, fi := range fileInfos {
			if fi.Name() == (nameCombine) {
				isAbsent = false
				break
			}
		}

		if isAbsent {
			file, err := os.Create(dirJSON + nameCombine)
			if err != nil {
				fmt.Println("Error when creating file" + nameCombine)
				os.Exit(1)
			}
			return file
		}
		isAbsent = true

	}
	return nil
}

//

//

//

// HasFileInDirectory Проверяет наличие файла в каталоге
func HasFileInDirectory(dirJSON string, fileName string) (bool, error) {
	dir, err := os.Open(dirJSON)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return false, err
	}

	for _, fi := range fileInfos {
		if fi.Name() == (fileName + fileExtension) {
			return true, nil
		}
	}

	return false, nil
}

//

//

// Tst Function
func Tst(dirJSON string) {
	dir, err := os.Open(dirJSON)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error in open directory %s, error is %s", dirJSON, err.Error()))
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error when read directory %s, error is %s", dirJSON, err.Error()))
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

}

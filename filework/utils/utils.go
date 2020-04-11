package utils

import (
	"fmt"
	"os"
)

// const counter int = 1000
// const fileExtension string = "json"

// GetNewNFile GetNFile
func GetNewNFile() NFile {
	return NFile{
		Name:      "newfile",
		Extension: "json",
		Directory: ".",
		MaxCount:  1000,
		OsFile:    nil,
	}
}

// NFile NFile
type NFile struct {
	Name      string
	Extension string
	Directory string
	MaxCount  int
	OsFile    *os.File
}

// FullName Возвращает имя файла с расширением
func (nf NFile) FullName() string {
	if nf.Extension == "" {
		return nf.Name
	} else {
		return nf.Name + "." + nf.Extension
	}
}

// Create Создает новый файл
func (nf NFile) Create() (*os.File, error) {
	dir, err := os.Open(nf.Directory)
	if err != nil {
		// if err == os.ErrNotExist {
		// 	dir, err = os.Create(nf.Directory)
		// } else {

		return nil, err
		// }
	}
	defer dir.Close()

	// Получение инфо файлов в каталоге
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
		// fmt.Println("Error when read directory" + nf.Directory)
		// os.Exit(1)
	}

	// Проверка на наличие файла в каталоге
	var isAbsent bool = true
	for _, fi := range fileInfos {
		if fi.Name() == (nf.FullName()) {
			isAbsent = false
			break
		}
	}

	// Если отсутствует, то создаем
	if isAbsent {

		file, err := os.Create(
			nf.Directory + nf.FullName())

		if err != nil {
			return nil, err
			// fmt.Println("Error when creating file" + nf.FullName())
			// os.Exit(1)
		}
		return file, nil
	}

	// Если присутствует, перебираем имена
	// и создаем с отсутствующим именем
	isAbsent = true
	for i := 0; i < nf.MaxCount; i++ {

		nameCombine := nf.Name + "_" + fmt.Sprintf("%.3d", i) + "." + nf.Extension

		for _, fi := range fileInfos {
			if fi.Name() == (nameCombine) {
				isAbsent = false
				break
			}
		}

		if isAbsent {
			file, err := os.Create(nf.Directory + nameCombine)
			if err != nil {
				return nil, err
				// fmt.Println("Error when creating file" + nameCombine)
				// os.Exit(1)
			}
			nf.Name = nf.Name + "_" + fmt.Sprintf("%.3d", i)
			return file, nil
		}
		isAbsent = true

	}
	return nil, nil
}

//

//

//

// IsFileInDirectory Проверяет наличие файла в каталоге
func IsFileInDirectory(dirName string, fileName string) (bool, error) {
	dir, err := os.Open(dirName)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return false, err
	}

	for _, fi := range fileInfos {
		if fi.Name() == (fileName) {
			return true, nil
		}
	}

	return false, nil
}

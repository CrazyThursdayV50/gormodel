package sql

import (
	"io/fs"
	"path/filepath"
)

// read the last dir name as .go package name
func ReadPackage(sqlFilePath string) string {
	dir := filepath.Dir(sqlFilePath)
	_, pkg := filepath.Split(dir)
	return pkg
}

func genWalkSqlFilesFunc(sqlFiles *[]string) func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == ".sql" {
			*sqlFiles = append(*sqlFiles, path)
		}

		return nil
	}
}

func ListAllSqlFiles(path string) []string {
	var sqlFiles []string
	fn := genWalkSqlFilesFunc(&sqlFiles)
	filepath.Walk(path, fn)
	return sqlFiles
}

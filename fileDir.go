package fileDir

import "os"
import "strings"
import "github.com/aurelienCastel/stringUtil"
import "github.com/aurelienCastel/errorUtil"

func CurrentDirName() string {
	var directoryName string
	var err error

	directoryName, err = os.Getwd()
	errorUtil.Check(err)

	return directoryName
}

func DirNamed(directoryName string) *os.File {
	var currentDir *os.File
	var err error

	currentDir, err = os.Open(directoryName)
	errorUtil.Check(err)
	err = currentDir.Close()
	errorUtil.Check(err)

	return currentDir
}

func CurrentDir() *os.File {
	return DirNamed(CurrentDirName())
}

func NameIsDir(fileName string) bool {
	fileInfo, err := os.Stat(fileName)
	errorUtil.Check(err)
	return fileInfo.IsDir()
}

func NamesInDir(directory *os.File) []string {
	var fileNames []string
	var err error

	fileNames, err = directory.Readdirnames(0)
	errorUtil.Check(err)

	return fileNames
}

func NamesInRecDir(directory *os.File) []string {
	var fileNames []string
	var file *os.File
	var err error

	for _, fileName := range NamesInDir(directory) {
		if NameIsDir(fileName) {
			file, err = os.Open(fileName)
			errorUtil.Check(err)
			err = file.Close()
			errorUtil.Check(err)
			fileNames = append(fileNames, NamesInRecDir(file)...)
		} else {
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames
}

func NamesWithExt(fileNames []string, extension string) []string {
	var validNames []string

	for _, fileName := range fileNames {
		if strings.HasSuffix(fileName, extension) {
			validNames = append(validNames, fileName)
		}
	}

	return validNames
}

func NamesWithExts(fileNames []string, extensions []string) []string {
	var validNames []string

	for _, fileName := range fileNames {
		if stringUtil.HasOneOfSuffixes(fileName, extensions) {
			validNames = append(validNames, fileName)
		}
	}

	return validNames
}

func NamesInDirWithExt(directory *os.File, extension string) []string {
	return NamesWithExt(NamesInDir(directory), extension)
}

func NamesInDirWithExts(directory *os.File, extensions []string) []string {
	return NamesWithExts(NamesInDir(directory), extensions)
}

func NamesInRecDirWithExt(directory *os.File, extension string) []string {
	return NamesWithExt(NamesInRecDir(directory), extension)
}

func NamesInRecDirWithExts(directory *os.File, extensions []string) []string {
	return NamesWithExts(NamesInRecDir(directory), extensions)
}

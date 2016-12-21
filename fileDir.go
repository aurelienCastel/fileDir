package fileDir

import "os"
import "strings"
import "github.com/aurelienCastel/stringUtil"
import "github.com/aurelienCastel/errorUtil"

func PrefixWithPath(fileName string, path string) string {
	return path + string(os.PathSeparator) + fileName
}

func PrefixEachWithPath(fileNames []string, path string) []string {
	var preffixeds []string

	for _, fileName := range fileNames {
		preffixeds = append(preffixeds, PrefixWithPath(fileName, path))
	}

	return preffixeds
}

func CurrentDirAbsoluteName() string {
	var directoryName string
	var err error

	directoryName, err = os.Getwd()
	errorUtil.Check(err)

	return directoryName
}

// Call close on the receiver of the returned *os.File once you finished with it
func DirNamed(directoryName string) *os.File {
	var currentDir *os.File
	var err error

	currentDir, err = os.Open(directoryName)
	errorUtil.Check(err)

	return currentDir
}

// Call close on the receiver of the returned *os.File once you finished with it
func CurrentDir() *os.File {
	return DirNamed(CurrentDirAbsoluteName())
}

func NameIsDir(fileName string) bool {
	var fileInfo os.FileInfo
	var err error

	fileInfo, err = os.Stat(fileName)
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

			fileNames = append(fileNames, NamesInRecDir(file)...)

			err = file.Close()
			errorUtil.Check(err)
		} else {
			fileNames = append(fileNames, fileName)
		}
	}

	return fileNames
}

func RelativeNamesInDir(directory *os.File) []string {
	return PrefixEachWithPath(NamesInDir(directory), directory.Name())
}

func RelativeNamesInRecDir(directory *os.File) []string {
	var fileNames []string
	var file *os.File
	var err error

	for _, fileName := range RelativeNamesInDir(directory) {
		if NameIsDir(fileName) {
			file, err = os.Open(fileName)
			errorUtil.Check(err)

			fileNames = append(fileNames, RelativeNamesInRecDir(file)...)

			err = file.Close()
			errorUtil.Check(err)
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

func RelativeNamesInDirWithExt(directory *os.File, extension string) []string {
	return NamesWithExt(RelativeNamesInDir(directory), extension)
}
func RelativeNamesInDirWithExts(directory *os.File, extensions []string) []string {
	return NamesWithExts(RelativeNamesInDir(directory), extensions)
}

func RelativeNamesInRecDirWithExt(directory *os.File, extension string) []string {
	return NamesWithExt(RelativeNamesInRecDir(directory), extension)
}
func RelativeNamesInRecDirWithExts(directory *os.File, extensions []string) []string {
	return NamesWithExts(RelativeNamesInRecDir(directory), extensions)
}

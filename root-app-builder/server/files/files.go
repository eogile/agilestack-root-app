package files

import (
	"log"
	"os"
	"path/filepath"
)

const (
	defaultSourcesDir       = "./"
	defaultHTTPDir          = "./build/"
	defaultSharedModulesDir = "./web_modules/"
)

var (
	SourcesDirectory       string
	OutputDirectory        string
	HTTPDirectory          string
	ReducersFile           string
	RoutesFile             string
	ComponentsFile         string
	SharedModulesDirectory string
)

func init() {
	SourcesDirectory = getFromEnvOrElse("SOURCES_DIRECTORY", defaultSourcesDir)
	HTTPDirectory = getFromEnvOrElse("HTTP_DIRECTORY", defaultHTTPDir)
	SharedModulesDirectory = getFromEnvOrElse("SHARED_MODULES_DIRECTORY", defaultSharedModulesDir)

	OutputDirectory = absolutePath(SourcesDirectory + "/build/")
	ReducersFile = absolutePath(SourcesDirectory + "/js/generated/reducers.js")
	RoutesFile = absolutePath(SourcesDirectory + "/js/generated/routes.js")
	ComponentsFile = absolutePath(SourcesDirectory + "/js/generated/components.js")

	log.Println("Sources directory:", SourcesDirectory)
	log.Println("Output directory :", OutputDirectory)
	log.Println("HTTP directory   :", HTTPDirectory)
	log.Println("Reducers files   :", ReducersFile)
	log.Println("Routes files     :", RoutesFile)
	log.Println("Components files :", ComponentsFile)
	log.Println("Modules directory:", SharedModulesDirectory)

	// Creating directories that may not exists
	createDirIfNotExist(HTTPDirectory)
	createDirIfNotExist(SharedModulesDirectory)
}

// Checking the validity of directories and files.
func CheckFilesExistence() {
	checkExists(SourcesDirectory, true)
	checkExists(HTTPDirectory, true)
	checkExists(SharedModulesDirectory, true)
	checkExists(ReducersFile, false)
	checkExists(RoutesFile, false)

}

func getFromEnvOrElse(envVariable string, defaultPath string) string {
	path := defaultPath
	if os.Getenv(envVariable) != "" {
		path = os.Getenv(envVariable)
	}
	return absolutePath(path)
}

func absolutePath(path string) string {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Invalid path \"%s\": %v", path, err)
	}
	return absolutePath
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func createDirIfNotExist(path string) {
	if exists(path) {
		return
	}
	log.Println("Creating directory:", path)
	err := os.MkdirAll(path, 0644)
	if err != nil {
		log.Fatalf("Error while creating directory %s: %v", path, err)
	}
}

func checkExists(path string, isDir bool) {
	if !exists(path) {
		log.Fatalln("The path does not exist: ", path)
	}
	stats, err := os.Stat(path)
	if err != nil {
		log.Fatalln("Error while checking path:", path, err)
	}
	if stats.IsDir() != isDir {
		log.Fatalf("Invalid path: is directory[expected:%b, actual:%b]\n",
			isDir, stats.IsDir())
	}
}

/*
 * This is the root application.
 *
 * Schemes: http, https
 * BasePath: /
 * Version: 0.0.1
 * Host:localhost:8080
 *
 * Consumes:
 * - application/json
 *
 * Produces:
 * - application/json
 *
 * swagger:meta
 */
package main

import (
	"log"
	"net/http"
	"os"

	"bytes"
	"io/ioutil"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/files"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/handlers"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/services"
	"github.com/eogile/agilestack-utils/plugins"
)

/*
 * Logs configuration.
 */
func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {

	initBaseURL()

	/*
	 * Checks the existence of the files and directories
	 * manipulated by the builder.
	 */
	files.CheckFilesExistence()

	err := services.BuildApplication()
	if err != nil {
		log.Fatalln("Unable to build the application: ", err)
	}

	http.HandleFunc("/status", plugins.HandleHttpStatusUrl)
	http.HandleFunc("/plugins", handlers.NewGenerationHandler(handlers.BuildApplication).HandlePluginsEndpoint)
	http.HandleFunc("/menu-entries", handlers.HandleMenuEntriesEndpoint)
	http.ListenAndServe(":8080", nil)
}
/*
 * modify the index.html to use the url prefix (basename) provided by environment variable
 */
func initBaseURL() {
	baseUrl := os.Getenv("HTTP_APP_BASE_URL")
	log.Println("Using HTTP_APP_BASE_URL = ", baseUrl)

	if baseUrl != "" {

		indexPath := files.SourcesDirectory + "/index.html"
		oldIndexContent, err := ioutil.ReadFile(indexPath)
		if err != nil {
			log.Fatalf("~/unable to find index.html :%v", err)
		}

		newIndexContent := bytes.Replace(oldIndexContent, []byte("window.baseUrl=\"/\""), []byte("window.baseUrl=\""+baseUrl+"\""), -1)
		err = ioutil.WriteFile(indexPath, newIndexContent, 0644)
	}
}

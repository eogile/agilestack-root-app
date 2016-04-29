package services

import (
	"log"

	appFiles "github.com/eogile/agilestack-root-app/root-app-builder/server/files"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/utils/npm"
	"github.com/eogile/agilestack-utils/files"
)

func BuildApplication() error {
	err := npm.LaunchWebpack()
	if err != nil {
		log.Println("Error while webpack compilation:", err)
		return err
	}

	if appFiles.HTTPDirectory != appFiles.OutputDirectory {
		from := appFiles.OutputDirectory
		to := appFiles.HTTPDirectory
		log.Printf("Copying content of %s into %s", from, to)
		err = files.CopyDir(from, to)
		if err != nil {
			log.Println("Error while copying files:", err)
		}
		return err
	} else {
		log.Println("Compilation artifacts not copied. " +
			"Output directory equals the HTTP directory.")
		return nil
	}

}

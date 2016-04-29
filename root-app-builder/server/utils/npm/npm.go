package npm

import (
	"log"
	"os"
	"os/exec"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/files"
)

/*
Links a NPM package via a symbolic link using the "npm link" command.
See : https://docs.npmjs.com/cli/link
*/
func LinkPackage(packageName string) error {
	packageDirectory := files.SharedModulesDirectory + "/" + packageName
	return runNPMCommand(files.SourcesDirectory, "link", packageDirectory)
}

/*
Compiles the JavaScript by launching a Webpack process in the sources directory.
*/
func LaunchWebpack() error {
	log.Printf("Compiling the %s directory with Webpack\n",
		files.SourcesDirectory)
	return runNPMCommand(files.SourcesDirectory, "run", "build")
}

/*
Runs a NPM command in the given directory for the given arguments.
*/
func runNPMCommand(directory string, commandArguments ...string) error {
	cmd := exec.Command("npm", commandArguments...)
	cmd.Dir = directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}

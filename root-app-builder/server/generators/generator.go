package generators

import (
	"bytes"
	"log"
	"os"
)

/*
 * Sets the given content into the given file.
 *
 * If the file already exists, then it's replaced by a new one.
 *
 * The file's permissions will be "666" so anyone will be able
 * to read its content.
 */
func generateFile(content string, fileName string) error {
	log.Println("Generating the file:", fileName)
	var buffer bytes.Buffer
	_, err := buffer.WriteString(content)
	if err != nil {
		log.Printf("Error while writing in the buffer: %v\n", err)
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Error while creating file %s : %v\n", fileName, err)
		return err
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	return err
}
/*
 * Models for Swagger
 */
package models

// swagger:response listOfMenuEntries
type MenuEntrySlice struct {

	//in:body
	Body []MenuEntry
}

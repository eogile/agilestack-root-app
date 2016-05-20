/*
 * Models for Swagger
 */
package models

import "github.com/eogile/agilestack-utils/plugins/menu"

// swagger:response listOfMenuEntries
type MenuEntrySlice struct {

	//in:body
	Body []menu.MenuEntry
}

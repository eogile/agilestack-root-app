package models

// swagger:model menuEntry
type MenuEntry struct {

	/*
	 * The technical ID.
	 */
	ID string `json:"id"`

	/*
	 * The name of the plugin that registered the menu entry.
	 *
	 * Required: true
	 */
	PluginName string `json:"pluginName"`

	/*
	 * The name of the menu entry.
	 *
	 * Unique: true
	 * Required: true
	 */
	Name string `json:"name"`

	/*
	 * The front-end route.
	 *
	 * Unique: true
	 * Required: true
	 */
	Route string `json:"route"`

	/*
	 * The menu entry weight (used to determine the rank of the
	 * menu entry).
	 *
	 * Required: true
	 * Minimum: 0
	 */
	Weight int `json:"weight"`
}

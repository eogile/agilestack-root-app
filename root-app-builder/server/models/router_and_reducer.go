package models

type (
	Route struct {
		Href          string `json:"href"`
		ComponentName string `json:"componentName"`
		ModuleName    string `json:"moduleName"`
	}

	Reducer struct {
		Name       string `json:"name"`
		ModuleName string `json:"moduleName"`
	}

	PluginConfiguration struct {
		ModuleName string    `json:"moduleName"`
		Reducers   []Reducer `json:"reducers"`
		Routes     []Route   `json:"routes"`
	}
)

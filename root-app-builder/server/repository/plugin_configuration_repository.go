package repository

import "github.com/eogile/agilestack-root-app/root-app-builder/server/models"

type PluginConfigurationRepository struct {
	configurations map[string]models.PluginConfiguration
}

func GetPluginConfigurationsRepository() *PluginConfigurationRepository {
	return &PluginConfigurationRepository{
		configurations: make(map[string]models.PluginConfiguration),
	}
}

func (repo *PluginConfigurationRepository) Add(config models.PluginConfiguration) {
	repo.configurations[config.ModuleName] = config
}

func (repo *PluginConfigurationRepository) GetRoutes() []models.Route {
	routes := make([]models.Route, 0)
	for _, config := range repo.configurations {
		routes = append(routes, config.Routes...)
	}
	return routes
}

func (repo *PluginConfigurationRepository) GetReducers() []models.Reducer {
	reducers := make([]models.Reducer, 0)
	for _, config := range repo.configurations {
		reducers = append(reducers, config.Reducers...)
	}
	return reducers
}
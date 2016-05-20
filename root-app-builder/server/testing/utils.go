package testing

import (
	"log"
	"testing"

	"github.com/eogile/agilestack-utils/dockerclient"
	"github.com/eogile/agilestack-utils/plugins/menu"
	"github.com/eogile/agilestack-utils/test"
	"os"
	"github.com/stretchr/testify/require"
	"github.com/hashicorp/consul/api"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func Main(m *testing.M) {
	menu.SetConsulAddress("127.0.0.1:8501")
	registration.SetConsulAddress("127.0.0.1:8501")

	/*
	 * Docker client for test utilities
	 */
	dockerClient := dockerclient.NewClient()

	/*
	 * Creating the Docker network if it does not exist.
	 */
	err := test.CreateNetworkIfNotExists(dockerClient, "testNetwork")
	if err != nil {
		log.Fatalln("Unable to create a Docker network:", err)
	}

	/*
	 * Starting a Consul Docker container.
	 */
	if err = test.StartConsulContainer(dockerClient); err != nil {
		log.Fatalln("Unable to start a Docker container", err)
	}

	exitCode := m.Run()

	/*
	 * Stopping the Consul Docker container
	 */
	test.RemoveConsulContainer(dockerClient)

	os.Exit(exitCode)
}

func DeleteAllStoreEntries(t *testing.T) {
	_, err := consulClient(t).KV().DeleteTree("agilestack/", &api.WriteOptions{})
	require.Nil(t, err)
}

func consulClient(t *testing.T) *api.Client {
	config := api.DefaultConfig()
	config.Address = "localhost:8501"
	client, err := api.NewClient(config)
	require.Nil(t, err)
	return client
}

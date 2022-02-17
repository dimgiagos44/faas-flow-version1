package openfaas

import (
	"log"

	"handler/config"
	"handler/function"

	consulStateStore "github.com/faasflow/faas-flow-consul-statestore"
	redisStateStore "github.com/faasflow/faas-flow-redis-statestore"
	"github.com/faasflow/sdk"
)

func initStateStore() (stateStore sdk.StateStore, err error) {
	//stateStore, err = function.OverrideStateStore()
	stateStore, err = redisStateStore.GetRedisStateStore(os.Getenv("redis_url"), os.Getenv("redis_master"))
	if err != nil {
		return nil, err
	}

	if stateStore == nil {
		log.Print("Using default state store (consul)")

		consulURL := config.ConsulURL()
		consulDC := config.ConsulDC()

		stateStore, err = consulStateStore.GetConsulStateStore(consulURL, consulDC)
	}

	return stateStore, err
}

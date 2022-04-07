package external

import "github.com/google/wire"

var ExternalSet = wire.NewSet(NewCache, NewDB, NewMQTTClient)

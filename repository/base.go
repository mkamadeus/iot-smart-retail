package repository

import "github.com/google/wire"

var RepositorySet = wire.NewSet(NewCache, NewDB)

package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewCacheService,
	NewEmployeeService,
	NewLoginService,
	NewLeaveService,
)

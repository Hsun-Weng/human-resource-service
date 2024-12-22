package controllers

import (
	v1 "github.com/Hsun-Weng/human-resource-service/internal/controllers/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	v1.NewEmployeeController,
	v1.NewLoginController,
	v1.NewLeaveController,
)

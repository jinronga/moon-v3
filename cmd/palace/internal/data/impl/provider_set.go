package impl

import (
	"github.com/google/wire"
)

// ProviderSetImpl is a set of providers.
var ProviderSetImpl = wire.NewSet(
	NewUserRepo,
	NewCaptchaRepo,
	NewCacheRepo,
)

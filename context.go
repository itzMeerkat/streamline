package streamline

import "gitee.com/fat_marmota/infra/log"

type Context struct {
	Logger log.Logger
	// These are for RBAC authentication
	// If any of them is nil, it means no authentication is enabled
	Roles []string
	Action string
	Resource string
}
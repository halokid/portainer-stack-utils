package cmd

import (
	"github.com/greenled/portainer-stack-utils/client"
	"github.com/greenled/portainer-stack-utils/common"
)

func init() {
	common.AccessCmdInitFunc(configCmd, client.ResourceConfig)
}

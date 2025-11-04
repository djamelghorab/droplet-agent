// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/digitalocean/droplet-agent/internal/config"
	"github.com/digitalocean/droplet-agent/internal/log"
	"github.com/digitalocean/droplet-agent/internal/metadata/watcher"
)

func newMetadataWatcher(agentCfg *config.Conf, cfg *watcher.Conf) watcher.MetadataWatcher {
	if agentCfg != nil && agentCfg.UseWebWatcher {
		log.Info("Launching Web-based Watcher")
		return watcher.NewWebBasedWatcher(cfg)
	}

	log.Info("Launching SSH Port Knocking Watcher")
	return watcher.NewSSHWatcher(cfg)
}

// SPDX-License-Identifier: Apache-2.0

//go:build !linux
// +build !linux

package main

import (
	"github.com/digitalocean/droplet-agent/internal/config"
	"github.com/digitalocean/droplet-agent/internal/log"
	"github.com/digitalocean/droplet-agent/internal/metadata/watcher"
)

func newMetadataWatcher(_ *config.Conf, cfg *watcher.Conf) watcher.MetadataWatcher {
	log.Info("Launching Web-based Watcher")
	return watcher.NewWebBasedWatcher(cfg)
}

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliep.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"github.com/lsutils/kind/pkg/internal/apis/config"
	"github.com/lsutils/kind/pkg/internal/sets"
)

// RequiredNodeImages returns the set of _node_ images specified by the config
// This does not include the loadbalancer image, and is only used to improve
// the UX by explicit pulling the node images prior to running
func RequiredNodeImages(cfg *config.Cluster) sets.String {
	images := sets.NewString()
	for _, node := range cfg.Nodes {
		images.Insert(node.Image)
	}
	return images
}

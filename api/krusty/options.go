// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package krusty

import (
	"sigs.k8s.io/kustomize/api/konfig"
	"sigs.k8s.io/kustomize/api/types"
)

// Options holds high-level kustomize configuration options,
// e.g. are plugins enabled, should the loader be restricted
// to the kustomization root, etc.
type Options struct {
	// When true, sort the resources before emitting them,
	// per a particular sort order.  When false, don't do the
	// sort, and instead respect the depth-first resource input
	// order as specified by the kustomization file(s).
	DoLegacyResourceSort bool

	// When true, a label
	//     app.kubernetes.io/managed-by: kustomize-<version>
	// is added to all the resources in the build out.
	AddManagedbyLabel bool

	// Restrictions on what can be loaded from the file system.
	// See type definition.
	LoadRestrictions types.LoadRestrictions

	// Create an inventory object for pruning.
	DoPrune bool

	// Options related to kustomize plugins.
	PluginConfig *types.PluginConfig

	// When true, use kyaml/ packages to manipulate KRM yaml.
	// When false, use k8sdeps/ instead (uses k8s.io/api* packages).
	UseKyaml bool

	// When true, allow name and kind changing via a patch
	// When false, patch name/kind don't overwrite target name/kind
	AllowResourceIdChanges bool
}

// MakeDefaultOptions returns a default instance of Options.
// TODO(#3343): UseKyaml: konfig.FlagEnableKyamlDefaultValue
func MakeDefaultOptions() *Options {
	return &Options{
		DoLegacyResourceSort:   false,
		AddManagedbyLabel:      false,
		LoadRestrictions:       types.LoadRestrictionsRootOnly,
		DoPrune:                false,
		PluginConfig:           konfig.DisabledPluginConfig(),
		UseKyaml:               false,
		AllowResourceIdChanges: false,
	}
}

func (o Options) IfApiMachineryElseKyaml(s1, s2 string) string {
	if !o.UseKyaml {
		return s1
	}
	return s2
}

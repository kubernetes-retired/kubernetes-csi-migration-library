/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package translate

import (
	"fmt"

	"github.com/kubernetes-csi/kubernetes-csi-migration-library/plugins"
	"k8s.io/api/core/v1"
)

var (
	inTreePlugins = map[string]plugins.InTreePlugin{
		plugins.GCEPDDriverName: &plugins.GCEPD{},
	}
)

// TranslateToCSI takes a volume.Spec and will translate it to a
// CSIPersistentVolumeSource if the translation logic for that
// specific in-tree volume spec has been implemented
func TranslateInTreePVToCSI(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	// TODO: probably shouldn't need to do this explicity copy if we had good mutation semantics...
	// The semantics should be that the original PV does NOT get mutated...
	if pv == nil {
		return nil, fmt.Errorf("persistent volume was nil")
	}
	for _, curPlugin := range inTreePlugins {
		if curPlugin.CanSupport(pv) {
			return curPlugin.TranslateInTreePVToCSI(pv)
		}
	}
	return nil, fmt.Errorf("could not find in-tree plugin translation logic for %#v", pv.Name)
}

// TranslateToIntree takes a CSIPersistentVolumeSource and will translate
// it to a volume.Spec for the specific in-tree volume specified by
//`inTreePlugin`, if that translation logic has been implemented
func TranslateCSIPVToInTree(pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	if pv == nil || pv.Spec.CSI == nil {
		return nil, fmt.Errorf("CSI persistent volume was nil")
	}
	for driverName, curPlugin := range inTreePlugins {
		if pv.Spec.CSI.Driver == driverName {
			return curPlugin.TranslateCSIPVToInTree(pv)
		}
	}
	return nil, fmt.Errorf("could not find in-tree plugin translation logic for %s", pv.Spec.CSI.Driver)
}

func IsMigratedByName(pluginName string) bool {
	for _, curPlugin := range inTreePlugins {
		if curPlugin.GetInTreePluginName() == pluginName {
			return true
		}
	}
	return false
}

func IsPVMigrated(pv *v1.PersistentVolume) bool {
	for _, curPlugin := range inTreePlugins {
		if curPlugin.CanSupport(pv) {
			return true
		}
	}
	return false
}

func IsInlineMigrated(vol *v1.Volume) bool {
	return false
}

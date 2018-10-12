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

package plugins

import (
	"k8s.io/api/core/v1"
)

type InTreePlugin interface {
	// TranslateToCSI takes a volume.Spec and will translate it to a
	// CSIPersistentVolumeSource if the translation logic for that
	// specific in-tree volume spec has been implemented
	TranslateInTreePVToCSI(pv *v1.PersistentVolume) (*v1.PersistentVolume, error)

	// TranslateToIntree takes a CSIPersistentVolumeSource and will translate
	// it to a volume.Spec for the specific in-tree volume specified by
	//`inTreePlugin`, if that translation logic has been implemented
	TranslateCSIPVToInTree(pv *v1.PersistentVolume) (*v1.PersistentVolume, error)

	// CanSupport tests whether the plugin supports a given volume
	// specification from the API.  The spec pointer should be considered
	// const.
	CanSupport(pv *v1.PersistentVolume) bool

	// GetInTreePluginName returns the in-tree plugin name this migrates
	GetInTreePluginName() string
}

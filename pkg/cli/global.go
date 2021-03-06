// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/spf13/viper"
)

const (
	AcceptedLicenseVersions = "license.acceptedVersions"
)

type PersistentGlobalConfig interface {
	LicenseAcceptedForVersion(string) bool
	SetLicenseAcceptedForVersion(string)
	PersistConfig() error
}

type viperPersistentGlobalConfig struct {
	viper *viper.Viper
}

func newViperPersistentGlobalConfig(persistentGlobalConfig *viper.Viper) PersistentGlobalConfig {
	config := &viperPersistentGlobalConfig{
		viper: persistentGlobalConfig,
	}
	return config
}

func (v *viperPersistentGlobalConfig) LicenseAcceptedForVersion(version string) bool {
	for _, acceptedVersion := range v.viper.GetStringSlice(AcceptedLicenseVersions) {
		if acceptedVersion == version {
			return true
		}
	}
	return false
}

func (v *viperPersistentGlobalConfig) SetLicenseAcceptedForVersion(version string) {
	acceptedVersions := v.viper.GetStringSlice(AcceptedLicenseVersions)
	included := false
	for _, acceptedVersion := range acceptedVersions {
		if acceptedVersion == version {
			included = true
		}
	}
	if !included {
		acceptedVersions = append(acceptedVersions, version)
		v.viper.Set(AcceptedLicenseVersions, acceptedVersions)
	}
}

func (v *viperPersistentGlobalConfig) PersistConfig() error {
	return persistViper(v.viper)
}

/*
Copyright 2020 Red Hat OpenShift Container Storage.

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

package v1alpha1

import "k8s.io/apimachinery/pkg/types"

const ApiServerExportedAddressAnnotationName = "ocs.openshift.io/api-server-exported-address"

// ConnectedClient describes the connected clients of the storage cluster key
type ConnectedClient struct {
	Name      string `yaml:"name"`
	ClusterID string `yaml:"clusterId"`
	ClientID  string `yaml:"clientId"`
}

type InfoCephBlockPool struct {
	Name          string `yaml:"name"`
	MirrorEnabled bool   `yaml:"mirrorEnabled"`
}

// InfoStorageCluster describes information regarding a storage cluster key
type InfoStorageCluster struct {
	NamespacedName          types.NamespacedName `yaml:"namespacedName"`
	StorageProviderEndpoint string               `yaml:"storageProviderEndpoint"`
	CephClusterFSID         string               `yaml:"cephClusterFSID"`
	StorageClusterUID       string               `yaml:"storageClusterUID"`
	Annotations             map[string]string    `yaml:"annotations"`
	InfoCephBlockPools      []InfoCephBlockPool  `yaml:"infoCephBlockPools"`
}

// OdfInfoData describes odf-info CM's data
type OdfInfoData struct {
	Version        string             `yaml:"version"`
	DeploymentType string             `yaml:"deploymentType"`
	Clients        []ConnectedClient  `yaml:"clients"`
	StorageCluster InfoStorageCluster `yaml:"storageCluster"`
}

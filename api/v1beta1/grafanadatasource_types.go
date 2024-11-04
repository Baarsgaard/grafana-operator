/*
Copyright 2022.

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

package v1beta1

import (
	"encoding/json"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GrafanaDatasourceInternal struct {
	UID           string `json:"uid,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
	URL           string `json:"url,omitempty"`
	Access        string `json:"access,omitempty"`
	Database      string `json:"database,omitempty"`
	User          string `json:"user,omitempty"`
	IsDefault     *bool  `json:"isDefault,omitempty"`
	BasicAuth     *bool  `json:"basicAuth,omitempty"`
	BasicAuthUser string `json:"basicAuthUser,omitempty"`

	// Deprecated field, it has no effect
	OrgID *int64 `json:"orgId,omitempty"`
	// Deprecated field, it has no effect
	Editable *bool `json:"editable,omitempty"`

	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
	// +optional
	JSONData json.RawMessage `json:"jsonData,omitempty"`

	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
	// +optional
	SecureJSONData json.RawMessage `json:"secureJsonData,omitempty"`
}

// GrafanaDatasourceSpec defines the desired state of GrafanaDatasource
type GrafanaDatasourceSpec struct {
	GrafanaUIDSpec    `json:",inline"`
	GrafanaCommonSpec `json:",inline"`
	Datasource        *GrafanaDatasourceInternal `json:"datasource"`

	// plugins
	// +optional
	Plugins PluginList `json:"plugins,omitempty"`

	// environments variables from secrets or config maps
	// +optional
	// +kubebuilder:validation:MaxItems=99
	ValuesFrom []ValueFrom `json:"valuesFrom,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GrafanaDatasource is the Schema for the grafanadatasources API
// +kubebuilder:printcolumn:name="No matching instances",type="boolean",JSONPath=".status.NoMatchingInstances",description=""
// +kubebuilder:printcolumn:name="Last resync",type="date",format="date-time",JSONPath=".status.lastResync",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""
// +kubebuilder:resource:categories={grafana-operator}
type GrafanaDatasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaDatasourceSpec `json:"spec,omitempty"`
	Status GrafanaCommonStatus   `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GrafanaDatasourceList contains a list of GrafanaDatasource
type GrafanaDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrafanaDatasource `json:"items"`
}

func (in *GrafanaDatasource) Unchanged(hash string) bool {
	return in.Status.Hash == hash
}

func (in *GrafanaDatasource) IsAllowCrossNamespaceImport() bool {
	if in.Spec.AllowCrossNamespaceImport != nil {
		return *in.Spec.AllowCrossNamespaceImport
	}
	return false
}

func (in *GrafanaDatasourceList) Find(namespace string, name string) *GrafanaDatasource {
	for _, datasource := range in.Items {
		if datasource.Namespace == namespace && datasource.Name == name {
			return &datasource
		}
	}
	return nil
}

func init() {
	SchemeBuilder.Register(&GrafanaDatasource{}, &GrafanaDatasourceList{})
}

func (in *GrafanaDatasource) ResyncPeriodHasElapsed() bool {
	deadline := in.Status.LastResync.Add(in.Spec.ResyncPeriod.Duration)
	return time.Now().After(deadline)
}

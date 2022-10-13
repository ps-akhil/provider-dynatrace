/*
Copyright 2020 The Crossplane Authors.

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

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// OnBoardingParameters are the configurable fields of a OnBoarding.
type OnBoardingParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// OnBoardingObservation are the observable fields of a OnBoarding.
type OnBoardingObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A OnBoardingSpec defines the desired state of a OnBoarding.
type OnBoardingSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       OnBoardingParameters `json:"forProvider"`
}

// A OnBoardingStatus represents the observed state of a OnBoarding.
type OnBoardingStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          OnBoardingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A OnBoarding is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,dynatrace}
type OnBoarding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OnBoardingSpec   `json:"spec"`
	Status OnBoardingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OnBoardingList contains a list of OnBoarding
type OnBoardingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OnBoarding `json:"items"`
}

// OnBoarding type metadata.
var (
	OnBoardingKind             = reflect.TypeOf(OnBoarding{}).Name()
	OnBoardingGroupKind        = schema.GroupKind{Group: Group, Kind: OnBoardingKind}.String()
	OnBoardingKindAPIVersion   = OnBoardingKind + "." + SchemeGroupVersion.String()
	OnBoardingGroupVersionKind = SchemeGroupVersion.WithKind(OnBoardingKind)
)

func init() {
	SchemeBuilder.Register(&OnBoarding{}, &OnBoardingList{})
}

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CassandraDatacenterObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CassandraDatacenterParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZonesEnabled *bool `json:"availabilityZonesEnabled,omitempty" tf:"availability_zones_enabled,omitempty"`

	// +kubebuilder:validation:Required
	CassandraClusterID *string `json:"cassandraClusterId" tf:"cassandra_cluster_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetID *string `json:"delegatedManagementSubnetId,omitempty" tf:"delegated_management_subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetIDRef *v1.Reference `json:"delegatedManagementSubnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetIDSelector *v1.Selector `json:"delegatedManagementSubnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
}

// CassandraDatacenterSpec defines the desired state of CassandraDatacenter
type CassandraDatacenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraDatacenterParameters `json:"forProvider"`
}

// CassandraDatacenterStatus defines the observed state of CassandraDatacenter.
type CassandraDatacenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraDatacenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraDatacenter is the Schema for the CassandraDatacenters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CassandraDatacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CassandraDatacenterSpec   `json:"spec"`
	Status            CassandraDatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraDatacenterList contains a list of CassandraDatacenters
type CassandraDatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraDatacenter `json:"items"`
}

// Repository type metadata.
var (
	CassandraDatacenter_Kind             = "CassandraDatacenter"
	CassandraDatacenter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraDatacenter_Kind}.String()
	CassandraDatacenter_KindAPIVersion   = CassandraDatacenter_Kind + "." + CRDGroupVersion.String()
	CassandraDatacenter_GroupVersionKind = CRDGroupVersion.WithKind(CassandraDatacenter_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraDatacenter{}, &CassandraDatacenterList{})
}
/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NodePoolsObservation struct {

	// Enable the auto scaler for the default node pool.
	AutoScaler *bool `json:"autoScaler,omitempty" tf:"auto_scaler,omitempty"`

	// The VKE cluster ID you want to attach this nodepool to.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Date of node pool creation.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// Date of node pool updates.
	DateUpdated *string `json:"dateUpdated,omitempty" tf:"date_updated,omitempty"`

	// The Nodepool ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The label to be used as a prefix for nodes in this node pool.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The maximum number of nodes to use with the auto scaler.
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// The minimum number of nodes to use with the auto scaler.
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The number of nodes in this node pool.
	NodeQuantity *float64 `json:"nodeQuantity,omitempty" tf:"node_quantity,omitempty"`

	// Array that contains information about nodes within this node pool.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// The plan to be used in this node pool. See Plans List Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// Status of node pool.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A tag that is assigned to this node pool.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NodePoolsParameters struct {

	// Enable the auto scaler for the default node pool.
	// +kubebuilder:validation:Optional
	AutoScaler *bool `json:"autoScaler,omitempty" tf:"auto_scaler,omitempty"`

	// The VKE cluster ID you want to attach this nodepool to.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// The label to be used as a prefix for nodes in this node pool.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The maximum number of nodes to use with the auto scaler.
	// +kubebuilder:validation:Optional
	MaxNodes *float64 `json:"maxNodes,omitempty" tf:"max_nodes,omitempty"`

	// The minimum number of nodes to use with the auto scaler.
	// +kubebuilder:validation:Optional
	MinNodes *float64 `json:"minNodes,omitempty" tf:"min_nodes,omitempty"`

	// The number of nodes in this node pool.
	// +kubebuilder:validation:Optional
	NodeQuantity *float64 `json:"nodeQuantity,omitempty" tf:"node_quantity,omitempty"`

	// The plan to be used in this node pool. See Plans List Note the minimum plan requirements must have at least 1 core and 2 gbs of memory.
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// A tag that is assigned to this node pool.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NodesObservation struct {

	// Date of node pool creation.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// The Nodepool ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The label to be used as a prefix for nodes in this node pool.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Status of node pool.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type NodesParameters struct {
}

// NodePoolsSpec defines the desired state of NodePools
type NodePoolsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolsParameters `json:"forProvider"`
}

// NodePoolsStatus defines the observed state of NodePools.
type NodePoolsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePools is the Schema for the NodePoolss API. Provides a Vultr Kubernetes Engine (VKE) Node Pool resource. This can be used to create, read, modify, and delete VKE clusters on your Vultr account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vultr}
type NodePools struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.clusterId)",message="clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.label)",message="label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.nodeQuantity)",message="nodeQuantity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.plan)",message="plan is a required parameter"
	Spec   NodePoolsSpec   `json:"spec"`
	Status NodePoolsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolsList contains a list of NodePoolss
type NodePoolsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePools `json:"items"`
}

// Repository type metadata.
var (
	NodePools_Kind             = "NodePools"
	NodePools_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePools_Kind}.String()
	NodePools_KindAPIVersion   = NodePools_Kind + "." + CRDGroupVersion.String()
	NodePools_GroupVersionKind = CRDGroupVersion.WithKind(NodePools_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePools{}, &NodePoolsList{})
}

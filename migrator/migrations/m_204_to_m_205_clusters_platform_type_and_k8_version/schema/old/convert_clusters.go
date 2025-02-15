// Code generated by pg-bindings generator. DO NOT EDIT.
package schema

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertClusterFromProto converts a `*storage.Cluster` to Gorm model
func ConvertClusterFromProto(obj *storage.Cluster) (*Clusters, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &Clusters{
		ID:                                obj.GetId(),
		Name:                              obj.GetName(),
		Labels:                            obj.GetLabels(),
		StatusProviderMetadataClusterType: obj.GetStatus().GetProviderMetadata().GetCluster().GetType(),
		Serialized:                        serialized,
	}
	return model, nil
}

// ConvertClusterToProto converts Gorm model `Clusters` to its protobuf type object
func ConvertClusterToProto(m *Clusters) (*storage.Cluster, error) {
	var msg storage.Cluster
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}

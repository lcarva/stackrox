// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *DiscoveredCluster) Size() int                   { return m.SizeVT() }
func (m *DiscoveredCluster) Clone() *DiscoveredCluster   { return m.CloneVT() }
func (m *DiscoveredCluster) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *DiscoveredCluster) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *DiscoveredCluster_Metadata) Size() int                          { return m.SizeVT() }
func (m *DiscoveredCluster_Metadata) Clone() *DiscoveredCluster_Metadata { return m.CloneVT() }
func (m *DiscoveredCluster_Metadata) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *DiscoveredCluster_Metadata) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

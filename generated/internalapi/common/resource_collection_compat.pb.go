// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package common

func (m *ResourceCollection) Size() int                   { return m.SizeVT() }
func (m *ResourceCollection) Clone() *ResourceCollection  { return m.CloneVT() }
func (m *ResourceCollection) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ResourceCollection) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ResourceCollection_EmbeddedResourceCollection) Size() int { return m.SizeVT() }
func (m *ResourceCollection_EmbeddedResourceCollection) Clone() *ResourceCollection_EmbeddedResourceCollection {
	return m.CloneVT()
}
func (m *ResourceCollection_EmbeddedResourceCollection) Marshal() ([]byte, error) {
	return m.MarshalVT()
}
func (m *ResourceCollection_EmbeddedResourceCollection) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ResourceSelector) Size() int                   { return m.SizeVT() }
func (m *ResourceSelector) Clone() *ResourceSelector    { return m.CloneVT() }
func (m *ResourceSelector) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ResourceSelector) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *SelectorRule) Size() int                   { return m.SizeVT() }
func (m *SelectorRule) Clone() *SelectorRule        { return m.CloneVT() }
func (m *SelectorRule) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SelectorRule) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *RuleValue) Size() int                   { return m.SizeVT() }
func (m *RuleValue) Clone() *RuleValue           { return m.CloneVT() }
func (m *RuleValue) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *RuleValue) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

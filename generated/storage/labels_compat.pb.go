// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *LabelSelector) Size() int                   { return m.SizeVT() }
func (m *LabelSelector) Clone() *LabelSelector       { return m.CloneVT() }
func (m *LabelSelector) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *LabelSelector) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *LabelSelector_Requirement) Size() int                         { return m.SizeVT() }
func (m *LabelSelector_Requirement) Clone() *LabelSelector_Requirement { return m.CloneVT() }
func (m *LabelSelector_Requirement) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *LabelSelector_Requirement) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *SetBasedLabelSelector) Size() int                     { return m.SizeVT() }
func (m *SetBasedLabelSelector) Clone() *SetBasedLabelSelector { return m.CloneVT() }
func (m *SetBasedLabelSelector) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *SetBasedLabelSelector) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *SetBasedLabelSelector_Requirement) Size() int { return m.SizeVT() }
func (m *SetBasedLabelSelector_Requirement) Clone() *SetBasedLabelSelector_Requirement {
	return m.CloneVT()
}
func (m *SetBasedLabelSelector_Requirement) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SetBasedLabelSelector_Requirement) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

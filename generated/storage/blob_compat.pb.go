// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *Blob) Size() int                   { return m.SizeVT() }
func (m *Blob) Clone() *Blob                { return m.CloneVT() }
func (m *Blob) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *Blob) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

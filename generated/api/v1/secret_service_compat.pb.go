// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v1

func (m *SecretList) Size() int                   { return m.SizeVT() }
func (m *SecretList) Clone() *SecretList          { return m.CloneVT() }
func (m *SecretList) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *SecretList) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ListSecretsResponse) Size() int                   { return m.SizeVT() }
func (m *ListSecretsResponse) Clone() *ListSecretsResponse { return m.CloneVT() }
func (m *ListSecretsResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ListSecretsResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *CountSecretsResponse) Size() int                    { return m.SizeVT() }
func (m *CountSecretsResponse) Clone() *CountSecretsResponse { return m.CloneVT() }
func (m *CountSecretsResponse) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *CountSecretsResponse) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

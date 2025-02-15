// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package storage

func (m *SignatureIntegration) Size() int                    { return m.SizeVT() }
func (m *SignatureIntegration) Clone() *SignatureIntegration { return m.CloneVT() }
func (m *SignatureIntegration) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *SignatureIntegration) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

func (m *CosignPublicKeyVerification) Size() int                           { return m.SizeVT() }
func (m *CosignPublicKeyVerification) Clone() *CosignPublicKeyVerification { return m.CloneVT() }
func (m *CosignPublicKeyVerification) Marshal() ([]byte, error)            { return m.MarshalVT() }
func (m *CosignPublicKeyVerification) Unmarshal(dAtA []byte) error         { return m.UnmarshalVT(dAtA) }

func (m *CosignPublicKeyVerification_PublicKey) Size() int { return m.SizeVT() }
func (m *CosignPublicKeyVerification_PublicKey) Clone() *CosignPublicKeyVerification_PublicKey {
	return m.CloneVT()
}
func (m *CosignPublicKeyVerification_PublicKey) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *CosignPublicKeyVerification_PublicKey) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *CosignCertificateVerification) Size() int                             { return m.SizeVT() }
func (m *CosignCertificateVerification) Clone() *CosignCertificateVerification { return m.CloneVT() }
func (m *CosignCertificateVerification) Marshal() ([]byte, error)              { return m.MarshalVT() }
func (m *CosignCertificateVerification) Unmarshal(dAtA []byte) error           { return m.UnmarshalVT(dAtA) }

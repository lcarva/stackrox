// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v2

func (m *ComplianceScanStatsShim) Size() int                       { return m.SizeVT() }
func (m *ComplianceScanStatsShim) Clone() *ComplianceScanStatsShim { return m.CloneVT() }
func (m *ComplianceScanStatsShim) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *ComplianceScanStatsShim) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *ComplianceProfileScanStats) Size() int                          { return m.SizeVT() }
func (m *ComplianceProfileScanStats) Clone() *ComplianceProfileScanStats { return m.CloneVT() }
func (m *ComplianceProfileScanStats) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *ComplianceProfileScanStats) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ComplianceClusterScanStats) Size() int                          { return m.SizeVT() }
func (m *ComplianceClusterScanStats) Clone() *ComplianceClusterScanStats { return m.CloneVT() }
func (m *ComplianceClusterScanStats) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *ComplianceClusterScanStats) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ComplianceScanClusterRequest) Size() int                            { return m.SizeVT() }
func (m *ComplianceScanClusterRequest) Clone() *ComplianceScanClusterRequest { return m.CloneVT() }
func (m *ComplianceScanClusterRequest) Marshal() ([]byte, error)             { return m.MarshalVT() }
func (m *ComplianceScanClusterRequest) Unmarshal(dAtA []byte) error          { return m.UnmarshalVT(dAtA) }

func (m *ListComplianceProfileScanStatsResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceProfileScanStatsResponse) Clone() *ListComplianceProfileScanStatsResponse {
	return m.CloneVT()
}
func (m *ListComplianceProfileScanStatsResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceProfileScanStatsResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ListComplianceClusterProfileStatsResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceClusterProfileStatsResponse) Clone() *ListComplianceClusterProfileStatsResponse {
	return m.CloneVT()
}
func (m *ListComplianceClusterProfileStatsResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceClusterProfileStatsResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

func (m *ListComplianceClusterScanStatsResponse) Size() int { return m.SizeVT() }
func (m *ListComplianceClusterScanStatsResponse) Clone() *ListComplianceClusterScanStatsResponse {
	return m.CloneVT()
}
func (m *ListComplianceClusterScanStatsResponse) Marshal() ([]byte, error) { return m.MarshalVT() }
func (m *ListComplianceClusterScanStatsResponse) Unmarshal(dAtA []byte) error {
	return m.UnmarshalVT(dAtA)
}

// Code generated by protoc-gen-go-compat. DO NOT EDIT.

package v2

func (m *ReportConfiguration) Size() int                   { return m.SizeVT() }
func (m *ReportConfiguration) Clone() *ReportConfiguration { return m.CloneVT() }
func (m *ReportConfiguration) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ReportConfiguration) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *VulnerabilityReportFilters) Size() int                          { return m.SizeVT() }
func (m *VulnerabilityReportFilters) Clone() *VulnerabilityReportFilters { return m.CloneVT() }
func (m *VulnerabilityReportFilters) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *VulnerabilityReportFilters) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ReportSchedule) Size() int                   { return m.SizeVT() }
func (m *ReportSchedule) Clone() *ReportSchedule      { return m.CloneVT() }
func (m *ReportSchedule) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ReportSchedule) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ReportSchedule_DaysOfWeek) Size() int                         { return m.SizeVT() }
func (m *ReportSchedule_DaysOfWeek) Clone() *ReportSchedule_DaysOfWeek { return m.CloneVT() }
func (m *ReportSchedule_DaysOfWeek) Marshal() ([]byte, error)          { return m.MarshalVT() }
func (m *ReportSchedule_DaysOfWeek) Unmarshal(dAtA []byte) error       { return m.UnmarshalVT(dAtA) }

func (m *ReportSchedule_DaysOfMonth) Size() int                          { return m.SizeVT() }
func (m *ReportSchedule_DaysOfMonth) Clone() *ReportSchedule_DaysOfMonth { return m.CloneVT() }
func (m *ReportSchedule_DaysOfMonth) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *ReportSchedule_DaysOfMonth) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ResourceScope) Size() int                   { return m.SizeVT() }
func (m *ResourceScope) Clone() *ResourceScope       { return m.CloneVT() }
func (m *ResourceScope) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ResourceScope) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *CollectionReference) Size() int                   { return m.SizeVT() }
func (m *CollectionReference) Clone() *CollectionReference { return m.CloneVT() }
func (m *CollectionReference) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *CollectionReference) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *NotifierConfiguration) Size() int                     { return m.SizeVT() }
func (m *NotifierConfiguration) Clone() *NotifierConfiguration { return m.CloneVT() }
func (m *NotifierConfiguration) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *NotifierConfiguration) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *EmailNotifierConfiguration) Size() int                          { return m.SizeVT() }
func (m *EmailNotifierConfiguration) Clone() *EmailNotifierConfiguration { return m.CloneVT() }
func (m *EmailNotifierConfiguration) Marshal() ([]byte, error)           { return m.MarshalVT() }
func (m *EmailNotifierConfiguration) Unmarshal(dAtA []byte) error        { return m.UnmarshalVT(dAtA) }

func (m *ListReportConfigurationsResponse) Size() int { return m.SizeVT() }
func (m *ListReportConfigurationsResponse) Clone() *ListReportConfigurationsResponse {
	return m.CloneVT()
}
func (m *ListReportConfigurationsResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ListReportConfigurationsResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *CountReportConfigurationsResponse) Size() int { return m.SizeVT() }
func (m *CountReportConfigurationsResponse) Clone() *CountReportConfigurationsResponse {
	return m.CloneVT()
}
func (m *CountReportConfigurationsResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *CountReportConfigurationsResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *GetReportHistoryRequest) Size() int                       { return m.SizeVT() }
func (m *GetReportHistoryRequest) Clone() *GetReportHistoryRequest { return m.CloneVT() }
func (m *GetReportHistoryRequest) Marshal() ([]byte, error)        { return m.MarshalVT() }
func (m *GetReportHistoryRequest) Unmarshal(dAtA []byte) error     { return m.UnmarshalVT(dAtA) }

func (m *ReportHistoryResponse) Size() int                     { return m.SizeVT() }
func (m *ReportHistoryResponse) Clone() *ReportHistoryResponse { return m.CloneVT() }
func (m *ReportHistoryResponse) Marshal() ([]byte, error)      { return m.MarshalVT() }
func (m *ReportHistoryResponse) Unmarshal(dAtA []byte) error   { return m.UnmarshalVT(dAtA) }

func (m *ReportStatusResponse) Size() int                    { return m.SizeVT() }
func (m *ReportStatusResponse) Clone() *ReportStatusResponse { return m.CloneVT() }
func (m *ReportStatusResponse) Marshal() ([]byte, error)     { return m.MarshalVT() }
func (m *ReportStatusResponse) Unmarshal(dAtA []byte) error  { return m.UnmarshalVT(dAtA) }

func (m *CollectionSnapshot) Size() int                   { return m.SizeVT() }
func (m *CollectionSnapshot) Clone() *CollectionSnapshot  { return m.CloneVT() }
func (m *CollectionSnapshot) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *CollectionSnapshot) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ReportSnapshot) Size() int                   { return m.SizeVT() }
func (m *ReportSnapshot) Clone() *ReportSnapshot      { return m.CloneVT() }
func (m *ReportSnapshot) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ReportSnapshot) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *ReportStatus) Size() int                   { return m.SizeVT() }
func (m *ReportStatus) Clone() *ReportStatus        { return m.CloneVT() }
func (m *ReportStatus) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *ReportStatus) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *RunReportRequest) Size() int                   { return m.SizeVT() }
func (m *RunReportRequest) Clone() *RunReportRequest    { return m.CloneVT() }
func (m *RunReportRequest) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *RunReportRequest) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *RunReportResponse) Size() int                   { return m.SizeVT() }
func (m *RunReportResponse) Clone() *RunReportResponse   { return m.CloneVT() }
func (m *RunReportResponse) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *RunReportResponse) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

func (m *DeleteReportRequest) Size() int                   { return m.SizeVT() }
func (m *DeleteReportRequest) Clone() *DeleteReportRequest { return m.CloneVT() }
func (m *DeleteReportRequest) Marshal() ([]byte, error)    { return m.MarshalVT() }
func (m *DeleteReportRequest) Unmarshal(dAtA []byte) error { return m.UnmarshalVT(dAtA) }

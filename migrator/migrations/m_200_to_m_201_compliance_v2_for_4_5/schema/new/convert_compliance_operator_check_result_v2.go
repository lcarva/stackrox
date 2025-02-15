// Code generated by pg-bindings generator. DO NOT EDIT.
package new

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/protocompat"
)

// ConvertComplianceOperatorCheckResultV2FromProto converts a `*storage.ComplianceOperatorCheckResultV2` to Gorm model
func ConvertComplianceOperatorCheckResultV2FromProto(obj *storage.ComplianceOperatorCheckResultV2) (*ComplianceOperatorCheckResultV2, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &ComplianceOperatorCheckResultV2{
		ID:             obj.GetId(),
		CheckID:        obj.GetCheckId(),
		CheckName:      obj.GetCheckName(),
		ClusterID:      obj.GetClusterId(),
		Status:         obj.GetStatus(),
		Severity:       obj.GetSeverity(),
		CreatedTime:    protocompat.NilOrTime(obj.GetCreatedTime()),
		ScanConfigName: obj.GetScanConfigName(),
		Rationale:      obj.GetRationale(),
		ScanRefID:      obj.GetScanRefId(),
		RuleRefID:      obj.GetRuleRefId(),
		Serialized:     serialized,
	}
	return model, nil
}

// ConvertComplianceOperatorCheckResultV2ToProto converts Gorm model `ComplianceOperatorCheckResultV2` to its protobuf type object
func ConvertComplianceOperatorCheckResultV2ToProto(m *ComplianceOperatorCheckResultV2) (*storage.ComplianceOperatorCheckResultV2, error) {
	var msg storage.ComplianceOperatorCheckResultV2
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}

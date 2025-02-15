// Code generated by pg-bindings generator. DO NOT EDIT.
package new

import (
	"github.com/stackrox/rox/generated/storage"
)

// ConvertComplianceOperatorProfileV2FromProto converts a `*storage.ComplianceOperatorProfileV2` to Gorm model
func ConvertComplianceOperatorProfileV2FromProto(obj *storage.ComplianceOperatorProfileV2) (*ComplianceOperatorProfileV2, error) {
	serialized, err := obj.MarshalVT()
	if err != nil {
		return nil, err
	}
	model := &ComplianceOperatorProfileV2{
		ID:             obj.GetId(),
		ProfileID:      obj.GetProfileId(),
		Name:           obj.GetName(),
		ProfileVersion: obj.GetProfileVersion(),
		ProductType:    obj.GetProductType(),
		Standard:       obj.GetStandard(),
		ClusterID:      obj.GetClusterId(),
		ProfileRefID:   obj.GetProfileRefId(),
		Serialized:     serialized,
	}
	return model, nil
}

// ConvertComplianceOperatorProfileV2_RuleFromProto converts a `*storage.ComplianceOperatorProfileV2_Rule` to Gorm model
func ConvertComplianceOperatorProfileV2_RuleFromProto(obj *storage.ComplianceOperatorProfileV2_Rule, idx int, complianceOperatorProfileV2ID string) (*ComplianceOperatorProfileV2Rules, error) {
	model := &ComplianceOperatorProfileV2Rules{
		ComplianceOperatorProfileV2ID: complianceOperatorProfileV2ID,
		Idx:                           idx,
		RuleName:                      obj.GetRuleName(),
	}
	return model, nil
}

// ConvertComplianceOperatorProfileV2ToProto converts Gorm model `ComplianceOperatorProfileV2` to its protobuf type object
func ConvertComplianceOperatorProfileV2ToProto(m *ComplianceOperatorProfileV2) (*storage.ComplianceOperatorProfileV2, error) {
	var msg storage.ComplianceOperatorProfileV2
	if err := msg.Unmarshal(m.Serialized); err != nil {
		return nil, err
	}
	return &msg, nil
}

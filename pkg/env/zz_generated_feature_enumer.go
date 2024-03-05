// Code generated by "enumer -type Feature -output zz_generated_feature_enumer.go"; DO NOT EDIT.

package env

import (
	"fmt"
)

const _FeatureName = "FeatureDisableDenyAssignmentsFeatureDisableSignedCertificatesFeatureEnableDevelopmentAuthorizerFeatureRequireD2sV3WorkersFeatureDisableReadinessDelayFeatureEnableOCMEndpoints"

var _FeatureIndex = [...]uint8{0, 29, 61, 95, 121, 149, 174}

func (i Feature) String() string {
	if i < 0 || i >= Feature(len(_FeatureIndex)-1) {
		return fmt.Sprintf("Feature(%d)", i)
	}
	return _FeatureName[_FeatureIndex[i]:_FeatureIndex[i+1]]
}

var _FeatureValues = []Feature{0, 1, 2, 3, 4, 5}

var _FeatureNameToValueMap = map[string]Feature{
	_FeatureName[0:29]:    0,
	_FeatureName[29:61]:   1,
	_FeatureName[61:95]:   2,
	_FeatureName[95:121]:  3,
	_FeatureName[121:149]: 4,
	_FeatureName[149:174]: 5,
}

// FeatureString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FeatureString(s string) (Feature, error) {
	if val, ok := _FeatureNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Feature values", s)
}

// FeatureValues returns all values of the enum
func FeatureValues() []Feature {
	return _FeatureValues
}

// IsAFeature returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Feature) IsAFeature() bool {
	for _, v := range _FeatureValues {
		if i == v {
			return true
		}
	}
	return false
}

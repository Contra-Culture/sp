package asset

type (
	AttributeName string
	AttributeType struct {
		Name     AttributeName
		Children map[string]AttributeType
	}
	AssetAttributeTypesSet map[AttributeName]AttributeType
	AssetAttribute         struct {
		Type     *AttributeType
		Value    string
		Children map[string]AssetAttribute
	}
)

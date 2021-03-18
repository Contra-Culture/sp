package asset

type (
	AttributeName string
	AttributeType struct {
		Name     AttributeName
		Children map[AttributeName]AttributeType
	}
	AssetAttributeTypes map[AttributeName]AttributeType
	AssetAttribute      struct {
		Type     *AttributeType
		Value    string
		Children map[AttributeName]AssetAttribute
	}
	AssetAttributes map[AttributeName]AssetAttribute
)

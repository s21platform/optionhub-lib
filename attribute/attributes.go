package attribute

// Регистрация атрибута в соответствии с ID атрибута
const (
	Attribute_Name     = 1
	Attribute_Surname  = 2
	Attribute_Patronym = 3
	Attribute_Age      = 4
	Attribute_City     = 5
	Attribute_OS       = 6
	Attribute_Hobby    = 7
)

type AttributeType string

const (
	AttributeType_Int     = AttributeType("integer")
	AttributeType_String  = AttributeType("string")
	AttributeType_IntEnum = AttributeType("int_enum")
	AttributeType_Unknown = AttributeType("unknown")
)

func AttributeTypeByValue(attribute int64) AttributeType {
	switch attribute {
	case Attribute_Name,
		Attribute_Surname,
		Attribute_Patronym:
		return AttributeType_String
	case Attribute_Age,
		Attribute_City:
		return AttributeType_Int
	case Attribute_OS,
		Attribute_Hobby:
		return AttributeType_IntEnum
	}
	return AttributeType_Unknown
}

package users

// Регистрация атрибута в соответствии с ID атрибута
const (
	Attribute_Name      = 1
	Attribute_Surname   = 2
	Attribute_Patronym  = 3
	Attribute_Age       = 4
	Attribute_City      = 5
	Attribute_OS        = 6
	Attribute_Hobby     = 7
	Attribute_Birthdate = 8
	Attribute_Telegram  = 9
	Attribute_Git       = 10
)

type AttributeType string

const (
	AttributeType_Int     = AttributeType("integer")
	AttributeType_String  = AttributeType("string")
	AttributeType_IntEnum = AttributeType("int_enum")
	AttributeType_Date    = AttributeType("date")
	AttributeType_Unknown = AttributeType("unknown")
)

var AttributeTypes = map[int]AttributeType{
	Attribute_Name:     AttributeType_String,
	Attribute_Surname:  AttributeType_String,
	Attribute_Telegram: AttributeType_String,
	Attribute_Git:      AttributeType_String,
	Attribute_City:     AttributeType_Int,
}

func AttributeTypeByValue(attribute int64) AttributeType {
	switch attribute {
	case Attribute_Name,
		Attribute_Surname,
		Attribute_Patronym,
		Attribute_Telegram,
		Attribute_Git:
		return AttributeType_String
	case Attribute_Age,
		Attribute_City,
		Attribute_OS:
		return AttributeType_Int
	case Attribute_Hobby:
		return AttributeType_IntEnum
	case Attribute_Birthdate:
		return AttributeType_Date
	}
	return AttributeType_Unknown
}

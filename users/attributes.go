package users

// Регистрация атрибута в соответствии с entity_attribute_id
const (
	Attribute_Nickanme_1 = 1
	Attribute_Name_2     = 2
	Attribute_Surname_3  = 3
	Attribute_Birthday_4 = 4
	Attribute_City_5     = 5
	Attribute_Telegram_6 = 6
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
	Attribute_Name_2:     AttributeType_String,
	Attribute_Surname_3:  AttributeType_String,
	Attribute_Birthday_4: AttributeType_Date,
	Attribute_City_5:     AttributeType_Int,
	Attribute_Telegram_6: AttributeType_String,
}

func AttributeTypeByValue(attribute int64) AttributeType {
	switch attribute {
	case Attribute_Name_2,
		Attribute_Surname_3,
		Attribute_Telegram_6:
		return AttributeType_String
	case Attribute_City_5:
		return AttributeType_Int
	case Attribute_Birthday_4:
		return AttributeType_Date
	}
	return AttributeType_Unknown
}

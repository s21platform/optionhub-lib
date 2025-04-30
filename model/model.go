package model

type AttributeValue struct {
	AttributeId  int64
	ValueInt     *int64
	ValueString  *string
	ValueIntEnum []int64
}

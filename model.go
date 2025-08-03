package optionhub_lib

import "time"

type AttributeValue struct {
	AttributeId  int64      `json:"attribute_id"`
	ValueInt     *int64     `json:"value_int,omitempty"`
	ValueString  *string    `json:"value_string,omitempty"`
	ValueIntEnum []int64    `json:"value_int_enum,omitempty"`
	ValueDate    *time.Time `json:"value_date,omitempty"`
}

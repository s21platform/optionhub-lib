package optionhub_lib

import "time"

type AttributeValue struct {
	AttributeId  int64
	ValueInt     *int64
	ValueString  *string
	ValueIntEnum []int64
	ValueDate    *time.Time
}

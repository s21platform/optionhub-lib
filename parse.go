package optionhub_lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	logger_lib "github.com/s21platform/logger-lib"

	"github.com/s21platform/optionhub-lib/users"
)

func ParseAttributes(ctx context.Context, data json.RawMessage) ([]AttributeValue, error) {
	var target map[int64]json.RawMessage
	err := json.Unmarshal(data, &target)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal origin bytes: %v", err)
	}
	return parseAttributeValues(ctx, target)
}

func parseAttributeValues(ctx context.Context, data map[int64]json.RawMessage) ([]AttributeValue, error) {
	var res []AttributeValue
	for k, v := range data {
		ctx = logger_lib.WithField(ctx, "attribute_id", k)
		ctx = logger_lib.WithField(ctx, "parsing_value", string(v))
		switch users.AttributeTypeByValue(k) {
		case users.AttributeType_Int:
			val, err := parseInt(k, v)
			if err != nil {
				logger_lib.Error(logger_lib.WithField(ctx, "error", err.Error()), "failed to parse `int` value")
				return nil, err
			}
			res = append(res, val)
		case users.AttributeType_String:
			val, err := parseString(k, v)
			if err != nil {
				logger_lib.Error(logger_lib.WithField(ctx, "error", err.Error()), "failed to parse `string` value")
				return nil, err
			}
			res = append(res, val)
		case users.AttributeType_IntEnum:
			val, err := parseIntEnum(k, v)
			if err != nil {
				logger_lib.Error(logger_lib.WithField(ctx, "error", err.Error()), "failed to parse `int enum` value")
				return nil, err
			}
			res = append(res, val)
		case users.AttributeType_Date:
			val, err := parseDate(k, v)
			if err != nil {
				logger_lib.Error(logger_lib.WithField(ctx, "error", err.Error()), "failed to parse `date` value")
				return nil, err
			}
			res = append(res, val)
		default:
			logger_lib.Error(ctx, fmt.Sprintf("failed to retrieve `unknown` value for attribute_id: %d", k))
			return nil, fmt.Errorf("failed to retrieve `unknown` value for attribute_id: %d", k)
		}
	}
	return res, nil
}

func parseInt(attributeId int64, data json.RawMessage) (AttributeValue, error) {
	bytes.Replace(data, []byte(`"`), []byte{}, -1)
	var result int64
	err := json.Unmarshal(data, &result)
	if err != nil {
		return AttributeValue{}, fmt.Errorf("failed to parse `int` attribute - %d: %v", attributeId, err)
	}
	return AttributeValue{
		AttributeId: attributeId,
		ValueInt:    &result,
	}, nil
}

func parseString(attributeId int64, data json.RawMessage) (AttributeValue, error) {
	var result string
	err := json.Unmarshal(data, &result)
	if err != nil {
		return AttributeValue{}, fmt.Errorf("failed to parse `string` attribute - %d: %v", attributeId, err)
	}
	return AttributeValue{
		AttributeId: attributeId,
		ValueString: &result,
	}, nil
}

func parseIntEnum(attributeId int64, data json.RawMessage) (AttributeValue, error) {
	result := []int64{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return AttributeValue{}, fmt.Errorf("failed to parse `int enum` attribute - %d: %v", attributeId, err)
	}
	return AttributeValue{
		AttributeId:  attributeId,
		ValueIntEnum: result,
	}, nil
}

func parseDate(attributeId int64, data json.RawMessage) (AttributeValue, error) {
	result := time.Time{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return AttributeValue{}, fmt.Errorf("failed to parse `date` attribute - %d: %v", attributeId, err)
	}
	return AttributeValue{
		AttributeId: attributeId,
		ValueDate:   &result,
	}, nil
}

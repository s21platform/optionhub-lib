package optionhub_lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/s21platform/optionhub-lib/users"
	"time"

	logger_lib "github.com/s21platform/logger-lib"
)

type OptionhubParser struct {
	logger logger_lib.LoggerInterface
}

func New(logger logger_lib.LoggerInterface) *OptionhubParser {
	return &OptionhubParser{
		logger: logger,
	}
}

func (op *OptionhubParser) ParseAttributes(data json.RawMessage) ([]AttributeValue, error) {
	var target map[int64]json.RawMessage
	err := json.Unmarshal(data, &target)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal origin bytes: %v", err)
	}
	return op.parseAttributeValues(target)
}

func (op *OptionhubParser) parseAttributeValues(data map[int64]json.RawMessage) ([]AttributeValue, error) {
	var res []AttributeValue
	for k, v := range data {
		switch users.AttributeTypeByValue(k) {
		case users.AttributeType_Int:
			val, err := op.parseInt(k, v)
			if err != nil {
				op.logger.Error(fmt.Sprintf("failed to parse `int` value: %v", err))
				continue
			}
			res = append(res, val)
		case users.AttributeType_String:
			val, err := op.parseString(k, v)
			if err != nil {
				op.logger.Error(fmt.Sprintf("failed to parse `string` value: %v", err))
				continue
			}
			res = append(res, val)
		case users.AttributeType_IntEnum:
			val, err := op.parseIntEnum(k, v)
			if err != nil {
				op.logger.Error(fmt.Sprintf("failed to parse `int enum` value: %v", err))
				continue
			}
			res = append(res, val)
		case users.AttributeType_Date:
			val, err := op.parseDate(k, v)
			if err != nil {
				op.logger.Error(fmt.Sprintf("failed to parse `date` value: %v", err))
				continue
			}
			res = append(res, val)
		default:
			op.logger.Error(fmt.Sprintf("failed to retrieve `unknown` value for attribute_id: %d", k))
		}
	}
	return res, nil
}

func (op *OptionhubParser) parseInt(attributeId int64, data json.RawMessage) (AttributeValue, error) {
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

func (op *OptionhubParser) parseString(attributeId int64, data json.RawMessage) (AttributeValue, error) {
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

func (op *OptionhubParser) parseIntEnum(attributeId int64, data json.RawMessage) (AttributeValue, error) {
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

func (op *OptionhubParser) parseDate(attributeId int64, data json.RawMessage) (AttributeValue, error) {
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

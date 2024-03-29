package binding

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"go.uber.org/multierr"
)

func RegisterValidation(valid *validator.Validate) error {
	valid.RegisterCustomTypeFunc(
		func(field reflect.Value) any {
			if valuer, ok := field.Interface().(decimal.Decimal); ok {
				return valuer.String()
			}
			return nil
		},
		decimal.Decimal{},
	)

	err := multierr.Combine(
		valid.RegisterValidation("decimal", ValidIsDecimal),
		valid.RegisterValidation("decimal_gt", ValidIsDecimalGt),
		valid.RegisterValidation("decimal_gte", ValidIsDecimalGte),
		valid.RegisterValidation("decimal_lt", ValidIsDecimalLt),
		valid.RegisterValidation("decimal_lte", ValidIsDecimalLte),
		valid.RegisterValidation("decimal_min", ValidDecimalMinOf),
		valid.RegisterValidation("decimal_max", ValidDecimalMaxOf),
		valid.RegisterValidation("number_gt0", ValidNumberGt0),
		valid.RegisterValidation("number_gte0", ValidNumberGte0),
	)
	if err != nil {
		return fmt.Errorf("validator: register validation failed, %w", err)
	}
	return nil
}

// ValidIsDecimal 校验是否为字符串 decimal
func ValidIsDecimal(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() == reflect.String {
		return IsDecimal(field.String())
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// ValidIsDecimalGt 校验是否为字符串数字且大于0
func ValidIsDecimalGt(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if field.Kind() == reflect.String {
		return IsDecimalGt(field.String(), param)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// ValidIsDecimalGte 校验是否为字符串数字且大于等于0
func ValidIsDecimalGte(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if field.Kind() == reflect.String {
		return IsDecimalGte(field.String(), param)
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func ValidIsDecimalLt(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if field.Kind() == reflect.String {
		return IsDecimalLt(field.String(), param)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func ValidIsDecimalLte(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	if field.Kind() == reflect.String {
		return IsDecimalLte(field.String(), param)
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func ValidDecimalMinOf(fl validator.FieldLevel) bool {
	return ValidIsDecimalGte(fl)
}

func ValidDecimalMaxOf(fl validator.FieldLevel) bool {
	return ValidIsDecimalLte(fl)
}

func ValidNumberGt0(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.String {
		return IsNumberGt0(field.String())
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func ValidNumberGte0(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.String {
		return IsNumberGte0(field.String())
	}
	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

package validator

import (
	"context"

	"github.com/go-playground/locales/en"
	playgroundTranslator "github.com/go-playground/universal-translator"
	playgroundValidator "github.com/go-playground/validator/v10"
	englishTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/senither/custodian/server/utils"
)

var (
	validate            *playgroundValidator.Validate
	translate           *playgroundTranslator.Translator
	universalTranslator *playgroundTranslator.UniversalTranslator
)

func Parse(ctx context.Context, values interface{}) *fiber.Map {
	errs := validator().StructCtx(ctx, values)
	if errs == nil {
		return nil
	}

	errors := make(fiber.Map)
	for _, err := range errs.(playgroundValidator.ValidationErrors) {
		key := utils.ToSnakeCase(err.Field())
		if _, ok := errors[key]; !ok {
			errors[key] = []string{}
		}

		errors[key] = append(errors[key].([]string), err.Translate(*translate))
	}

	return &errors
}

func validator() *playgroundValidator.Validate {
	if validate == nil {
		en := en.New()
		universalTranslator = playgroundTranslator.New(en, en)

		trans, _ := universalTranslator.GetTranslator("en")
		translate = &trans

		validate = playgroundValidator.New()
		englishTranslations.RegisterDefaultTranslations(validate, trans)
	}

	return validate
}

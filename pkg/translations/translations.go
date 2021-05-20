// Copyright (c) 2021 Circutor S.A. All rights reserved.

package translations

import (
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/translations/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// TranslateEnError translate english error of validator structs.
func TranslateEnError(err error, v *validator.Validate) (errs []string) {
	trans := config.CreateEnTranslation(v)

	return translateError(err, trans)
}

// translateError translate error of validator structs.
func translateError(err error, trans ut.Translator) (errs []string) {
	if err != nil {
		//nolint:errorlint
		validatorErrs := err.(validator.ValidationErrors)

		for _, e := range validatorErrs {
			translatedErr := errors.NewErrFound(e.Translate(trans))
			errs = append(errs, translatedErr.Error())
		}

		return errs
	}

	return nil
}

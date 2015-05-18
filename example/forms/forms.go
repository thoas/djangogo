package forms

import (
	"github.com/bluele/gforms"
)

var authentication = gforms.DefineForm(
	gforms.NewFields(
		gforms.NewTextField(
			"email",
			gforms.Validators{
				gforms.Required(),
				gforms.MinLengthValidator(4),
				gforms.EmailValidator(),
			},
		),
		gforms.NewTextField(
			"password",
			gforms.Validators{
				gforms.Required(),
				gforms.MinLengthValidator(4),
				gforms.MaxLengthValidator(16),
			},
			gforms.PasswordInputWidget(map[string]string{}),
		),
	))

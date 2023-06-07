package validators

import (
	"github.com/anfern777/cicd-dashboard/service"
	"github.com/go-playground/validator/v10"
)

func ValidateUserRole(field validator.FieldLevel) bool {
	roles := service.NewUserService().GetAllUserRoles()

	for _, role := range roles {
		if role == field.Field().String() {
			return true
		}
	}
	return false
}

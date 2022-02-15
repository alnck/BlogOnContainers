package models

import . "blog-on-containers/constants"

type LoginRequest struct {
	UserName   string `json:"UserName" form:"UserName" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}

func (lr LoginRequest) IsValid() (errs []ErrorDetail) {
	if lr.UserName == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_USERNAME_IS_EMPTY})
	}
	if len(lr.UserName) < 5 || len(lr.UserName) > 10 {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_USERNAME_IS_NOT_AVAİBLE})
	}
	if lr.Password == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_PASSWORD_IS_EMPTY})
	}
	if len(lr.Password) < 5 || len(lr.Password) > 10 {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_PASSWORD_IS_NOT_AVAİBLE})
	}
	return errs
}

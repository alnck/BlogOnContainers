package models

type LoginRequest struct {
	UserName   string `json:"UserName" form:"UserName" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}

func (lr LoginRequest) IsValid() (errs []ErrorDetail) {
	if lr.UserName == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The username is required!"})
	}
	if len(lr.UserName) < 5 || len(lr.UserName) > 10 {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The username field must be between 5-10 chars!"})
	}
	if lr.Password == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The password is required!"})
	}
	if len(lr.Password) < 5 || len(lr.Password) > 10 {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The password field must be between 5-10 chars!"})
	}
	return errs
}

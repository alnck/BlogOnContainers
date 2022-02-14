package models

type StoryRequest struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

func (sr StoryRequest) IsValid() (errs []ErrorDetail) {
	if sr.Title == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The title is required!"})
	}
	if sr.Content == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: "The content is required!"})
	}
	return errs
}

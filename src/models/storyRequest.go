package models

import . "blog-on-containers/constants"

type StoryRequest struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

func (sr StoryRequest) IsValid() (errs []ErrorDetail) {
	if sr.Title == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_STORY_TITLE_IS_EMPTY})
	}
	if sr.Content == "" {
		errs = append(errs, ErrorDetail{ErrorType: ErrorTypeValidation, ErrorMessage: ERROR_MESSAGE_STORY_CONTENT_IS_EMPTY})
	}
	return errs
}

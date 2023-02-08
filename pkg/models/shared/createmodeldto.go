package shared

type CreateModelDto struct {
	SubjectIdentifier *string `json:"subjectIdentifier,omitempty"`
	SubjectKeyword    string  `json:"subjectKeyword"`
	Title             string  `json:"title"`
}

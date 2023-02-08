package shared

type ModelEntity struct {
	ID                string `json:"id"`
	SubjectIdentifier string `json:"subjectIdentifier"`
	SubjectKeyword    string `json:"subjectKeyword"`
	Title             string `json:"title"`
}

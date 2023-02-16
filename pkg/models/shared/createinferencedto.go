package shared

type CreateInferenceDto struct {
	Height         *float64 `json:"height,omitempty"`
	NegativePrompt *string  `json:"negativePrompt,omitempty"`
	NumberOfImages *float64 `json:"numberOfImages,omitempty"`
	Prompt         string   `json:"prompt"`
	PromptStrength *float64 `json:"promptStrength,omitempty"`
	RestoreFaces   *bool    `json:"restoreFaces,omitempty"`
	Seed           *float64 `json:"seed,omitempty"`
	Steps          *float64 `json:"steps,omitempty"`
	Version        *string  `json:"version,omitempty"`
	WebhookURL     *string  `json:"webhookUrl,omitempty"`
	Width          *float64 `json:"width,omitempty"`
}

package aireviewer

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/responses"
)

type Severity string

const (
	SeverityLow    Severity = "low"
	SeverityMedium Severity = "medium"
	SeverityHigh   Severity = "high"
)

type Review struct {
	Findings []Finding
}

type Finding struct {
	File        string   `json:"file"`
	Line        int      `json:"line"`
	Severity    Severity `json:"severity"`
	Title       string   `json:"title"`
	Explanation string   `json:"explanation"`
	Suggestion  string   `json:"suggestion"`
}

func schemaBuilderForReview() map[string]any {
	return map[string]any{
		"type":                 "object",
		"additionalProperties": false,
		"required":             []string{"findings"},
		"properties": map[string]any{
			"findings": map[string]any{
				"type": "array",
				"items": map[string]any{
					"$ref": "#/$defs/finding",
				},
			},
		},
		"$defs": map[string]any{
			"finding": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"file": map[string]any{"type": "string"},
					"line": map[string]any{"type": "integer"},
					"severity": map[string]any{
						"type":        "string",
						"enum":        []string{string(SeverityLow), string(SeverityMedium), string(SeverityHigh)},
						"description": "severity of the bug or issue found in the code.",
					},
					"title": map[string]any{"type": "string"},
					"explanation": map[string]any{
						"type":        "string",
						"description": "description of the issue or bug found in the code",
					},
					"suggestion": map[string]any{
						"type":        []string{"string"},
						"description": "code suggestion to fix the issue",
					},
				},
				"additionalProperties": false,
				"required":             []string{"file", "line", "severity", "title", "explanation", "suggestion"},
			},
		},
	}
}

func clientResponse(f []byte, instruction string) (*responses.Response, error) {
	client := openai.NewClient()
	systemMsg := responses.ResponseInputItemParamOfMessage(
		responses.ResponseInputMessageContentListParam{responses.ResponseInputContentParamOfInputText(instruction)},
		responses.EasyInputMessageRoleSystem,
	)
	userMsg := responses.ResponseInputItemParamOfMessage(
		responses.ResponseInputMessageContentListParam{responses.ResponseInputContentParamOfInputText(string(f))},
		responses.EasyInputMessageRoleUser,
	)
	jsonConfig := &responses.ResponseFormatTextJSONSchemaConfigParam{
		Name:   "ai_reviewer_v0",
		Schema: schemaBuilderForReview(),
		Strict: openai.Bool(true),
	}
	response, err := client.Responses.New(context.Background(), responses.ResponseNewParams{
		Model: openai.ChatModelGPT5_6Terra,
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{
				systemMsg,
				userMsg,
			},
		},
		Text: responses.ResponseTextConfigParam{
			Format: responses.ResponseFormatTextConfigUnionParam{
				OfJSONSchema: jsonConfig,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if response.Status == responses.ResponseStatusFailed ||
		response.Status == responses.ResponseStatusCancelled ||
		response.Status == responses.ResponseStatusIncomplete {
		return nil, errors.New("failed to get review")
	}
	return response, nil

}

func GetReviewFor(f []byte, instruction string) (Review, error) {
	response, err := clientResponse(f, instruction)
	if err != nil {
		return Review{}, err
	}
	var reviews Review
	if err := json.Unmarshal([]byte(response.OutputText()), &reviews); err != nil {
		return Review{}, err
	}

	return reviews, nil
}

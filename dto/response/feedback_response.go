package response

import "github.com/kelompok4-loyaltypointagent/backend/models"

type FeedbackResponse struct {
	User                 *UserResponse `json:"user"`
	IsInformationHelpful *bool         `json:"is_information_helpful"`
	IsArticleHelpful     *bool         `json:"is_article_helpful"`
	IsArticleEasyToFind  *bool         `json:"is_article_easy_to_find"`
	IsDesignGood         *bool         `json:"is_design_good"`
	Review               string        `json:"review"`
}

func NewFeedbackResponse(feedback models.Feedbacks) *FeedbackResponse {
	response := &FeedbackResponse{
		IsInformationHelpful: feedback.IsInformationHelpful,
		IsArticleHelpful:     feedback.IsArticleHelpful,
		IsArticleEasyToFind:  feedback.IsArticleEasyToFind,
		IsDesignGood:         feedback.IsDesignGood,
		Review:               feedback.Review,
	}

	if response.User != nil {
		response.User = NewUserResponse(*feedback.User)
	}

	return response
}

func NewFeedbackResponseList(feedback []models.Feedbacks) *[]FeedbackResponse {
	var list []FeedbackResponse
	for _, v := range feedback {
		list = append(list, *NewFeedbackResponse(v))
	}
	return &list
}

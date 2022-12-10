package payload

type FeedbackPayload struct {
	IsInformationHelpful *bool  `json:"is_information_helpful"`
	IsArticleHelpful     *bool  `json:"is_article_helpful"`
	IsArticleEasyToFind  *bool  `json:"is_article_easy_to_find"`
	IsDesignGood         *bool  `json:"is_design_good"`
	Review               string `json:"review" validate:"required"`
}

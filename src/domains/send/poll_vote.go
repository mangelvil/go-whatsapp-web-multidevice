package send

type PollVoteRequest struct {
	BaseRequest
	MessageID   string   `json:"message_id" form:"message_id"`
	PollSender  string   `json:"poll_sender" form:"poll_sender"`
	OptionNames []string `json:"option_names" form:"option_names"`
}

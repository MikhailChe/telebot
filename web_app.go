package telebot

// WebApp represents a parameter of the inline keyboard button
// or the keyboard button used to launch Web App.
type WebApp struct {
	URL string `json:"url"`
}

// WebAppMessage describes an inline message sent by a Web App on behalf of a user.
type WebAppMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// WebAppData object represents a data sent from a Web App to the bot
type WebAppData struct {
	// The data. Be aware that a bad client can send arbitrary data in this field
	Data string `json:"data"`
	// Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
	Text string `json:"button_text"`
}

// WriteAccessAllowed represents a service message about a user allowing a bot to write messages after adding it to the attachment menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess.
// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	// Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	FromRequest bool `json:"from_request,omitempty"`
	// Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	WebAppName string `json:"web_app_name,omitempty"`
	// Optional. True, if the access was granted when the bot was added to the attachment or side menu
	FromAttachmentMenu bool `json:"from_attachment_menu,omitempty"`
}

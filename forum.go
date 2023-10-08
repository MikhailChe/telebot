package telebot

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name"`
	// Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconEmoji string `json:"icon_custom_emoji_id"`
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat. Currently holds no information.
// https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// ForumTopicEdited represents a service message about an edited forum topic
// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	// Optional. New name of the topic, if it was edited
	Name string `json:"name,omitempty"`
	// Optional. Unique identifier of the custom emoji shown as the topic icon
	IconEmoji string `json:"icon_custom_emoji_id"`
}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat. Currently holds no information.
// https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat. Currently holds no information
// https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat. Currently holds no information.
// https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct{}

package message

const (
	// Used for reference table.
	tableDefMessages = "messages"

	// List of queries for messages.
	querySelectMessagesClause = `SELECT
									id,
									channel_id,
									sender_type,
									sender_id,
									content_type,
									content,
									created_at,
									updated_at
								FROM ` + tableDefMessages

	querySelectMessageById        = querySelectMessagesClause + ` WHERE id = $1`
	querySelectMessageByChannelId = querySelectMessagesClause + ` WHERE channel_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
)

const (
	// Used for reference table.
	tableDefMessageInputs = "message_inputs"

	// List of queries for message inputs.
	querySelectMessageInputsClause = `SELECT
									id,
									channel_id,
									source,
									sender,
									receiver,
									receiver_pronoun,
									created_at,
									updated_at
								FROM ` + tableDefMessageInputs

	querySelectMessageInputById        = querySelectMessageInputsClause + ` WHERE id = $1`
	querySelectMessageInputByChannelId = querySelectMessageInputsClause + ` WHERE channel_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
)

const (
	// Used for reference table.
	tableDefMessageSources = "message_sources"

	// List of queries for messages.
	querySelectMessageSourcesClause = `SELECT
									id,
									message_input_id,
									sender,
									content_type,
									content,
									sent_at,
									created_at,
									updated_at
								FROM ` + tableDefMessageSources

	querySelectMessageSourceById             = querySelectMessageSourcesClause + ` WHERE id = $1`
	querySelectMessageSourceByMessageInputId = querySelectMessageSourcesClause + ` WHERE message_input_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
)

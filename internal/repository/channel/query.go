package channel

const (
	// Used for reference table.
	tableDefChannels = "channels"

	// List of queries for channels.
	querySelectChannelsClause = `SELECT
									id,
									user_id,
									name,
									created_at,
									updated_at
								FROM ` + tableDefChannels

	querySelectChannelByID     = querySelectChannelsClause + ` WHERE id = $1`
	querySelectChannelByUserID = querySelectChannelsClause + ` WHERE user_id = $1`
)

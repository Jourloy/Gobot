package messages

import (
	c "github.com/bwmarrin/discordgo"
)

func DiscordMessageCreate(session *c.Session, msg *c.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if msg.Author.ID == session.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if msg.Content == "ping" {
		session.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}

}

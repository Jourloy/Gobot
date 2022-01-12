package messages

import (
	"os"
	"strings"

	"github.com/Jourloy/GoLogger"

	c "github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	logger    GoLogger.Logger
	modRoleID string
	adminID   string
	prefix    string
)

func prepare() {
	logger = GoLogger.New(5, "Discord Mesg")

	if err := godotenv.Load(); err != nil {
		logger.Error(".env file not find")
		return
	}

	modRoleID = os.Getenv("DISCORD_MOD_ROLE")
	adminID = os.Getenv("DISCORD_ADMIN")
	prefix = os.Getenv("DISCORD_COMMAND_PREFIX")
}

// Check user on have moderator permissions
func isMod(m c.Member) bool {
	check := false
	for i := range m.Roles {
		if m.Roles[i] == modRoleID {
			check = true
		}
	}
	if check == false {
		if m.User.ID == adminID {
			check = true
		}
	}
	return check
}

func DiscordMessageCreate(session *c.Session, msg *c.MessageCreate) {
	prepare()

	if msg.Author.ID == session.State.User.ID {
		return
	} else if !strings.HasPrefix(msg.Content, prefix) {
		return
	}

	content := strings.Split(msg.Content, prefix)[1]
	splited := strings.Split(content, " ")
	command := splited[0]

	if command == "ping" {
		if err := session.ChannelMessageDelete(msg.ChannelID, msg.ID); err != nil {
			logger.Error(err.Error())
		}
	}

}

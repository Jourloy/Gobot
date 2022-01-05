package commands

import (
	logger "github.com/Jourloy/GoLogger"

	c "github.com/bwmarrin/discordgo"
)

var (
	commands = []*c.ApplicationCommand{
		{
			Name:        "me",
			Description: "Not test command",
		},
	}
)

func CreateCommands(s *c.Session, g string) {
	logger.New("all", "Discord commands")

	notCreatedCommands, err := CheckCommands(s, g)
	if err != "" {
		logger.Error("Error with getting commands: " + err)
		return
	}

	logger.Debug("Start creating commands...")
	for i := range notCreatedCommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, g, notCreatedCommands[i])
		if err != nil {
			logger.Error("Error with creating command: " + err.Error())
		} else {
			logger.Debug("Command successfully created: " + cmd.Name)
		}
	}

	logger.Log("Commands started")
}

func CheckCommands(s *c.Session, g string) ([]*c.ApplicationCommand, string) {
	logger.Debug("Start check commands...")

	appID := s.State.User.ID

	cmds, err := s.ApplicationCommands(appID, g)
	if err != nil {
		return nil, err.Error()
	}

	logger.Debug("Start finding duplicated commands...")

	var count int
	var duplicated []*c.ApplicationCommand

	for i := range cmds {
		name := cmds[i].Name
		count = -1

		for j := range cmds {
			if cmds[j].Name == name {
				count++
			}
		}

		if count > 0 {
			duplicated = append(duplicated, cmds[i])
		}
	}

	if len(duplicated) > 0 {
		logger.Warn("Guild have duplicated commands! Removing...")
		for i := range duplicated {
			s.ApplicationCommandDelete(appID, g, duplicated[i].ID)
		}
	}

	logger.Debug("Start finding old commands...")

	var oldCommands []*c.ApplicationCommand
	var state bool

	for i := range cmds {
		state = false
		for j := range commands {
			if commands[j].Name == cmds[i].Name {
				state = true
			}
		}
		if !state {
			oldCommands = append(oldCommands, cmds[i])
		}
	}

	if len(oldCommands) > 0 {
		logger.Warn("Guild have old commands! Removing...")
		for i := range oldCommands {
			s.ApplicationCommandDelete(appID, g, oldCommands[i].ID)
		}
	}

	logger.Debug("Start define not created commands...")

	var notCreatedCommands []*c.ApplicationCommand

	for i := range commands {
		state = false
		for j := range cmds {
			if commands[i].Name == cmds[j].Name {
				state = true
			}
		}
		if !state {
			notCreatedCommands = append(notCreatedCommands, commands[i])
		}
	}

	return notCreatedCommands, ""
}

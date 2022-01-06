package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	commands "github.com/Jourloy/Cyberbot/commands"
	logger "github.com/Jourloy/GoLogger"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var client *discordgo.Session

func init() {
	logger.New("all", "Discord Main")

	logger.Debug("Loading .env file")
	if err := godotenv.Load(); err != nil {
		logger.Error(".env file not find")
		return
	}

	logger.Debug("Loading discord client...")
	if discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN")); err != nil {
		logger.Error("Error with loading discord client")
		fmt.Println(err)
		return
	} else {
		client = discord
	}

	logger.Log("Success launch")
}

func main() {

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		commands.CreateCommands(s, os.Getenv("DISCORD_GUILD"))
	})

	client.AddHandler(messageCreate)

	client.Identify.Intents = discordgo.IntentsAll

	err := client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	logger.Log("Bot started")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	fmt.Println("")
	logger.Log("Close discord client...")
	client.Close()
	logger.Log("Shutdown")
}

func messageCreate(session *discordgo.Session, msg *discordgo.MessageCreate) {

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

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jourloy/Cyberbot/commands"
	"github.com/Jourloy/Cyberbot/messages"
	"github.com/Jourloy/GoLogger"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var client *discordgo.Session
var logger GoLogger.Logger

func init() {
	logger = GoLogger.New(5, "Discord Main")

	if err := godotenv.Load(); err != nil {
		logger.Error(".env file not find")
		return
	}

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
	client.Identify.Intents = discordgo.IntentsAll

	client.AddHandler(discordReady)
	client.AddHandler(messages.DiscordMessageCreate)

	if err := client.Open(); err != nil {
		logger.Error("Error with opening connection: " + err.Error())
		logger.Error("Shutdown")
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

func discordReady(s *discordgo.Session, r *discordgo.Ready) {
	commands.CreateCommands(s, os.Getenv("DISCORD_GUILD"))
}

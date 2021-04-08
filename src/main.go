package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/spidernest-go/logger"
)

const GUILDIDENT = "82930961544646656"
const CHANJOINPART = "334152713594208257"
const ROLESILENCED = "194607151086305282"

var discord *discordgo.Session

func main() {
	go logger.Info().Msg("AutoMod 0.2.7 Starting Up.")

	connectDatabase()

	// search for discord websocket gateway
	err := *new(error)
	discord, err = discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Initial Discord connection was refused.")
	}

	// add event and command handlers
	discord.AddHandler(evtJoin)
	discord.AddHandler(evtPart)

	discord.AddHandler(cmdSilence)

	// set intents
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	// open a new discord connection
	err = discord.Open()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Discord websocket connection could not be established.")
	}

	// stay connected until interrupted
	logger.Info().Msg("AutoMod 0.2.7 Startup Finshed.")
	logger.Info().Msg("Retrieving Tasks.")
	retrieveAllTasks()

	logger.Info().Msg("Delegating Tasks.")
	for _, task := range TASKS {
		delegateTask(task)
	}

	logger.Info().Msg("Waiting For Interrupt Signal.")
	<-make(chan struct{})
}

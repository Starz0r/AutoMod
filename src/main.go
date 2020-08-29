package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spidernest-go/logger"
	"os"
)

func main() {
	logger.Info().Msg("AutoMod 0.1.0 Starting Up.")

	// open a new discord connection
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Initial Discord connection was refused.")
	}

	err = discord.Open()
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Discord websocket connection could not be established.")
	}

	// stay connected until interrupted
	<-make(chan struct{})

}

package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spidernest-go/logger"
)

func evtJoin(s *discordgo.Session, j *discordgo.GuildMemberAdd) {
	logger.Debug().Msg("New User Join Event.")
	// build user info string
	userinfo := strings.Join([]string{"Username: ",
		j.Member.User.Username, "#",
		j.Member.User.Discriminator,
		"\nID: ", j.Member.User.ID,
		"\nTimestamp: ", time.Now().UTC().Format(time.RFC1123)}, "")

	// get avatar and set a default one if it doesn't exist
	avatar := "http://is1.mzstatic.com/image/thumb/Purple117/v4/a1/d8/3a/a1d83a42-e84e-5965-c006-610fb8a1fd45/source/300x300bb.jpg"
	if j.Member.User.Avatar != "" {
		avatar = strings.Join([]string{"https://cdn.discordapp.com/avatars",
			j.Member.User.ID, "/",
			j.Member.User.Avatar, ".webp?size=256"}, "")
	}

	msg := NewEmbed().
		SetTitle("A user has joined the guild.").
		AddField("User Information", userinfo).
		SetThumbnail(avatar).
		SetColor(0x00ff2b).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(CHANJOINPART, msg)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("User join event was not logged. " + userinfo)
	}
}

func evtPart(s *discordgo.Session, j *discordgo.GuildMemberRemove) {
	logger.Debug().Msg("New User Part Event.")
	// build user info string
	userinfo := strings.Join([]string{"Username: ",
		j.Member.User.Username, "#",
		j.Member.User.Discriminator,
		"\nID: ", j.Member.User.ID,
		"\nTimestamp: ", time.Now().UTC().Format(time.RFC1123)}, "")

	// get avatar and set a default one if it doesn't exist
	avatar := "http://is1.mzstatic.com/image/thumb/Purple117/v4/a1/d8/3a/a1d83a42-e84e-5965-c006-610fb8a1fd45/source/300x300bb.jpg"
	if j.Member.User.Avatar != "" {
		avatar = strings.Join([]string{"https://cdn.discordapp.com/avatars",
			j.Member.User.ID, "/",
			j.Member.User.Avatar, ".webp?size=256"}, "")
	}

	msg := NewEmbed().
		SetTitle("A user has parted from the guild.").
		AddField("User Information", userinfo).
		SetThumbnail(avatar).
		SetColor(0xff0000).MessageEmbed

	_, err := s.ChannelMessageSendEmbed(CHANJOINPART, msg)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("User part event was not logged. " + userinfo)
	}
}

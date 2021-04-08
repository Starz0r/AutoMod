package main

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/spidernest-go/logger"
)

func cmdSilence(s *discordgo.Session, j *discordgo.MessageCreate) {
	// authz check
	if j.Author.ID == "749152747735089162" {
		return
	}

	for i, role := range j.Member.Roles {
		if role == "729002351398092904" {
			break
		}

		if i >= len(j.Member.Roles) {
			logger.Info().Msg("Unauthorized user attempted to use the silence command.")
			return
		}
	}

	// command check
	args := strings.Split(j.Content, " ")
	if args[0] != "!silence" {
		return
	}

	// get the duration deadline
	duration, err := ParseDuration(args[2])
	if err != nil {
		logger.Error().Err(err).Msg("Duration argument was malformed and could not be parsed.")
		s.ChannelMessageSend(j.ChannelID, "The specified duration could not be processed, please respecify.")
		return
	}
	durdate := time.Now().Add(duration)
	durstr := humanize.Time(durdate)

	// silence targetted users
	// TODO: range over this instead of taking the 0th value
	target := j.Mentions[0]
	if target == nil {
		return
	}

	err = s.GuildMemberRoleAdd(GUILDIDENT, target.ID, ROLESILENCED)
	if err != nil {
		logger.Error().Err(err).Msg("User could not be silenced.")
		s.ChannelMessageSend(j.ChannelID, "User was unable to be silenced, an internal error has occurred. Please check the logs for more context!")
		return
	}

	// create a task and write it to the current datastore
	task, err := writeTask(
		&Task{
			ID:          0,
			Event:       "unsilence",
			RequestedBy: j.Author.ID,
			Affects:     target.ID,
			Deadline:    durdate,
		},
	)
	if err != nil {
		s.GuildMemberRoleRemove(GUILDIDENT, target.ID, ROLESILENCED)
		logger.Error().Err(err).Msg("Task was not written to datastore, undoing silence.")
		s.ChannelMessageSend(j.ChannelID, "User was unable to be silenced due to a database error. Please check the logs for more context!")
		return
	}

	delegateTask(task)
	s.ChannelMessageSend(j.ChannelID, "User was silenced successfully! They will be unsilenced at "+durdate.Format("Monday, January 2, 2006 15:04:05 MST"))

	// send the targets a message
	reason := *new(string)
	if len(args) < 4 {
		reason = "No Reason Given."
	} else {
		var combine = func(strs []string) string {
			longstr := ""
			for _, str := range strs {
				longstr += str
			}
			return longstr
		}
		reason = combine(args[3:])
	}
	dm, err := s.UserChannelCreate(target.ID)
	if err != nil {
		logger.Warn().Err(err).Msg("User was not notified of silence.")
		return
	}
	_, err = s.ChannelMessageSend(dm.ID, "You've been silenced in the `I Wanna Community` guild server until "+durdate.Format("Monday, January 2, 2006 15:04:05 MST")+".\n\n**Duration:** "+durstr+".\n**Reason:** `"+reason+"`.")
	if err != nil {
		logger.Warn().Err(err).Msg("User was not notified of silence.")
		return
	}
}

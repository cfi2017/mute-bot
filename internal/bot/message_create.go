package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var muted = false

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "!mute" {
		return
	}

	muted = !muted
	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Muted: %v", muted))
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
		log.Println(err)
		return
	}
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
		log.Println(err)
		return
	}

	log.Printf("%d voice states", len(guild.VoiceStates))
	for _, state := range guild.VoiceStates {
		if state.Mute == muted {
			continue
		}

		err := s.GuildMemberMute(m.GuildID, state.UserID, muted)
		if err != nil {
			log.Println(err)
		}
	}
}

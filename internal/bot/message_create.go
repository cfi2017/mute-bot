package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

var Muted = false

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "!mute" {
		return
	}

	Muted = !Muted
	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Muted: %v", Muted))
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
		log.Println(err)
		return
	}
	guild, err := s.Guild(m.GuildID)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
		log.Println(err)
		return
	}

	for _, state := range guild.VoiceStates {
		if state.Mute != Muted {
			err := s.GuildMemberMute(m.GuildID, state.UserID, Muted)
			if err != nil {
				log.Println(err)
			}
		}
	}

}

package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"

	"log"
)

var muted = false

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "!mute" {
		return
	}

	muted = !muted
	if viper.GetBool("verbose") {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Muted: %v", muted))
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
			log.Println(err)
			return
		}
	} else {
		_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Unknown error.")
		log.Println(err)
		return
	}

	log.Printf("%d voice states", len(guild.VoiceStates))
	muteAll(s, guild, muted)
}

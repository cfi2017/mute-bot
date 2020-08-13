package bot

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func VoiceJoin(s *discordgo.Session, u *discordgo.VoiceStateUpdate) {
	if u.ChannelID == "" {
		return
	}

	if u.Mute == muted {
		return
	}

	err := s.GuildMemberMute(u.GuildID, u.UserID, muted)
	if err != nil {
		log.Println(err)
	}

}

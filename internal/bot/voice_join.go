package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var oldState = make(map[string]*discordgo.VoiceStateUpdate)

func VoiceJoin(s *discordgo.Session, u *discordgo.VoiceStateUpdate) {
	if u.ChannelID == "" {
		// remove state on leave
		if _, ok := oldState[u.UserID]; ok {
			oldState[u.UserID] = nil
		}
		return
	}

	// state exists
	if _, ok := oldState[u.UserID]; ok {
		oldState[u.UserID] = u
		return
	}

	oldState[u.UserID] = u

	if u.Mute == muted {
		return
	}

	err := s.GuildMemberMute(u.GuildID, u.UserID, muted)
	if err != nil {
		log.Println(err)
	}

}

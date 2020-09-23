package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func muteAll(s *discordgo.Session, guild *discordgo.Guild, m bool) {
	for _, state := range guild.VoiceStates {
		if state.Mute == m {
			continue
		}

		go mute(s, guild, state, m)
	}
}

func mute(s *discordgo.Session, g *discordgo.Guild, state *discordgo.VoiceState, m bool) {
	err := s.GuildMemberMute(g.ID, state.UserID, m)
	if err != nil {
		log.Println(err)
	}
}

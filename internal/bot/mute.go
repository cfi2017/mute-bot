package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func muteAll(s *discordgo.Session, guild *discordgo.Guild, m bool) {
	log.Printf("setting mute state for all users in guild %s to %v", guild.ID, m)
	for _, state := range guild.VoiceStates {
		if state.Mute == m {
			continue
		}
		log.Printf("muting %s", state.UserID)
		go mute(s, guild, state, m)
	}
}

func mute(s *discordgo.Session, g *discordgo.Guild, state *discordgo.VoiceState, m bool) {
	err := s.GuildMemberMute(g.ID, state.UserID, m)
	if err != nil {
		log.Println(err)
	}
}

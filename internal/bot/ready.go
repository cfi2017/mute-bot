package bot

import "github.com/bwmarrin/discordgo"

func Ready(s *discordgo.Session, _ *discordgo.Ready) {

	// Set the playing status.
	_ = s.UpdateStatus(0, "!mute")
}

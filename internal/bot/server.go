package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/mute-bot/internal/model"
	"github.com/gin-gonic/gin"
)

func InitialiseServer(r *gin.Engine, s *discordgo.Session) {
	r.POST("/event", func(c *gin.Context) {
		req := model.Event{}
		_ = c.BindJSON(&req)

		switch req.Type {
		case model.GameStartedEventType:
			{
				// mute all
				g, err := s.State.Guild(req.GuildID)
				if err != nil {
					log.Println(err)
					_ = c.AbortWithError(500, err)
					return
				}
				muted = true
				muteAll(s, g, muted)
			}
		case model.GameEndedEventType:
			{
				// unmute all
				g, err := s.State.Guild(req.GuildID)
				if err != nil {
					log.Println(err)
					_ = c.AbortWithError(500, err)
					return
				}
				muted = false
				muteAll(s, g, muted)
			}
		case model.MeetingStartedEventType:
			{
				// unmute all
				g, err := s.State.Guild(req.GuildID)
				if err != nil {
					log.Println(err)
					_ = c.AbortWithError(500, err)
					return
				}
				muted = false
				muteAll(s, g, muted)
			}
		case model.MeetingEndedEventType:
			{
				// mute all
				g, err := s.State.Guild(req.GuildID)
				if err != nil {
					log.Println(err)
					_ = c.AbortWithError(500, err)
					return
				}
				muted = true
				muteAll(s, g, muted)
			}
		}
	})
}

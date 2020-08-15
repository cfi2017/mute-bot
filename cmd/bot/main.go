package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/mute-bot/internal/bot"
	"github.com/cfi2017/mute-bot/internal/util"
	"github.com/spf13/viper"
)

func main() {
	bot.InitialiseFlags()
	util.InitialiseConfig("mute-bot")

	token := viper.GetString("token")
	if token == "" {
		log.Fatalf("No token provided. Please run with -t <token>, a DISCORD_TOKEN env or a valid config file.")
	}

	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		log.Fatalln("Error creating Discord session: ", err)
	}

	dg.AddHandler(bot.Ready)
	dg.AddHandler(bot.MessageCreate)
	dg.AddHandler(bot.VoiceJoin)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}

	dg.State.TrackVoice = true

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()

}

package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/mute-bot/internal/bot"
	"github.com/cfi2017/mute-bot/internal/util"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot.InitialiseFlags()
	util.InitialiseConfig("mute-bot")

	token := viper.GetString("token")
	if token == "" {
		log.Fatalf("No token provided. Please run with -t <token>, a BOT_TOKEN env or a valid config file.")
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

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()

}

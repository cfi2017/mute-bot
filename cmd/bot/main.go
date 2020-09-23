package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cfi2017/mute-bot/internal/bot"
	"github.com/cfi2017/mute-bot/internal/util"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()
	bot.InitialiseServer(r, dg)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", "0.0.0.0", 8080),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// probably timeout exceeded
		log.Fatal("Server forced to shutdown:", err)
	}

	// Cleanly close down the Discord session.
	_ = dg.Close()

}

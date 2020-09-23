package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/cfi2017/mute-bot/internal/model"
	"github.com/cfi2017/mute-bot/internal/sniffer"
	"github.com/cfi2017/mute-bot/internal/util"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/viper"
)

func main() {
	// set up dependencies
	sniffer.InitialiseFlags()
	util.InitialiseConfig("sniffer")

	// start sniffing
	const (
		snapLength  = 1024
		promiscuous = false
		timeout     = 30 * time.Second
	)
	handle, err := pcap.OpenLive(viper.GetString("device"), snapLength, promiscuous, timeout)
	//handle, err := pcap.OpenOffline("among_us.pcapng")
	if err != nil {
		log.Fatal("error accessing device", err)
	}
	defer handle.Close()

	filter := fmt.Sprintf("udp and port %d", viper.GetInt("server_port"))
	log.Println(filter)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal("error setting display filter", err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	cooldown := time.Now().Add(-2 * time.Second)
	for packet := range packetSource.Packets() {
		// Process packet here
		// fmt.Println(packet)
		data := packet.Data()
		hexData := fmt.Sprintf("%x", data)
		meetingEndedRegexp := regexp.MustCompile("^(01).{6}(0005).{6}(80).{2}(0002)")
		meetingStartedRegexp := regexp.MustCompile("^(01).{6}(0005).*(401c460000401c46)")
		if time.Since(cooldown) >= 0*time.Second {
			if (bytes.Contains(data, []byte("EndGame")) &&
				bytes.Equal(data[3:6], []byte{0x12, 0x00, 0x05})) ||
				bytes.Equal(data[3:6], []byte{0x06, 0x00, 0x08}) {
				log.Println("game ended")
				// game ended
				report(viper.GetString("guild_id"), model.GameEndedEventType)
			} else if len(data) >= 200 && strings.Contains(hexData, "460000000001010101010101010101010101000000000") {
				// game started
				log.Println("game started")
				report(viper.GetString("guild_id"), model.GameStartedEventType)
			} else if len(data) <= 16 && len(data) >= 15 {
				if meetingEndedRegexp.MatchString(hexData) {
					// meeting ended
					log.Println("meeting ended")
					report(viper.GetString("guild_id"), model.MeetingEndedEventType)
				}
			} else if meetingStartedRegexp.MatchString(hexData) {
				// meeting started
				log.Println("meeting started")
				report(viper.GetString("guild_id"), model.MeetingStartedEventType)
			}
		}
	}

	// filter packets and output commands to channel

	// handle output commands (file, server?)
}

func report(guildID string, eventType model.EventType) {
	event := model.Event{
		Type:    eventType,
		GuildID: guildID,
	}
	bs, _ := json.Marshal(&event)
	_, _ = http.Post(viper.GetString("report_endpoint"), "application/json", bytes.NewReader(bs))
}

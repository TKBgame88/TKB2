package main

import (
	"fmt"
	constants "go_bot/src/const"
	"go_bot/src/utils"
	"log"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	dg, err := discordgo.New("Bot " + token)
	if err != nil {

		fmt.Println("error:start\n", err)
		return
	}

	//on message
	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error:wss\n", err)
		return
	}
	dg.UpdateGameStatus(1, "Test")
	fmt.Println("BOT Running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var utils utils.Utils

	if m.Author.Bot {
		return
	}
	nick := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nick = member.Nick
	}
	fmt.Println("< " + m.Content + " by " + nick)

	// 持ち越しTL変換のマッチ
	r := regexp.MustCompile(constants.CARRY_OVER_REGEX)
	if r.MatchString(m.Content) {
		msg := utils.Convert(m.Content)

		s.ChannelMessageSend(m.ChannelID, msg)
	}
}

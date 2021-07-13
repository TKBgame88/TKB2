package main

import (
	"fmt"
	constants "go_bot/src/const"
	"go_bot/src/utils"
	"log"
	"os"
	"os/signal"
	"strconv"
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

	// 持ち越しTL変換のマッチ
	if constants.CARRY_OVER_REGEX.MatchString(m.Content) {
		msg, sec := utils.Convert(m.Content)

		desc := "```\n\n" + strconv.Itoa(sec) + "秒の持ち越しTLだ!\n" + msg + "```"

		s.ChannelMessageSend(m.ChannelID, desc)
	}
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	constants "pricess_connect_lite_bot/src/const"
	"pricess_connect_lite_bot/src/utils"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	token := ""
	if os.Getenv("BOT_TOKEN") != "" {
		token = os.Getenv("BOT_TOKEN")
	}

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
		utils.Convert(m.Content, s, m)
	}
}

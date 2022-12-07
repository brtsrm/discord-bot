package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var token = "DISCORD-TOKEN"
var botChID = "BOT_TOKEN"
var Dg *discordgo.Session

func initSession() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("DC Oturumu acilmadi Err", err)
		l.Fatalf("DC Oturumu acilmadi Err", err)
	}
	Dg = dg
}

func ConnectToDc() {
	initSession()
	go ParseRSS()

	Dg.AddHandler(messageCreate)

	Dg.Identify.Intents = discordgo.IntentsGuildMessages

	fmt.Println("Bot calisiyor. Cikma icin CTRL+C basiniz.")
	l.Println("[INFO] Bot calisiyor. Cikma icin CTRL+C basiniz.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	Dg.Close()
	l.Println("[INFO] Bot durduldu")
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Yardım menüsüne hoş geldiniz. <URL>")
}
func handleAddBlog(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Yardım menüsüne hoş geldiniz. <URL>")
	splits := strings.Split(m.Content, " ")
	url := splits[2]
	if len(splits) != 3 {
		s.ChannelMessageSend(m.ChannelID, "Lütfen add_blog komutunu yugun formatta kulanlnı: !aaa add_blog")
		l.Printf("Rss parserr %s kullanicisi tarafinda ynalis komut girildi", m.Author)
	} else {
		file, err := os.OpenFile("blog_wishlist.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		if !strings.Contains(readFile("blog_wishlist.txt"), url) {
			if _, err := file.WriteString(url + "\n"); err != nil {
				l.Fatal(err)
			} else {
				s.ChannelMessageSend(m.ChannelID, "Blog Listeye eklendi")
				l.Printf("Listeye eklendi", m.Author, url)
			}
		} else {
			s.ChannelMessageSend(m.ChannelID, "Blog url zaten listede blulunıyor")
			l.Printf("%s")
		}

	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!aaa" {
		s.ChannelMessageSend(m.ChannelID, "Selam bebek naber")
		l.Printf("[INFO] %s kullanici beni cagridi", m.Author)
	}
	if len(strings.Split(m.Content, " ")) > 1 {
		if strings.Split(m.Content, " ")[0] == "!aaa" {

			if strings.Split(m.Content, " ")[1] == "help" {
				handleHelp(s, m)
			}

			if strings.Split(m.Content, " ")[1] == "add_blog" {
				handleAddBlog(s, m)
			}

		}

	}
}

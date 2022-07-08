package main

// https://api.betterttv.net/3/emotes/shared/top?offset=740&limit=50

// [{"id":"54fa8f1401e468494b85b537","code":":tf:","imageType":"png","userId":"5561169bd6b9d206222a8c19"},{"id":"54fa8fce01e468494b85b53c","code":"CiGrip","imageType":"png","userId":"5561169bd6b9d206222a8c19"},
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Emote struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
}

func main() {
	godotenv.Load()

	token := os.Getenv("DISCORD_BOT_TOKEN")

	fmt.Println("Getting emote data")

	client := &http.Client{}

	emotesByCode := make(map[string]Emote)
	for _, emote := range getGlobalEmotes(client) {
		emotesByCode[emote.Code] = emote
	}

	for i := 0; i < 25; i++ {
		for _, emote := range getTrendingEmotes(client, i*100) {
			if _, ok := emotesByCode[emote.Code]; !ok {
				emotesByCode[emote.Code] = emote
			}
		}
	}

	fmt.Println("Starting bot")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { onMessageCreate(emotesByCode, s, m) })
	discord.Identify.Intents = discordgo.IntentsGuildMessages

	// Setup connection
	err = discord.Open()
	if err != nil {
		panic(err)
	}
	defer discord.Close()
	defer fmt.Println("Closing bot")

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func getGlobalEmotes(client *http.Client) []Emote {
	url := "https://api.betterttv.net/3/cached/emotes/global"

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result := []Emote{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

type TrendingEmote struct {
	Emote Emote `json:"emote"`
}

func getTrendingEmotes(client *http.Client, offset int) []Emote {
	url := "https://api.betterttv.net/3/emotes/shared/trending?limit=100&offset=" + fmt.Sprint(offset)

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result := []TrendingEmote{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	emotes := []Emote{}
	for _, trendingEmote := range result {
		emotes = append(emotes, trendingEmote.Emote)
	}

	return emotes
}

func onMessageCreate(emotesByCode map[string]Emote, s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if emote, ok := emotesByCode[m.Content]; ok {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+":")
		s.ChannelMessageSend(m.ChannelID, "https://cdn.betterttv.net/emote/"+emote.ID+"/2x."+emote.ImageType)
	}
}

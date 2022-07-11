package main

// https://api.betterttv.net/3/emotes/shared/top?offset=740&limit=50

// [{"id":"54fa8f1401e468494b85b537","code":":tf:","imageType":"png","userId":"5561169bd6b9d206222a8c19"},{"id":"54fa8fce01e468494b85b53c","code":"CiGrip","imageType":"png","userId":"5561169bd6b9d206222a8c19"},
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Emote struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
	Order     int
}

func getDefaultEmotes() map[string]Emote {
	client := &http.Client{}

	emotesChannel := make(chan []Emote, 26)
	wg := new(sync.WaitGroup)

	emotesByCode := make(map[string]Emote)
	wg.Add(1)
	go func() {
		getGlobalEmotes(client, emotesChannel)
		wg.Done()
	}()

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			getTrendingEmotes(client, i*100, emotesChannel)
			wg.Done()
		}()
	}

	wg.Wait()
	close(emotesChannel)

	for result := range emotesChannel {
		for _, emote := range result {
			existingEmote, existsEmote := emotesByCode[emote.Code]
			if !existsEmote || existingEmote.Order < emote.Order {
				emotesByCode[emote.Code] = emote
			}
		}
	}

	return emotesByCode
}

func main() {
	godotenv.Load()
	token := os.Getenv("DISCORD_BOT_TOKEN")

	fmt.Println("Getting emote data")
	emoteByCode := getDefaultEmotes()

	fmt.Println("Starting bot")
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) { onMessageCreate(emoteByCode, s, m) })
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

func getGlobalEmotes(client *http.Client, c chan []Emote) {
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

	for _, emote := range result {
		emote.Order = -1
	}

	c <- result
}

type TrendingEmote struct {
	Emote Emote `json:"emote"`
}

func getTrendingEmotes(client *http.Client, offset int, c chan []Emote) {
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
		trendingEmote.Emote.Order = offset
		emotes = append(emotes, trendingEmote.Emote)
	}

	c <- emotes
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

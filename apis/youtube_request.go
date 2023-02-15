package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type ThumbnailDetail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnail struct {
	Default  ThumbnailDetail `json:"default"`
	Medium   ThumbnailDetail `json:"medium"`
	High     ThumbnailDetail `json:"high"`
	Standard ThumbnailDetail `json:"standard"`
	Maxres   ThumbnailDetail `json:"maxres"`
}

type Sni struct {
	PublishedAt          string    `json:"publishedAt"`
	ChannelId            string    `json:"channelId"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Thumbnails           Thumbnail `json:"thumbnails"`
	ChannelTitle         string    `json:"channelTitle"`
	DefaultAudioLanguage string    `json:"defaultAudioLanguage"`
}

type Statistic struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

type Item struct {
	Id         string    `json:"id"`
	Snippet    Sni       `json:"snippet"`
	Statistics Statistic `json:"statistics"`
}

type YoutubeVideoData struct {
	Items []*Item `json:"items"`
}

func (me *YoutubeVideoData) String() string {
	return fmt.Sprintf("チャンネル名: %s\nチャンネルID: %s\n動画タイトル: %s\n再生数: %s", me.Items[0].Snippet.ChannelTitle, me.Items[0].Snippet.ChannelId, me.Items[0].Snippet.Title, me.Items[0].Statistics.ViewCount)
}

var (
	Log      = log.New(os.Stderr, "", 0)
	errorLog = log.New(os.Stderr, "[Error]", 0)
)

func YoutubeRequest(movieId string) {

	url := []string{"https://www.googleapis.com/youtube/v3/videos?id=", movieId, "&key=", os.Getenv("YOUTUBE_API"), "&part=snippet,statistics"}

	res, err := http.Get(strings.Join(url, ""))
	if err != nil {
		errorLog.Println(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorLog.Printf("http status code: %d", res.StatusCode)
		return
	}

	var (
		movie   = &YoutubeVideoData{}
		decoder = json.NewDecoder(res.Body)
	)

	err = decoder.Decode(movie)
	if err != nil {
		errorLog.Println(err)
		return
	}

	Log.Printf("status: %d\n%s\n", res.StatusCode, movie)

}

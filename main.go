package main

import (
	"go-api/apis"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Log      = log.New(os.Stderr, "", 0)
	errorLog = log.New(os.Stderr, "[Error]", 0)
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		errorLog.Printf("環境変数を読み込み出来ませんでした: %v", err)
	}

	//動画詳細の出力
	youtube_movieId := "RbDcICCxm5k"
	apis.YoutubeRequest(youtube_movieId)
}

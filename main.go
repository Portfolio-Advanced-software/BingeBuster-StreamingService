package main

import "github.com/Portfolio-Advanced-software/BingeBuster-StreamingService/messaging"

func main() {
	messaging.ProduceMessage("get title for movie 3", "streaming_movie")
	messaging.ConsumeMessage("streaming_movie")
}

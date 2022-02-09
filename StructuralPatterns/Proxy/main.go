package main

import "fmt"

func main() {
	// Get Third Party lib service
	thirdParty := NewThirdPartyYoutubeLib()

	// Get proxy
	proxy := NewYoutubeProxy(thirdParty)

	// call service with proxy
	video := proxy.getVideoInfo("1")
	fmt.Printf("Video 1 info is: %+v\n", video)

	videos := proxy.listVideos()
	fmt.Printf("Videos:\n %+v\n", videos)

	videoDowloaded := proxy.downloadVideo("1")
	fmt.Printf("Video 1 downloaded: %+v\n", videoDowloaded)
}

package main

import "fmt"

type YoutubeProxy struct {
	service IThirdPartyYoutubeLib
}

func (proxy *YoutubeProxy) getVideoInfo(id string) VideoInfo {
	fmt.Println("========== PROXY ==========")
	fmt.Printf("Getting Video %s Info...\n", id)
	video := proxy.service.getVideoInfo(id)
	fmt.Println("Success!")
	fmt.Println("========== PROXY ==========")
	return video
}

func (proxy *YoutubeProxy) listVideos() []VideoInfo {
	fmt.Println("========== PROXY ==========")
	fmt.Println("Listing videos...")
	videos := proxy.service.listVideos()
	fmt.Println("Success!")
	fmt.Println("========== PROXY ==========")
	return videos
}

func (proxy *YoutubeProxy) downloadVideo(id string) VideoDownloaded {
	fmt.Println("========== PROXY ==========")
	fmt.Printf("Downloading video %s...\n", id)
	video := proxy.service.downloadVideo(id)
	fmt.Println("Video downloaded!")
	fmt.Println("========== PROXY ==========")
	return video
}

func NewYoutubeProxy(service IThirdPartyYoutubeLib) *YoutubeProxy {
	return &YoutubeProxy{
		service: service,
	}
}

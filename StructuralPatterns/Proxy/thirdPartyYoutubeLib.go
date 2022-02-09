package main

import "encoding/json"

type VideoInfo struct {
	id             string
	title          string
	duration       float64
	visualizations uint
}

type VideoDownloaded struct {
	encoder string
	data    []byte
}

type ThirdPartyYoutubeLib struct{}

func (lib *ThirdPartyYoutubeLib) getVideoInfo(id string) VideoInfo {
	return VideoInfo{
		id:             id,
		title:          "Video 1",
		duration:       1000.2,
		visualizations: 1000,
	}
}

func (lib *ThirdPartyYoutubeLib) listVideos() []VideoInfo {
	videos := []VideoInfo{
		{
			id:             "1",
			title:          "Video 1",
			duration:       1000.2,
			visualizations: 1000,
		},
		{
			id:             "2",
			title:          "Video 2",
			duration:       2000.2,
			visualizations: 2000,
		},
		{
			id:             "3",
			title:          "Video 3",
			duration:       3000.2,
			visualizations: 3000,
		},
	}

	return videos

}

func (lib *ThirdPartyYoutubeLib) downloadVideo(id string) VideoDownloaded {

	video := VideoInfo{
		id:             "1",
		title:          "Video 1",
		duration:       1000.2,
		visualizations: 1000,
	}

	videoStr, _ := json.Marshal(video)

	return VideoDownloaded{
		encoder: "mmpeg",
		data:    []byte(videoStr),
	}

}

func NewThirdPartyYoutubeLib() *ThirdPartyYoutubeLib {
	return &ThirdPartyYoutubeLib{}
}

package main

type IThirdPartyYoutubeLib interface {
	getVideoInfo(id string) VideoInfo
	listVideos() []VideoInfo
	downloadVideo(id string) VideoDownloaded
}

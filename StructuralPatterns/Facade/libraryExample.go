package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Video struct {
	name string
	size float64
}

type ILibraryExample interface {
	EncodeVideo(Video) string
	DecodeVideo(string) Video
	PublishVideo(Video) bool
	UnpublishVideo(Video) bool
	GetVideo(string) Video
	DeleteVideo(string) bool
}

type LibraryExample struct{}

func (library *LibraryExample) EncodeVideo(video Video) string {
	fmt.Println("Encoding video...")
	return video.name + "/" + fmt.Sprintf("%f", video.size)
}

func (library *LibraryExample) DecodeVideo(video string) Video {
	fmt.Println("Decoding video...")
	videoString := strings.Split(video, "/")
	name := videoString[0]
	size, err := strconv.ParseFloat(videoString[1], 64)
	if err != nil {
		log.Fatalf("error trying to parseFloat video string %s", err)
	}

	return Video{
		name: name,
		size: size,
	}
}

func (library *LibraryExample) PublishVideo(video Video) bool {
	fmt.Println("Publishing video...")
	fmt.Printf("Published! %+v\n", video)
	return true
}

func (library *LibraryExample) UnpublishVideo(video Video) bool {
	fmt.Println("Unpublishing video...")
	fmt.Printf("Unpublished! %+v\n", video)
	return true
}

func (library *LibraryExample) GetVideo(id string) Video {
	fmt.Println("Getting video...")
	return Video{
		name: "Video #" + id,
		size: 240.5,
	}
}

func (library *LibraryExample) DeleteVideo(id string) bool {
	fmt.Println("Deleting video...")
	fmt.Printf("Deleted #%s!\n", id)
	return true
}

func NewLibraryExample() ILibraryExample {
	return &LibraryExample{}
}

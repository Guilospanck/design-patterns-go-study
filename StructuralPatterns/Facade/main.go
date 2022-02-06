package main

import "fmt"

func main() {
	// Whole library
	library := NewLibraryExample()

	// facade of library using only the methods needed
	encodeDecodeFacade := NewEncodeDecodeFacade(library)

	// new video
	video := Video{
		name: "Video1",
		size: 222,
	}
	fmt.Printf("Original video: %+v\n", video)

	// encode video
	encoded := encodeDecodeFacade.EncodeVideo(video)
	fmt.Printf("Encoded: %s\n", encoded)

	// decode video
	decoded := encodeDecodeFacade.DecodeVideo(encoded)
	fmt.Printf("Decoded: %+v\n", decoded)

}

package main

type EncodeDecodeFacade struct {
	library ILibraryExample
}

func (edf *EncodeDecodeFacade) EncodeVideo(video Video) string {
	return edf.library.EncodeVideo(video)
}

func (edf *EncodeDecodeFacade) DecodeVideo(video string) Video {
	return edf.library.DecodeVideo(video)
}

func NewEncodeDecodeFacade(library ILibraryExample) *EncodeDecodeFacade {
	return &EncodeDecodeFacade{
		library: library,
	}
}

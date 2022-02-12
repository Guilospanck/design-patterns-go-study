package main

import "fmt"

type Data struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func (d *Data) Marshal() string {
	return d.Name + "<>" + fmt.Sprintf("%d", d.Size)
}

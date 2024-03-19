package main

import (
	"image"
	"sync"
)

var loadOnce sync.Once

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是线程安全的
func Icon(name string) image.Image {
	loadOnce.Do(loadIcons)
	return icons[name]
}

func loadIcon(name string) image.Image {
	// do something...
	return nil
}

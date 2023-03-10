package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"

	"github.com/sqweek/dialog"
)

// Bridge between frontend and backend
type Bridge struct {
	ctx context.Context
}

// NewBridge initializes a new runtime instance
func NewBridge() *Bridge {
	return &Bridge{}
}

// Startup is called at application startup
func (b *Bridge) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// ExportProject chooses a file then saves the current project
func (b *Bridge) ExportProject(data string) {
	filename, err := dialog.File().SetStartFile("project.json").Title("Export Project").Filter("Project (*.json)", "json").Save()
	if err != nil {
		return
	}
	_ = os.WriteFile(filename, []byte(data), 0644)
}

// ImportProject loads and returns an existing project from a file
func (b *Bridge) ImportProject() string {
	filename, err := dialog.File().Title("Import Project").Filter("Project (*.json)", "json").Load()
	if err != nil {
		return err.Error()
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return err.Error()
	}
	return string(data)
}

// LoadImage loads an image and converts it to base64
func (b *Bridge) LoadImage() string {
	path, err := dialog.File().Title("Load Image").Filter("Image Files", "png", "jpeg", "jpg", "gif", "webp").Load()
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	var base64Encoding string
	mimeType := http.DetectContentType(data)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}
	base64Encoding += base64.StdEncoding.EncodeToString(data)

	return base64Encoding
}

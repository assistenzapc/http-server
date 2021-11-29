package main

import (
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/net/html/charset"
)

// extensionList holds a key-value store with the most common
// file extensions and their corresponding associations.
// There's also file names which are common across platform
// to identify certain file types.
var extensionList = map[string]string{
	// File extensions from https://www.computerhope.com/issues/ch001789.htm
	".pdf":       "PDF",
}

// fileNameList is a map from filename to the file type
// associated with it
var fileNameList = map[string]string{
	".dockerignore":     "Docker Ignore",
	".gitattribute":     "Git attribute",
	".bash_profile":     "Bash Profile",
	".profile":          "Bash Profile",
	".bash_history":     "Bash History",
	".bash_logout":      "Bash Logout",
	".bashrc":           "Bash RC",
	".gemrc":            "Ruby Gem Config",
	".minttyrc":         "MinTTY Config",
	".npmjs":            "NPM Config",
	".yarnrc":           "Yarn Config",
	".vim":              "Vim Config",
	".vimrc":            "Vim Config",
	".vimtags":          "Vim Tags Config",
	".babelrc":          "Babel Config",
	".wget-hsts":        "Wget HSTS Config",
	".tmux.conf":        "Tmux Config",
	"webpack.config.js": "Webpack Configuration",
	"Dockerfile":        "Dockerfile",
	"LICENSE":           "License",
	"CONTRIBUTE":        "Contributor README",
	"README":            "README",
	"README.md":         "README Markdown",
	"README.markdown":   "README Markdown",
	"Makefile":          "GNU Make",
	"Makefile.inc":      "GNU Make include",
	"Gemfile":           "Ruby Gem",
	"Rakefile":          "Ruby Rake",
	"config.ru":         "Ruby Config",
	"Vagrant":           "Vagrant VM",
	"config":            "Config",
	"go.mod":            "Go Module File",
}

// detectByName tries to find the filetype based on the
// file name using the map above
func detectByName(name string) string {
	// Get the content type based off the full file name
	if content, found := fileNameList[name]; found {
		return content
	}

	// Get the content type based off the file extension
	if content, found := extensionList[filepath.Ext(name)]; found {
		return content
	}

	// Get the content type based off the file name without extension
	if content, found := fileNameList[strings.TrimSuffix(name, filepath.Ext(name))]; found {
		return content
	}

	return ""
}

var overrideCTypeExtension = map[string]string{}

// generateContentTypeCharset tries to find the filetype based on the
// file content using the map above
func generateContentTypeCharset(name string, content []byte) string {
	if s, found := overrideCTypeExtension[name]; found {
		return s
	}

	s := http.DetectContentType(content)

	if _, name, certain := charset.DetermineEncoding(content, s); certain && !strings.Contains(s, ";") {
		return s + "; charset=" + name
	}

	return s
}

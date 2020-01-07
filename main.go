package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Basic example of how to commit changes to the current branch to an existing
// repository.
func main() {
	CheckArgs("<directory>")
	directory := os.Args[1]


}
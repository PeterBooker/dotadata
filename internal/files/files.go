// +build dev

package files

import (
	"net/http"
	"path/filepath"
)

// Assets contains static assets.
var Assets http.FileSystem = http.Dir(filepath.Join("web", "build"))

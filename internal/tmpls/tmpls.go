// +build dev

package tmpls

import (
	"net/http"
	"path/filepath"
)

// Templates contains HTML templates.
var Assets http.FileSystem = http.Dir(filepath.Join("web", "templates"))

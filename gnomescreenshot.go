package gnomescreenshot

import ( 
	"bytes"
	"log"
	"os/exec"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"os"
)

var dir string

// ErrNoDir is returned when trying to access the screenshot directory
var ErrNoDir = errors.New("directory not set for output")

// SetDirectory name must be called before MakeDirectory to allow us to 
// create a directory for screenshots...
func SetDirectoryName(name string) {
	dir = name
}

// MakeDirectory allows us to generate a directory to store screenshots
func MakeDirectory(dir string) error {
	if dir == "" {
		return ErrNoDir
	}	

	return os.Mkdir(dir, 0770)
}

// getUUID returns a UUID string to be used for naming
func getUUID() string {
	uid, _ := uuid.NewRandom()
	return uid.String()
}

// GrabScreenshot returns a link to a screenshot for retrieval on localhost
// and and errors encountered along the way with the calling application
func GrabScreenshot(link string) (string, string, error) {

	if dir != "" {
		err := MakeDirectory(dir)
		if err.Error() != "mkdir " + dir + ": file exists" {
			return "", "", err
		}
	}	

	command := "gnome-web-photo" 
	thumbnail := "--mode=thumbnail" 
	size := "-s" 
	sizeVal := "256" 
	hyperlink := link
	filename := dir + "/" + getUUID() + ".png"

	args := []string{thumbnail, size, sizeVal, hyperlink, filename}
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	return filename, out.String(), err
}

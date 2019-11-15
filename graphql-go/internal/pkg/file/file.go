package file

import "io/ioutil"

// Read Reads a file and returns a string from the file
func Read(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

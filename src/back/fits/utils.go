package fits

import (
	"os"

	"github.com/siravan/fits"
)

func GetFitsUnitsFromFile(filePath string) ([]*fits.Unit, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	f, err := fits.Open(file)
	if err != nil {
		return nil, err
	}
	return f, nil
}

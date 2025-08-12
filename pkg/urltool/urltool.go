package urltool

import (
	"net/url"
	"path"
)

func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if err != nil {
		return "", err
	}

	basePath := path.Base(myUrl.Path)

	return basePath, nil
}

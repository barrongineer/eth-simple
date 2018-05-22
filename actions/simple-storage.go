package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

const (
	SimpleStoragePage  = "simple-storage.html"
	StoredDataModelKey = "storedData"
)

var storedData int

func SimpleStorageGetHandler(c buffalo.Context) error {
	c.Set(StoredDataModelKey, storedData)
	return c.Render(http.StatusOK, r.HTML(SimpleStoragePage))
}

func SimpleStoragePutHandler(c buffalo.Context) error {
	var form StoredDataRequest
	err := c.Bind(&form)
	if err != nil {
		return err
	}

	storedData = form.StoredData

	c.Set(StoredDataModelKey, storedData)
	return c.Render(http.StatusOK, r.HTML(SimpleStoragePage))
}

type StoredDataRequest struct {
	StoredData int `form:"storedData"`
}

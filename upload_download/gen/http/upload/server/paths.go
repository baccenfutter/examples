// Code generated by goa v3.0.9, DO NOT EDIT.
//
// HTTP request path constructors for the upload service.
//
// Command:
// $ goa gen goa.design/examples/upload/design

package server

import (
	"fmt"
)

// UploadUploadPath returns the URL path to the upload service upload HTTP endpoint.
func UploadUploadPath() string {
	return "/"
}

// DownloadUploadPath returns the URL path to the upload service download HTTP endpoint.
func DownloadUploadPath(id string) string {
	return fmt.Sprintf("/%v", id)
}

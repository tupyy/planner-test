// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

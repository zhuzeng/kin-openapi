package openapi3

import (
  "path/filepath"
  "testing"

  "github.com/stretchr/testify/require"
)

func TestIssue220(t *testing.T) {
  // on Linux e.g. /dev/my-openapi.json
  // on Windows e.g. C:/dev/my-openapi.json
  // openapi := filepath.FromSlash("C:/dev/my-openapi.json")
  specPath := filepath.FromSlash("testdata/my-openapi.json")
  loader := NewSwaggerLoader()
  loader.IsExternalRefsAllowed = true
  doc, err := loader.LoadSwaggerFromFile(specPath)
  require.NoError(t, err)
  err = doc.Validate(loader.Context)
  require.NoError(t, err)
}

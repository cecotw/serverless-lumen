//go:generate go-bindata -ignore=.(go|DS_Store) -pkg=schema ../...

package schema

import (
	"bytes"

	"github.com/flytedesk/foundation/services/graphql-go/internal/app/graphql/todo"
)

// QueryResolver : Root query resolver
type QueryResolver struct {
	*todo.Resolver
}

// MergeSchema merges the schema files to one string
func MergeSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}

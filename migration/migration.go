package migration

import "embed"

//go:embed postgres/*.sql
var MigrationsFS embed.FS

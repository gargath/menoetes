package migration

const Latest = 2018110601

func MigrationsFor(version int64) (stmts []string) {
	if version == 0 {
		stmts = migration_0
	}
	return stmts
}

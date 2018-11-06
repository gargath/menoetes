package migration

var migration_0 = []string{
	`CREATE TABLE modules (
  id varchar(255) PRIMARY KEY,
  owner varchar(100),
  namespace varchar(100) NOT NULL,
  name varchar(100) NOT NULL,
  version varchar(20) NOT NULL,
  provider varchar(50) NOT NULL,
  description text,
  source varchar(255),
  published_ad timestamp NOT NULL,
  downloads integer default 0,
  verified boolean default false
)`,
	`CREATE INDEX idx_module_name on modules(
  name
)`,
	`CREATE UNIQUE INDEX ixd_module_name_namespace on modules (
  name,
  namespace
)`,
	`UPDATE schemaversion SET version = 2018110601`,
}

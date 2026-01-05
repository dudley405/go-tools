module app/clients

go 1.25.5

replace app/types => ../types

replace app/clients => ../clients

replace app/logging => ../logging

require (
	app/logging v0.0.0-00010101000000-000000000000
	app/types v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.1
)

require go.uber.org/multierr v1.10.0 // indirect

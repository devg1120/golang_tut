module okiraku

go 1.23.1

replace foo => ./sample/foo

replace bar => ./sample/bar

require (
	bar v0.0.0-00010101000000-000000000000 // indirect
	foo v0.0.0-00010101000000-000000000000 // indirect
)

module meb.com/hello

go 1.18

require (
	golang.org/x/example v0.0.0-20220412213650-2e68773dfca0
	meb.com/greet v0.0.0-00010101000000-000000000000
)

replace meb.com/greet => ../greeting

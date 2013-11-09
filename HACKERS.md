to run with gdb

cd singularity
go build -o singularity *.go
gdb --args singularity server

gdb> run

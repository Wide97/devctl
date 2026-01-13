devctl

Small Go CLI used to run a few development utilities:
- sys info: fetches OS/arch/runtime info from a remote service
- ping: checks service reachability
- file exists: checks a remote file via API
- file ls: lists local directory contents
- calc: basic calculator (add/sub/mul/div)

Requirements
- Go 1.22+

Quick start
1) Run the mock server
   make mock

2) Set the base URL (optional if using the default in Makefile)
   set DEVCTL_BASE_URL=http://localhost:8080

3) Run a command
   make sys
   make ping
   make file-exists FILE_PATH=./README.md
   make file-ls DIR=.
   make calc CALC_OP=add X=2 Y=3

Environment
- DEVCTL_BASE_URL: base URL for remote service (used by sys/ping/file exists)

CLI usage
devctl sys info
devctl ping
devctl file exists <path>
devctl file ls [dir]
devctl calc <add|sub|mul|div> <val1> <val2>

Testing
make test

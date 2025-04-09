$ bat launch_config_file.sh
───────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────
│ File: launch_config_file.sh
───────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────
1   │ #!/bin/bash
2   │ # no arguments, no environment variables, use config file
3   │ # expected result:
4   │ # Host:  host_from_file
5   │ # Port:  8081
6   │ # Debug: true
7   │ go run config_loader.go
───────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────
$ ./launch_config_file.sh
Host:  host_from_file
Port:  8081
Debug: true

$ bat launch_env.sh
───────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────
│ File: launch_env.sh
───────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────
1   │ #!/bin/bash
2   │ # no arguments, use environment variables, ignore config file
3   │ # expected result:
4   │ # Host:  host_from_env
5   │ # Port:  8082
6   │ # Debug: false
7   │ APP_HOST="host_from_env" APP_PORT=8082 APP_DEBUG=false \
8   │     go run config_loader.go
───────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────
$ ./launch_env.sh
Host:  host_from_env
Port:  8082
Debug: false

$ bat launch_cli.sh
───────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────
│ File: launch_cli.sh
───────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────
1   │ #!/bin/bash
2   │ # use arguments, ignore environment variables, ignore config file
3   │ # expected result:
4   │ # Host:  host_from_cli
5   │ # Port:  8083
6   │ # Debug: true
7   │ APP_HOST="host_from_env" APP_PORT=8082 APP_DEBUG=false \
8   │     go run config_loader.go \
9   │     --host="host_from_cli" \
10   │     --port=8083 \
11   │     --debug=true
───────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────
$ ./launch_cli.sh
Host:  host_from_cli
Port:  8083
Debug: true

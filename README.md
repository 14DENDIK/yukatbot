# Telgram Bot on Golang
## Description
Telegram Bot written in Golang with standard libraries.<br>
Architecture can be used for any monolit server.<br>
### Run
1. For running the code first create ` configs/yukat.toml ` with all<br>
the neccessary fields. Fields can be found in ` internal/yukat/config/config.go ` file.
2. Create PostgresDatabase(with any name) and run migrations with `make migrations` 
3. After which run ` make ` that will create binary file

## Usage
In template branch the template version will be pushed.
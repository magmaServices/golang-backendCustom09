# `fesl-backend`

`backend` is an implementation of the `GameSpy` network adapted for BFHeroes use.

## Configuration

Below there is table with all enviroment variables which are used by the `openheroes/backend`.
You can refer to `config/config.go` file if you need more information about specific variable.

Instead of editing `~/bashrc` you may create a `.env` file and insert there configuration of your choice .
Optionally, you can provide a path to a `.env` file, which can be found in other directory by typing `./backend --config ../configs/.dev.env`.

| Name                  | Default value        |
|-----------------------|----------------------|
| `LOG_LEVEL`           | `INFO`               |

Note: It is okay to run on default configuration if you run server on your local PC for testing and development. But if you are thinking about exposing your server you probably should change some variables (i.e. `THEATER_ADDR`).

### Example `.env` file

```ini
DATABASE_NAME=open-heroes
DATABASE_HOST=192.168.33.10
DATABASE_PASSWORD=test
```

## Development

Before diving into the development you will probably need to download and [install Go](https://golang.org/dl/) programming lanugage compiler and set `GOPATH` env variable (`~/go` is used by default) - [see Linux installation manual](https://docs.minio.io/docs/how-to-install-golang).

### Installation in the `GOPATH`

To download code from the repository you could use a terminal (i.e. mingw/gitbash on Windows or preferable: built-in terminal emulator on Linux):

```bash
mkdir -p $GOPATH/src/gitlab.com/oiacow && \
cd $GOPATH/src/gitlab.com/oiacow && \
git clone https://gitlab.com/oiacow/fesl3.git && \
cd backend
```

### Prerequisites

`openheroes/backend`  currently only uses `MySQL` as a backing services. If you are on platform where `docker` is available, you may use following command to quickly download and start container with a MySQL database:

```bash
sudo docker-compose start
```

### Start

During development you might appreciate `Makefile` scripts, which in one simple command it will compile and run the compiled binary:

```bash
make run
```

Which is alias to:

```bash
go build -o main cmd/backend/main.go && sudo -H ./main
```

Unfortunately, Windows is not really great for running any console-based applications, but if you use `powershell` you might also appreciate following command:

```powershell
go build -o main.exe cmd/backend/main.go ; if ($?) { .\main.exe }
```

Or following if you are using custom `.env` file:

```powershell
go build -o main.exe cmd/backend/main.go ; if ($?) { .\main.exe --config .dev.env }
```

Note: PowerShell has one big advantage over other terminal in Windows - text coloring of logs.

### Dependencies

Currently golang dependencies are resolved thanks to [glide](https://github.com/Masterminds/glide).

Note: It is recommended to commit `vendor` directory to repository.

## Credits

This repository was forked from `github.com/HeroesAwaken/GoFesl`.

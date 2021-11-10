# URL Shortener

URL Shortener is a Go & Vite & Vue 3 & TypeScript & Tailwind project for dealing with URL shorten service.

## Installation

Clone the project, run `build.sh` or `build.bat`, then change dir to `dist` and run following command:

```bash
CONFIGFILE=../config/config.json WEBPATH=../web/dist ./urlshortener
```

or

```bat
set CONFIGFILE=../config/config.json&set WEBPATH=../web/dist&./urlshortener&set CONFIGFILE=&set WEBPATH=
```

## Development

Frontend, change to `web` directory frist:

- Install dependencies

```bash
yarn
```

- Run development server

```bash
yarn dev
```

Backend:

- Download dependencies

```bash
go mod download
```

- Run development server

```bash
go run main.go
```

## Docker

Build:

```
docker build . -t urlshortener
```

Run:

```
docker run -p 8082:8082 urlshortener
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](LICENSE.md)

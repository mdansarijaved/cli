# Cli tool for Younicorn Project a self hosting cloud solution for everyone

This is a command-line interface (CLI) tool for self hosting your project using you own VPS.


## Installation
on you ubuntu server run the following command to install the cli tool

```bash
wget https://github.com/mdansarijaved/cli/releases/download/v1.1.5/main
```

## Usage

```bash
sudo chmod +x main
sudo ./main
```

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

```bash
"DATABASE_URL",
"GITHUB_CLIENT_ID",
"GITHUB_CLIENT_SECRET",
"NEXTAUTH_SECRET",
"NEXTAUTH_URL",
"ADMIN_MAIL",
"ADMIN_PASS",
"DOMAIN",
```
just be sure to add the correct values to the variables and you're all set.



## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](LICENSE.md)

```

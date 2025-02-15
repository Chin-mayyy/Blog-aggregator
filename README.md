# Gator CLI

Gator is a command-line tool written in Go that interacts with a PostgreSQL database. This guide will help you install and configure Gator for your system.

## Prerequisites

Before using Gator, ensure you have the following installed:

- Go (version 1.20 or later recommended)
- PostgreSQL
- Git (for cloning the repository)
- Sqlc (for generating SQL code)

## Installing Go

To install Go using terminal commands, follow these steps:

### **Linux/macOS**
```sh
curl -OL https://go.dev/dl/go1.20.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### **Windows (Using PowerShell)**
```sh
winget install --id=GoLang.Go -e
```

Verify the installation:
```sh
go version
```

## Installation

### Clone the Repository

```sh
git clone https://github.com/Chin-mayyy/Blog-aggregator.git
cd Blog-aggregator
```

### Install Gator

Use `go install` to build and install the Gator CLI:

```sh
go install ./...
```

After installation, ensure the `GOPATH/bin` directory is in your system's `PATH` so you can run `gator` globally.

```sh
export PATH=$PATH:$(go env GOPATH)/bin
```

## Installing SQLC

SQLC is required for generating Go code from SQL queries. Install it using the following commands:

### **Linux/macOS**
```sh
brew install sqlc
```

Or manually:
```sh
curl -sSL https://github.com/sqlc-dev/sqlc/releases/latest/download/sqlc_$(uname -s)_$(uname -m).tar.gz | tar -xz
sudo mv sqlc /usr/local/bin
```

### **Windows (Using PowerShell)**
```sh
scoop install sqlc
```

Verify the installation:
```sh
sqlc version
```

## Configuration

Gator requires a configuration file to connect to your PostgreSQL database. Create a `gatorconfig.json` file in the root directory with the following content:

```json
{
  "postgres_url": "postgres://your-username:your-password@localhost:5432/your-database?sslmode=disable",
  "username": "your-username"
}
```

You can create the file using the following command:

```sh
echo '{"postgres_url": "postgres://your-username:your-password@localhost:5432/your-database?sslmode=disable", "username": "your-username"}' > gatorconfig.json
```

Gator requires a configuration file to connect to your PostgreSQL database. Create a `config.yaml` file in the root directory and populate it with the necessary credentials:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "your-username"
  password: "your-password"
  dbname: "your-database"
  sslmode: "disable"
```

## Running Gator

Once installed and configured, you can start using Gator with the following command:

```sh
gator
```

## Available Commands

- `gator init` – Initializes the database.
- `gator fetch` – Fetches new records from the database.
- `gator list` – Lists all stored records.
- `gator help` – Displays help information.
- `gator login` – Logs in a user.
- `gator register` – Registers a new user.
- `gator reset` – Resets a user’s password.
- `gator users` – Lists all users.
- `gator agg` – Runs data aggregation.
- `gator addfeed` – Adds a new feed (requires login).
- `gator feeds` – Lists all feeds.
- `gator follow` – Follows a user.
- `gator following` – Lists users being followed.
- `gator unfollow` – Unfollows a user (requires login).

## Running in Development Mode

For development, you can use:

```sh
go run .
```

For production, simply run:

```sh
gator
```

## Submitting Your Repository

After pushing Gator to GitHub, submit your repository link in the format:

```
https://github.com/github-username/repo-name
```

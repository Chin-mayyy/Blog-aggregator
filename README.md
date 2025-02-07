Gator CLI

Gator is a command-line tool written in Go that interacts with a PostgreSQL database. This guide will help you install and configure Gator for your system.

Prerequisites

Before using Gator, ensure you have the following installed:

Go (version 1.20 or later recommended)

PostgreSQL

Git (for cloning the repository)

Installation

Clone the Repository

 git clone https://github.com/github-username/repo-name.git
 cd repo-name

Install Gator

Use go install to build and install the Gator CLI:

go install ./...

After installation, ensure the GOPATH/bin directory is in your system's PATH so you can run gator globally.

export PATH=$PATH:$(go env GOPATH)/bin

Configuration

Gator requires a configuration file to connect to your PostgreSQL database. Create a config.yaml file in the root directory and populate it with the necessary credentials:

database:
  host: "localhost"
  port: 5432
  user: "your-username"
  password: "your-password"
  dbname: "your-database"
  sslmode: "disable"

Running Gator

Once installed and configured, you can start using Gator with the following command:

gator

Available Commands

gator init – Initializes the database.

gator fetch – Fetches new records from the database.

gator list – Lists all stored records.

gator help – Displays help information.

Running in Development Mode

For development, you can use:

go run .

For production, simply run:

gator

Submitting Your Repository

After pushing Gator to GitHub, submit your repository link in the format:

https://github.com/github-username/repo-name

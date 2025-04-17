# GitHub User Activity CLI

A command-line application to fetch and display recent GitHub user activity, built with Go. The application uses the GitHub API to retrieve user events and presents them in the terminal.

## Features
- Fetch recent activity for a specified GitHub user
- Display activities such as pushing commits, starring repositories, and opening issues

## Requirements
- Go 1.21 or later
- A GitHub Personal Access Token with public_repo scope

## Installation
1. Clone or download this repository.
2. Navigate to the project directory:
```
cd github-user-activity
```
3. Initialize Go modules if not already done:
```
go mod init github-user-activity
```
4. Install dependencies (e.g., godotenv):
```
go get github.com/joho/godotenv
```
5. Create a .env file in the project root with your GitHub token:
```
GITHUB_TOKEN=your_token_here
```

Alternatively, set the GITHUB_TOKEN environment variable manually:
```
export GITHUB_TOKEN=your_token_here
```

6. Build the application:
```
go build -o github-user-activity cmd/main.go
```

## Usage
Run the application using the compiled binary (e.g., ./github-activity).

### Commands
| Command                           | Description                              | Example                                    |
|-----------------------------------|------------------------------------------|--------------------------------------------|
| `<username>`                      | Fetch and display user activity          | `./github-user-activity john`              |

```
Example Output
- Pushed 3 commits to octocat/Hello-World
- Starred octocat/Hello-World
- Opened a new issue in octocat/Hello-World
```

Notes:
- Ensure the GITHUB_TOKEN environment variable is set or a .env file is present in the project root.
- Ensure your GitHub token has the necessary permissions (public_repo scope) to access user events.
- For more details on the GitHub API, see the [official documentation](https://docs.github.com/en/rest/activity/events?apiVersion=2022-11-28).

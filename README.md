# AI Command Line Assistant

A command-line tool that leverages OpenAI's API to provide help with commands and generate commands based on user input.

## Features

- **AI Help**: Get detailed explanations and examples for any command, similar to the `man` command but with examples.
- **AI Command Generation**: Generate command-line commands based on natural language descriptions.

## Installation

### Prerequisites

- Go 1.24 or higher
- OpenAI API key

### Steps

1. Clone the repository:
   ```
   git clone <repository-url>
   cd awesomeProject
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Build the application:
   ```
   go build -o ai-helper
   ```

## Configuration

The application requires the following environment variables:

- `API_KEY`: Your OpenAI API key
- `API_MODEL`: The OpenAI model to use (e.g., "gpt-4", "gpt-3.5-turbo")

Example of setting environment variables:

```bash
# On Windows
set API_KEY=your-api-key
set API_MODEL=gpt-4

# On Linux/macOS
export API_KEY=your-api-key
export API_MODEL=gpt-4
```

## Usage

### AI Help Command

Get help with any command with examples:

```bash
./ai-helper ai-help -c "git commit"
```

This will provide detailed information about the specified command, similar to a man page but with practical examples.

### AI Command Generation

Generate a command based on a natural language description:

```bash
./ai-helper ai-cmd -c "how to find all files modified in the last 24 hours"
```

This will generate the appropriate command-line command based on your description, including an example of how to use it.

## Error Messages

- If you don't provide a command with the `-c` flag, you'll see: "You must provide a command for help with ai"
- If the `API_KEY` environment variable is not set, you'll see: "please provide env variable API_KEY"
- If the `API_MODEL` environment variable is not set, you'll see: "You must provide a command for help with ai"

## Dependencies

- [github.com/openai/openai-go](https://github.com/openai/openai-go): OpenAI Go client
- [github.com/urfave/cli/v2](https://github.com/urfave/cli/v2): Command line interface framework


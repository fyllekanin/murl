
# Murl

Murl is a command-line tool designed to help manage software versions available on your system. It allows you to easily list and manage versions for various software like Go, Node.js, Maven, and Java.

## Features

- Manage software versions for common tools.
- List available stable and unstable versions for supported software.
- Simple CLI interface.

TEST

## Installation

### Prerequisites

- [Any prerequisites, like Go, Node.js, etc.]

### Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/murl.git
   ```
2. Navigate into the project directory:
   ```bash
   cd murl
   ```
3. Build the project:
   ```bash
   go build
   ```

## Usage

Murl works by specifying the software you want to manage and the command you want to execute.

### General Command Syntax
```bash
murl [flags] <software> <command>
```

### Available Software

- **go**
- **nodejs**

### Available Commands

- `list` - Lists the available versions for the specified software.
  - `-unstable` - Include unstable versions in the list if applicable.
- `install <version>` - Install the specific version

### Example Usage

- List available versions of Go:
   ```bash
   murl go list -unstable
   ```

- List available versions of Node.js, including unstable versions:
   ```bash
   murl nodejs list
   ```

# License

Cobra is released under the Apache 2.0 license. See [LICENSE.txt](LICENSE.txt)

# Golang AI Project

A development environment for building AI-powered Go applications using Genkit Go framework with Docker Model Runner and Docker Agentic Compose.

## Overview

This project provides a containerized development workspace for creating Go applications that integrate with AI models. It includes:

- **Genkit Go Framework**: Firebase Genkit framework for building AI-powered applications in Go
- **Docker Model Runner**: Provides AI model inference capabilities via OpenAI-compatible API
- **Docker Agentic Compose**: Orchestrates AI models and services using Docker Compose
- **Development Container**: A fully configured Ubuntu-based development environment with Go 1.25.5, Docker, and essential tools

## Getting Started

### Prerequisites

- Docker Desktop with DevContainer support
- Visual Studio Code with Remote - Containers extension

### Launch the Development Environment

1. Open the project in VS Code
2. When prompted, click "Reopen in Container" or run the command palette action: `Dev Containers: Reopen in Container`
3. VS Code will build and start the development container automatically
4. Once initialized, you'll have a complete Go development environment with all dependencies installed

The container will:
- Mount your workspace to `/workspaces/${localWorkspaceFolderBasename}`
- Forward ports: 3000, 4000, 8080, 9090, 7070, 6443
- Share your Git configuration and SSH keys for seamless repository access
- Provide Docker-in-Docker capabilities for container orchestration

## DevContainer Configuration Explained

### Architecture Components

#### [Dockerfile](.devcontainer/Dockerfile)
Builds an Ubuntu 22.04 image with:
- **Go 1.25.5**: Complete Go toolchain with gopls, go-outline, gocode, and revive
- **Docker**: Full Docker CE installation for container management
- **Development Tools**: git, curl, wget, jq, build-essential, bat, eza, socat
- **User Configuration**: Creates a user matching your local username with sudo and docker group permissions
- **Oh My Bash**: Enhanced shell experience

#### [compose.yml](.devcontainer/compose.yml)
Defines the development workspace service:
- Uses `network_mode: "host"` for simplified networking with local services
- Mounts the workspace directory with cached consistency for better performance
- Configurable Go version via `GO_VERSION` arg (default: 1.25.5)
- User context passed via `USER_NAME` environment variable

#### [devcontainer.json](.devcontainer/devcontainer.json)
VS Code DevContainer configuration:

**Features**:
- Git integration (latest version)

**Extensions** (automatically installed):
- Docker tools (ms-azuretools.vscode-docker, docker.docker)
- Go development (golang.go)
- Python support (ms-python.python)
- UI enhancements (Material Icon Theme, Color Picker)
- Documentation tools (Markdown Mermaid, DrawIO)
- Code highlighting and utilities

**Environment Variables**:
- `LOCAL_WORKSPACE_FOLDER`: Reference to your host workspace path
- `MODEL_RUNNER_HOST`: Endpoint for Docker Model Runner (`http://model-runner.docker.internal`)

**Mounts**:
- Docker socket for Docker-in-Docker functionality
- Git configuration for commit authorship
- SSH keys for secure repository access

**Post-Create Commands**:
- Configures Git safe directory
- Sets Docker socket permissions

## Running the Compose Agent Example

The [compose-agent](compose-agent) directory contains a sample Genkit Go application that demonstrates AI model integration.

### What It Does

The application uses Genkit Go to create a D&D dungeon master AI that generates NPC (Non-Player Character) profiles. It connects to Docker Model Runner to access the `jan-nano` model via an OpenAI-compatible API.

### Launch Instructions

1. **Navigate to the compose-agent directory**:
   ```bash
   cd compose-agent
   ```

2. **Start the application with Docker Compose**:
   ```bash
   docker compose up
   ```

   This command:
   - Builds the `budgie-agent` service defined in [compose-agent/compose.yml](compose-agent/compose.yml)
   - Starts Docker Model Runner with the `qwen2_5` model (hf.co/menlo/jan-nano-gguf:q4_k_m)
   - Injects environment variables `MODEL_RUNNER_BASE_URL` and `MODEL_RUNNER_CHAT_MODEL` into the container
   - Runs the Go application which streams AI-generated D&D NPC characteristics

3. **Expected Output**:
   The application will stream the generated NPC details to the console, including the elf name and characteristics.

### Configuration Details

**[compose-agent/compose.yml](compose-agent/compose.yml)**:
- Defines `budgie-agent` service that builds from local Dockerfile
- Maps the `chat-agen` model to environment variables for the agent
- Declares the `chat-agen` model using Hugging Face's jan-nano model (quantized q4_k_m version)

**[compose-agent/main.go](compose-agent/main.go)**:
- Initializes Genkit with OpenAI plugin configured for Docker Model Runner
- Reads `MODEL_RUNNER_BASE_URL` and `MODEL_RUNNER_CHAT_MODEL` from environment
- Sends a system prompt defining the AI as a D&D dungeon master
- Requests NPC generation with streaming output
- Temperature set to 0.7 for creative responses

### Customization

To use different models or change behavior:
- Edit the `model` field in [compose-agent/compose.yml](compose-agent/compose.yml) to use alternative Hugging Face models
- Modify the prompts in [compose-agent/main.go](compose-agent/main.go) to change the AI's role and output
- Adjust `temperature` and other parameters in the `ai.WithConfig()` call

## Project Structure

```
.
├── .devcontainer/
│   ├── Dockerfile          # Container image definition
│   ├── compose.yml         # Docker Compose service configuration
│   └── devcontainer.json   # VS Code DevContainer settings
└── compose-agent/
    ├── main.go             # Genkit Go application
    ├── compose.yml         # Docker Agentic Compose configuration
    ├── Dockerfile          # Agent container definition
    └── go.mod              # Go dependencies
```

## Development Workflow

1. Make code changes in VS Code within the container
2. Use the integrated terminal for Go commands (`go run`, `go build`, `go test`)
3. Test Docker Compose applications using `docker compose up`
4. Access forwarded ports from your host machine
5. Git operations work seamlessly with your host credentials

## Troubleshooting

- **Docker socket permission issues**: The post-create command sets permissions, but if problems persist, run `sudo chmod 666 /var/run/docker.sock`
- **Model Runner connection errors**: Ensure `MODEL_RUNNER_HOST` is correctly set and Model Runner is accessible
- **Port conflicts**: Check if forwarded ports (3000, 4000, 8080, 9090, 7070, 6443) are available on your host


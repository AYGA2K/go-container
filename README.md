# Lightweight Containerization Tool in Go

## Overview

This project is a lightweight containerization tool developed in Go. It simulates basic Docker functionalities, including process isolation and filesystem management, using Linux namespaces and `chroot`. Additionally, it incorporates cgroups for basic resource management.

## Features

- **Process Isolation**: Uses Linux namespaces to isolate processes.
- **Filesystem Management**: Implements `chroot` to set up a container-like filesystem.
- **Resource Management**: Integrates cgroups to control resource usage.

## How It Works

1. **Process Isolation**: The tool uses Linux namespaces (`CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS`) to create isolated environments for processes.
2. **Filesystem Setup**: Uses `chroot` to change the root directory to a specified filesystem, simulating a container environment.
3. **Resource Management**: Sets up cgroups to manage resources such as memory limits for the isolated processes.

## Getting Started

### Installation

Clone the repository:

```bash
git clone https://github.com/AYGA2K/go-container.git
cd go-container
```

Download the ubuntu base filesystem:

```bash
curl -O https://cdimage.ubuntu.com/ubuntu-base/releases/22.04/release/ubuntu-base-22.04-base-amd64.tar.gz
```

Extract the Ubuntu base filesystem into a directory that will be used as the container's root filesystem (rootfs):

```bash
tar -xvzf ubuntu-base-22.04-base-amd64.tar.gz -C rootfs/
```

Build the project:

```bash
go build -o go-container
```

### Usage

To run a command within an isolated environment:

```bash
sduo ./go-container run /bin/bash
```

This command will execute `/bin/bash` in a new container-like environment with process isolation and resource management.

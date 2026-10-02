# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Goplant** is a tool for generating simple reverse-shell implant binaries. It's designed for pentesting use with tools like Penelope or nc, aiming to bypass basic AV detection and make obtaining a reverse shell as straightforward as possible.

## Build & Run Commands

```bash
# Build the CLI tool
go build -o goplant .

# Run with required arguments (LHOST and LPORT)
./goplant -H 192.168.1.100 -p 4444

# Run with custom output path
./goplant -H 192.168.1.100 -p 4444 -o /path/to/output

# View help
./goplant -h
```

## Project Structure

### Core Components

- **`internal/cli/`** — Command-line argument parsing
  - `Parse()` returns an `Args` struct with LHOST, LPORT, and Output fields
  - Uses `argparse` library; `-h`/`--help` handled automatically

- **`internal/template/`** — Template and build system
  - `template.go` — Loads embedded templates and exposes `GetAvailableTemplates()` and `ListTemplateNames()`
  - `builder.go` — Build command generation; `BuildGoImplant()` returns a go build command with randomized ldflags
  - `templates/` — Go text template files (`.tmpl`) embedded at compile time

- **`main.go`** — Entry point; parses CLI args, prints implant generation info, and lists available templates

### Key Design Notes

- Templates are embedded using Go's `//go:embed` directive, making the binary self-contained
- `BuildGoImplant()` generates random `BuildID` and mutex names for each build (obfuscation via ldflags)
- Planned features (encryption, code obfuscation, stager/server) remain as commented-out scaffolding in `main.go`

## Future Work

The README lists planned features:
- Binary encryption and randomization
- Code obfuscation
- Clean CLI interface
- GitHub installation support
- Stager/implant server

Most infrastructure for these exists as stubs; the main focus currently is on the template and build system.

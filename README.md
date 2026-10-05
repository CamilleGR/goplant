
# Goplant

![goplant logo](img/goplant_logo.png)

Goplant generates executable reverse shell implants for pentesting and CTF. 


```bash
Goplant generates executable reverse shell implants for pentesting

Usage:
  goplant [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a reverse shell implant
  help        Help about any command
  list        List available templates
  serve       Start an implant server

Flags:
  -h, --help   help for goplant

Use "goplant [command] --help" for more information about a command.
```

## Generate 

The `generate` command generate a new executable accordingly to targetted arch ans operating system. It take lhost and lport parameters as metasploit or msfvenom.

## Serve 

The serve command allows you to directly generate an implant and serve it on a web server. It reduce the number of action to get your reverse shell. 

Goplant show you some example commands to execute to download and execute your payload
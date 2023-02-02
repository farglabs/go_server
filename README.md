# go_server
An extremely simple base server in golang for jump-starting an HTTP API endpoint microservice project. Secure, light-weight and can be built for any platform to a single executable!

At this point, all served code/data should be contained within main.go (or at least called from there). The inability to serve static HTML files (and other files) is intentional; if you want to take on that security risk, go use Apache or Nginx.

## Instructions
Just build and run!

Build:

``$ go build``

Run on Linux:

``$ ./server``

Run on Windows:

``$ server.exe``


## Alternate Instructions
Download a prebuilt executable:

[linux-build]: https://github.com/farglabs/go_server/raw/master/build/go_server "Linux pre-built executable"

[windows-build]: https://github.com/farglabs/go_server/raw/master/build/go_server.exe "Windows pre-built executable"

[settings-ini]: https://github.com/farglabs/go_server/raw/master/build/.example.settings.ini "Example settings.ini"

Linux [pre-built executable][linux-build] download

Windows [pre-built executable][windows-build] download


Download [the .example.settings.ini file][settings-ini]
Rename the file to `settings.ini` and modify the settings as you see fit

Move the `settings.ini` file to the same directory as the pre-built executable

Run the executable (you may need to `chmod +x` the executable first if you're on Linux)

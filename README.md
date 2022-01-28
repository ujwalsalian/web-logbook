# Web-logbook

This is a simple free EASA style logbook local web app written in golang.

You can clone the repo and compile the binaries yourself, or just download the latest ones for your operating system from the [releases](https://github.com/vsimakhin/web-logbook/releases).

Once you start the app it automatically creates a sqlite local db and start listening on a port 4000 by default. So you can open it in your standard web-browser on http://localhost:4000

You also can easily export all flight records into EASA style pdf format, print it, sign and use as a usual paper logbook.

# Usage

1. Download the latest release from https://github.com/vsimakhin/web-logbook/releases
1. Extract archive to some folder/directory
1. Run: 
  * Windows:
    * Double click on the `web-logbook.exe` file. It will show you some warning about how unsafe it can be (need to solve it later), but just run it.
  * Linux:
    * Open a terminal and navigate to the directory
    * Run `./web-logbook`
  * MacOS:
    * *I still didn't test it for the MacOS, so the binaries will be added later*
4. Open your browser and type http://localhost:4000
5. Once you finished, use `Ctrl+C` or just close the terminal window

# Supported operating systems

Since it's written in golang it can run on any system if you compile the sources. For now in the `Release` page there are 2 binaries for linux amd64 and windows.

# Interface

Currently there are implemented several modules in the logbook app:
* Logbook itself
* Export to EASA PDF format
* Settings
* Map
* Licensing & Certification
* Statistics

## Logbook

![Main logbook page](https://github.com/vsimakhin/web-logbook-assets/raw/main/logbook-main.png)

## Export

![Export to PDF](https://github.com/vsimakhin/web-logbook-assets/raw/main/logbook-export.png)

## Flight record

![Flight record](https://github.com/vsimakhin/web-logbook-assets/raw/main/flight-record-example.png)

## Settings

![Settings](https://github.com/vsimakhin/web-logbook-assets/raw/main/settings.png)

## Stats

![Flight stats](https://github.com/vsimakhin/web-logbook-assets/raw/main/stats.png)


![Map](https://github.com/vsimakhin/web-logbook-assets/raw/main/stats-map.png)

## Licensing & Certifications

![Licensing](https://github.com/vsimakhin/web-logbook-assets/raw/main/licensing.png)


# Used libraries

* Bootstrap https://getbootstrap.com/
* Datatables https://datatables.net/
* Openlayers https://openlayers.org/
* Golang gofpdf https://github.com/jung-kurt/gofpdf
* Golang chi web-server https://github.com/go-chi/chi

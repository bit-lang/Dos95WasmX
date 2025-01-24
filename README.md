# Dos 95 Wasm X
a simple wrapper around [doswasmx](https://github.com/nbarkhina](https://github.com/nbarkhina/DosWasmX)).

credits goes to nbarkhina.

## play your favorite windows 95 games, ***in browser***
![windows 95 gaming](https://github.com/bit-lang/Dos95WasmX/blob/master/images/win95-games01.png)
![windows 95 gaming](https://github.com/bit-lang/Dos95WasmX/blob/master/images/win95-games02.png)
![windows 95 gaming](https://github.com/bit-lang/Dos95WasmX/blob/master/images/win95-games03.png)
![windows 95 gaming](https://github.com/bit-lang/Dos95WasmX/blob/master/images/win95-games04.png)

## what is included in the release
- the app is launching a simple web server, then users can use Chrome / Edge / Firefox to access the embedded Windows instance page
- Windows executable file in "dos95wasmxen-win-AMD64.zip", embedding a full windows 95 installation, with 512MB hard disk. the executable can run on Windows 7 / 10 / 11
- linux program in "dos95wasmxen-linux-AMD64", with same embedded contents, runs on AMD64 versions of debian11/12, or ubuntu 22.04/24.04

## how to use
- prerequisites: unzip the zip file matching your operating system. directly run executables from within the zip container ***DOES NOT WORK***

1. in windows command line, run dos95wasmxen.exe, ***leave the command console open, do not close it***
![run the app](https://github.com/bit-lang/Dos95WasmX/blob/master/images/dos95_start.png)
2. use your browser to navigate to [URL](http://localhost:8088) <http://localhost:8088>
3. drag a game CD into the center frame. the ISO image will begin to load and then windows will start
![ISO loading](https://github.com/bit-lang/Dos95WasmX/blob/master/images/win95-start.png)
4. click on PowerOn button at the top of browser page, if windows does not automatically start
5. install your game and enjoy
6. pre-installed startcraft 1.0 is playable without installation, provided you use the same version of StarCraft game CD.

## where to get windows 95 game CDs
- [starcrarft 1 is free now](https://www.cnet.com/tech/computing/how-to-download-the-original-starcraft-for-free/)
- or, you can get game CDs from <archive.org>

## special notes on how to properly stop the app
1. please ***always shutdown windows*** after exiting your game
2. click on Power Off button and close the browser, this way the data will be saved to your browser cache. if you close browser without shutting down windows first and "Power Off", your game data is lost.
3. remember the command prompt where your started and left open? go there and type "stop" at the prompt to stop the app.

## please refer to doswasmx readme for details on build instructions.

## DOS Wasm X features (from the author)
A browser based DOS emulator designed around ease of use and stability. It is based on the newer DosBox-X codebase and thus supports both Windows 95 and Windows 98 installations. However if you just want to use DOS applications and games you can stay in DOS mode. To begin using it simply drag and drop any application or game files onto the emulator. You can then save your hard disk with the click of a button or just exit if you want to discard your changes. I went with a simple and clean interface to try and make it approachable and non-intimidating. The hard disk saves directly in your browser so you can come back later and continue where you left off. It's like your own personal virtual machine on the web!

Supports the following features -
- Fully web based application - using web assembly
- Save hard drive to the browser (512mb, 1 gig, or 2 gig options)
- Automatic support for a variety of file formats (Iso, Zip, Bin, Cue, Img, 7z)
- Customize RAM (32mb, 64mb, 128mb)
- Import/export files into and out of the emulator
- Export your entire hard disk image for local saving
- Load/change CD while emulator is running
- Floppy Disk Support
- Audio support
- Full screen
- Zoom controls
- Mouse capture
- Resize resolution
- 16 and 32 bit color (via 16 bit color fix)
- Customize CPU speed
- Host the application yourself
- Customize startup hard drive image
- Send CTRL/ALT/DELETE
- Pause/Unpause
- Import existing IMG hard disk if you already have one

# Installing your own Windows version
DOS Wasm X supports installing Windows 95 or Windows 98 using your own copy of Windows. Simply drag and drop the ISO onto the startup page. DOS Wasm X will detect the Windows CD and begin the installation process. If you choose to Install Windows 95 you may get the error below. Simply click OK and then cancel when it asks you for the Path to the CD. This will allow you to continue with the installation. The reason for this error is because at this stage of the process the CD drivers have not yet been loaded. However after restarting Windows it will detect the CD Drive and finish installing the drivers successfuly. Always remember to shut down windows in the guest OS before exiting the page. This will automatically save your hard drive changes to the browser and prevent scandisk from running the next time you boot into Windows.

# Disclaimer
This app was made for fun and is not affiliated or associated with Microsoft.

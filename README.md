# Clacks Micro

> The Clacks is the informal nickname for the semaphore system that is the fastest means of non-magical communication on the Discworld. [1](https://wiki.lspace.org/mediawiki/Clacks)

Clacks Micro displays Clacks-style messages using micro-controllers.

It uses [tiny go](https://tinygo.org/) to compile to hardware devices.

It is intended to be used by [Clex](https://github.com/GrandTrunkSemaphoreCompany/clex) as a generic bitstream client target. This should also make the transition to a more physical (and louder) mechanical clacks system easier, and hopefully relatively seamless.

## Supported boards / devices

It supports as micro controllers:
- [https://www.adafruit.com/product/3857](Adafruit Feather M4)

It supports as output devices:
- [SSD1306 OLED display](https://cdn-shop.adafruit.com/datasheets/SSD1306.pdf)

## Compiling
1. Run `make flash`

## Running
1. Compile
    - `$ make flash`
2. Determine usb port
    - `$ ./findusb.sh`
    - `/dev/ttyACM0 - Adafruit_Adafruit_Feather_M4`
3. Connect using screen
    - `$ screen /dev/ttyACM0 9600`
4. Type message, use return to send.    

## Debugging

Debugging can be done via `fmt.Print` messages which will be echoed back onto a screen connection.
# Clacks Micro

> The Clacks is the informal nickname for the semaphore system that is the fastest means of non-magical communication on the Discworld. [1](https://wiki.lspace.org/mediawiki/Clacks)

Clacks Micro displays Clacks-style messages using micro-controllers.

It uses [tiny go](https://tinygo.org/) to compile to hardware devices

It is intended to be used by [Clex](https://github.com/GrandTrunkSemaphoreCompany/clex) as a generic bitstream client target. This should also make the transition to a more physical (and louder) mechanical clacks system easier, and hopefully relatively seamless.

## Supported boards / devices

It supports as micro controllers:
- [https://www.adafruit.com/product/3857](Adafruit Feather M4)

It supports as output devices:
- [SSD1306 OLED display](https://cdn-shop.adafruit.com/datasheets/SSD1306.pdf)
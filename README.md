LimitlessLED
============
This is a small library for controlling [LimitlessLED lightbulbs][1].  These bulbs operate via a small wifi bridge, which translates UDP packets (control sequences are documented in control.go) into 2.4GHz RF commands which are sent to the bulbs.

This package also provides support for storing bulb state in a database, seeing as the bulbs themselves don't allow you to simply send states.  The database implementation is an ugly hack, but I wanted to get it committed so that I could start consuming this library from the [atomic][2] project.

[1]: http://www.limitlessled.com/
[2]: http://github.com/frio/atomic

Limitations
-----------
Currently, only a subset of the controls are implemented -- notably, there is no support yet for the current bulb (because I don't have a spare socket to put mine in!) or controlling individual zones.  The library will evolve to support these.
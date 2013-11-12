LimitlessLED
============
This is a small library for controlling [LimitlessLED lightbulbs][1].  These bulbs operate via a small wifi bridge, which translates UDP packets (control sequences are documented in control.go) into 2.4GHz RF commands which are sent to the bulbs.

[1]: http://www.limitlessled.com/

Limitations
-----------
Currently, only a subset of the controls are implemented -- notably, there is no support yet for the current bulb (because I don't have a spare socket to put mine in!) or controlling individual zones.  The library will evolve to support these.

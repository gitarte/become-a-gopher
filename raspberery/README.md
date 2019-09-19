### WHICH LIBRARY
Basically what you have is a choice of two:

* https://gobot.io/
* https://github.com/stianeikeland/go-rpio

Both have their pros and cons. As for this example we choose go-rpio, because it is lean and pure Go solution. It laks some features like I2C, but for basic usage it is fast, simple and railable. 

### CROSS COMPILATION FOR RASPBERRY
```bash
GOOS=linux GOARCH=arm GOARM=5 go build 
```

### RASPBERRY PINOUT
![PINOUT](header_pinout.jpg)

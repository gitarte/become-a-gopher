# SMS Tool
Application serves HTTP endpoint that sends given messages to given numbers

### Prereq
- GSM device that is able to be discovered as TTY device
- GAMMU

### Install gammu
```bash
apt install gammu
gammu-config -c ./gammu-config
```

### Test gammu
Below `TEXT` is not the message but SMS protocol parameter. 
```bash
/bin/echo {message} | \
sudo /usr/bin/gammu \
    -c ./gammu-config \
    sendsms \
    TEXT \
    {destination-number}
```

### Usage
```bash
curl \
    -X POST \
    -H "Content-Type:application/json" \
    -d '{"content":"elo!","number":"123456789"}' \
    localhost:8080/sms
```
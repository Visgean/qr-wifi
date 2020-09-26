# CLI WiFi passwords sharing on Ubuntu

This is **very hacky** way to create a QR code with wifi password on Ubuntu. 

It is hacky because:
- it just goes into `/etc/NetworkManager/system-connections` and reads the `.ini` files. This is almost certainly a bad way of doing this. 
- It needs sudo access.
- I wrote it for fun to try doing a bit of go. 
- I don't plan to maintain it whatsoever :) 


![](example.png)

# Kudos:

This project is based on [@fumiyas/qrc](https://github.com/fumiyas/qrc).


# Installation:

```
go get github.com/Visgean/qr-wifi
```

(Plus you need to make sure that you have a `$GO_PATH/bin` in your `.bashrc` path.)

# Usage

Show QR codes for all networks:

```
$ sudo qr-wifi 
```

Filter networks by name:

```
$ sudo qr-wifi Starbucks
```


# Hacky way to get rid of sudo:

I only use it on personal, encrypted computer to which no-one else has access to, so I don't care much for security on this level. Therefore I am happy to use this permission to avoid typing sudo: 

```
sudo chown root:root `which qr-wifi`
sudo chmod 4775 `which qr-wifi`
```

Read more about it [here](https://unix.stackexchange.com/questions/18830/how-to-run-a-specific-program-as-root-without-a-password-prompt)
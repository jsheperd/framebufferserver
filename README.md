# framebufferserver
Linux framebuffer can be used to present images on linux device without X11. It is just a simple REST API around it.

# Use case

Create snapshot from your framebuffer with 
```
cp /dev/fb0 snapshot.data
```

You can present it on a remote machine after the X session disabled
```
sudo killall gdm-x-session
sudo dd if=snapshot.data of=/dev/fb0
```

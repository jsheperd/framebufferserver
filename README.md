# framebufferserver
Linux framebuffer can be used to present images on linux device without X11. It is just a simple REST API around it.

# Use case
'''
sudo killall gdm-x-session

dd if=snapshot.data of=/dev/fb0
'''

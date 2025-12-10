## gpio2mqtt
This is a simple utility to read GPIO pin states and publish their states to an MQTT broker. It can be useful for home automation projects where you want to monitor physical switches or sensors connected to the GPIO pins.

1. Build  
   1. Install go
       ```bash
       # for RedHat/CentOS/Rocky/Fedora/Alma
       sudo dnf install golang
    
       # for Debian/Ubuntu
       sudo apt install golang
       ```  
   2. Clone the repository
       ```bash
       git clone https://github.com/Biiddd/gpio2mqtt.git
       ```  
   3. Resolve dependencies and build
      ```bash
      cd gpio2mqtt
      go build
      ```
2. Configure
3. Run
   1. Run at foreground
       ```bash
       ./gpio2mqtt
       ```  
   2. Run as service
       ```bash
       # modify gpio2mqtt.service to set the correct path to the binary
       sudo cp gpio2mqtt.service /etc/systemd/system/
       sudo systemctl daemon-reload
       sudo systemctl enable gpio2mqtt
       sudo systemctl start gpio2mqtt
       ```
4. License
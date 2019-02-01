# urodynamics

##Notes
1. https://github.com/EddyVerbruggen/nativescript-bluetooth/tree/v2/demo-ng
- Use this demo to get started with Bluetooth LE connection
- Understand this is a Bluetooth Low Energy connection!
- This uses a beta version of nativescript-bluethooth@2.0.0-beta.2
  - Run tns plugin add nativescript-bluetooth@2.0.0-beta.2
2. Defualt bluetooth service on RPi3 does not support low energy advertisements. Use the following link to enable experimental low enery feature: [link]https://stackoverflow.com/questions/41351514/leadvertisingmanager1-missing-from-dbus-objectmanager-getmanagedobjects
3. To access the bluetooth settings on the RPi run: 'sudo bluetoothctl'
4. Found working implementation of Golang BLE [link]https://github.com/currantlabs/ble/tree/master/examples/blesh
5. Connect with app in 1
6. Steps to recreate
   - Use link in 4 to start up BLE server
   - In app, scan fot any peripheral
   - Connect to Gopher
   - Access service: "0001000000011000800000805F9B34FB"
   - From here you can read, write
   - And subscribe/start notify
   - Sends message every second
   - To change whatis sent change code in this function: c.HandleNotify(ble.NotifyHandlerFunc(func(req ble.Request, n ble.Notifier) {}

##TODO

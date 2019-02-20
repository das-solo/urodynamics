import board
import busio
import time

i2c = busio.I2C(board.SCL, board.SDA)

import adafruit_ads1x15.ads1115 as ADS

from adafruit_ads1x15.analog_in import AnalogIn



starttime=time.time()
while True:
    ads = ADS.ADS1115(i2c)
    ads.gain = 1

    # chan = AnalogIn(ads, ADS.P0)
    # chan1 = AnalogIn(ads, ADS.P1)
    chan = AnalogIn(ads, ADS.P0, ADS.P1)

    print("chan0:", chan.value, chan.voltage)
    # print("chan1:", chan1.value, chan1.voltage)

    time.sleep(0.5)

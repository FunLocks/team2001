import RPi.GPIO as GPIO
import time
import signal
import sys
import GPS 
import requests
import json
import subprocess

def handler(signal, frame):
	GPIO.cleanup()
	print("GPIO cleanup...")
	print("exit!")
	sys.exit(0)

def playSoiya():
	playsound("~/soiyaMix.wav")

if __name__ == '__main__':
	signal.signal(signal.SIGINT, handler)
	
	GPIO.setmode(GPIO.BCM)
	GPIO.setup(23,GPIO.IN) # switch
	GPIO.setup(24,GPIO.OUT) # ssr

	gps = GPS.GPS()
	gps.update()

	url = "http://153.120.166.49:8080/ahchoo/post"
	headers = {"Content-Type" : "application/json"}

	flag = True;
	while True:
		sw_raw = GPIO.input(23)
		sw = 0
		
		if flag == True and sw_raw == 1:
			sw = 1
			flag = False
		
		elif sw_raw == 0:
			flag = True

		if sw == 1:
			print("switch pushed!!!")
			gps.update()
			lat = str(gps.latitude()[0])
			lng = str(gps.longitude()[0])
			print("gps data format done")
			data = {"latitude" : lat, "longitude": lng}
			print(" ---- data ----")
			json_data = json.dumps(data).encode("utf-8")
			print(json_data)
			print("---------------")

			try:
				r = requests.post(url, data=json_data, headers=headers)
			except:
				print("err")

			GPIO.output(24,1)
			print("lamp on!")
			res = subprocess.run(["aplay", "--device=hw:1,0", "/home/pi/soiyaMix.wav"])
			GPIO.output(24,0)
			print("lamp off!")

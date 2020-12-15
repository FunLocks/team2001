import RPi.GPIO as GPIO
import time
import signal
import sys
import GPS 

def handler(signal, frame):
	GPIO.cleanup()
	print("GPIO cleanup...")
	print("exit!")
	sys.exit(0)

if __name__ == '__main__':
	signal.signal(signal.SIGINT, handler)
	
	GPIO.setmode(GPIO.BCM)
	GPIO.setup(23,GPIO.OUT) # switch

	gps = GPS.GPS()
	gps.start()

	flag = True;
	while True:
		#print(s.readline())
		sw_raw = GPIO.input(23)
		sw = 0
		if flag == True and sw_raw == 1:
			GPIO.input(23)
			sw = 1
			flag = False
		
		elif sw_raw == 0:
			flag = True

		if sw == 1:
			print('SW')
		

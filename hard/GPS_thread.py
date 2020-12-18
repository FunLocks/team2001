import serial
import threading
import time
from micropyGPS import micropyGPS

class GPS(threading.Thread):
	gps = micropyGPS.MicropyGPS(9, 'dd')
	port = '/dev/serial0'
	s = serial.Serial(port, 9600, timeout=10)
	s.readline()
	stack = []

	def __init__(self):
		super(GPS, self).__init__()

	def run(self):
		while True:
			string = self.s.readline().decode('utf-8')
			if string[0] != '$':
				continue
			for x in string:
				self.gps.update(x)
			
			print(self.gps.latitude)

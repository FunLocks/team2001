import serial
import time
from micropyGPS import micropyGPS

class GPS:
	gps = micropyGPS.MicropyGPS(9, 'dd')
	port = '/dev/serial0'
	s = serial.Serial(port, 9600, timeout=10)
	s.readline()

	def __init__(self):
		pass

	def update(self):
		string = self.s.readline().decode('utf-8')
		#if string[0] != '$':
		for x in string:
			self.gps.update(x)
		
	def latitude(self):
		return self.gps.latitude

	def longitude(self):
		return self.gps.longitude

	def latitude_to_string(self):
		return self.gps.latitude_string()
	
	def longitude_to_string(self):
		return self.gps.longitude_string()

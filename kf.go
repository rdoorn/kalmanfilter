package kalmanfilter

import (
	"math"
)

type Filter struct {
	R float64 // process noise // desirable
	Q float64 // measurement noise // estimated

	A float64 // state vector
	B float64 // control vector
	C float64 // measurement vector

	cov float64 //
	x   float64 // estimated signal without noise
}

func New(r, q float64) *Filter {
	return &Filter{
		R:   r,
		Q:   r,
		A:   1,
		B:   0,
		C:   1,
		cov: math.NaN(),
		x:   math.NaN(),
	}
}

func (f *Filter) Filter(m float64) float64 {

	u := float64(0)
	if math.IsNaN(f.x) {
		f.x = (1 / f.C) * m
		f.cov = (1 / f.C) * f.Q * (1 / f.C)
	} else {
		predX := (f.A * f.x) + (f.B * u)
		predCov := ((f.A * f.cov) * f.A) + f.R

		// Kalman Gain
		K := predCov * f.C * (1 / ((f.C * predCov * f.C) + f.Q))

		// Correction
		f.x = predX + K*(m-(f.C*predX))
		f.cov = predCov - (K * f.C * predCov)
	}

	return f.x

}

/*

class KalmanFilter:

    cov = float('nan')
    x = float('nan')

    def __init__(f, R, Q):
        """
        Constructor
        :param R: Process Noise
        :param Q: Measurement Noise
        """
        f.A = 1
        f.B = 0
        f.C = 1

        f.R = R
        f.Q = Q

    def filter(f, measurement):
        """
        Filters a measurement
        :param measurement: The measurement value to be filtered
        :return: The filtered value
        """
        u = 0
        if math.isnan(f.x):
            f.x = (1 / f.C) * measurement
            f.cov = (1 / f.C) * f.Q * (1 / f.C)
        else:
            predX = (f.A * f.x) + (f.B * u)
            predCov = ((f.A * f.cov) * f.A) + f.R

            # Kalman Gain
            K = predCov * f.C * (1 / ((f.C * predCov * f.C) + f.Q));

            # Correction
            f.x = predX + K * (measurement - (f.C * predX));
            f.cov = predCov - (K * f.C * predCov);

        return f.x

    def last_measurement(f):
        """
        Returns the last measurement fed into the filter
        :return: The last measurement fed into the filter
        """
        return f.x

    def set_measurement_noise(f, noise):
        """
        Sets measurement noise
        :param noise: The new measurement noise
        """
        f.Q = noise

    def set_process_noise(f, noise):
        """
        Sets process noise
        :param noise: The new process noise
        """
f.R = noise
*/

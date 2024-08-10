package length

func KtoM(k Kilometer) Mile {
	return Mile(k * 1.609)
}

func MtoK(m Mile) Kilometer {
	return Kilometer(m / 1.609)
}

package temperature

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KtoF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

func FtoK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

func KtoC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func CtoK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

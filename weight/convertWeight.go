package weight

func KgToP(k Killogram) Pound { return Pound(k * 2.2046) }

func PToKg(p Pound) Killogram { return Killogram(p / 2.2046) }

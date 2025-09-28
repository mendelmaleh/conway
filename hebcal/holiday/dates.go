package holiday

import "github.com/mendelmaleh/conway/hebcal"

var (
	// always two days
	RoshHashana = Event{
		Name:   "Rosh Hashana",
		Date:   hebcal.HebrewDate{Month: hebcal.Tishrei, Day: 1},
		Length: 2,
		Type:   Major | Holiday,
	}

	// regular fast
	TzomGedalia = Event{
		Name:  "Tzom Gedalia",
		Date:  hebcal.HebrewDate{Month: hebcal.Tishrei, Day: 3},
		Type:  Fast,
		Start: Dawn,
	}

	// always one day, is it really a holiday
	YomKippur = Event{
		Name: "Yom Kippur",
		Date: hebcal.HebrewDate{Month: hebcal.Tishrei, Day: 10},
		Type: Major | Holiday | Fast,
	}

	// one/two days + six/five chol hamoed
	Sukkot = Event{
		Name:   "Sukkot",
		Date:   hebcal.HebrewDate{Month: hebcal.Tishrei, Day: 15},
		Length: 7,
		Type:   Major | Holiday,
	}

	SheminiAtzeret = Event{
		Name: "Shemini Atzeret",
		Date: hebcal.HebrewDate{Month: hebcal.Tishrei, Day: 22},
		Type: Major | Holiday,
	}

	Chanukah = Event{
		Name:   "Chanukah",
		Date:   hebcal.HebrewDate{Month: hebcal.Kislev, Day: 25},
		Type:   Holiday,
		Length: 8,
	}

	TenthOfTevet = Event{
		Name:  "Asara Be'Tevet",
		Date:  hebcal.HebrewDate{Month: hebcal.Tevet, Day: 10},
		Type:  Fast,
		Start: Dawn,
	}

	FifteenOfShevat = Event{
		Name: "Tu Bi'Shevat",
		Date: hebcal.HebrewDate{Month: hebcal.Shevat, Day: 15},
		Type: Holiday,
	}

	TaanitEsther = Event{
		Name:  "Taanit Esther",
		Date:  hebcal.HebrewDate{Month: hebcal.AdarII, Day: 13},
		Type:  Fast,
		Start: Dawn,
	}

	Purim = Event{
		Name: "Purim",
		Date: hebcal.HebrewDate{Month: hebcal.AdarII, Day: 14},
		Type: Holiday,
	}

	ShushanPurim = Event{
		Name: "Shushan Purim",
		Date: hebcal.HebrewDate{Month: hebcal.AdarII, Day: 15},
		Type: Holiday,
	}

	TaanitBechorot = Event{
		Name:  "Taanit Bechorot",
		Date:  hebcal.HebrewDate{Month: hebcal.Nissan, Day: 14},
		Type:  Fast,
		Start: Dawn,
	}

	// one/two days + five/four chol hamoed + one/two days
	Pesach = Event{
		Name:   "Pesach",
		Date:   hebcal.HebrewDate{Month: hebcal.Nissan, Day: 15},
		Length: 7,
		Type:   Major | Holiday,
	}

	// one/two days
	Shavuot = Event{
		Name: "Shavuot",
		Date: hebcal.HebrewDate{Month: hebcal.Sivan, Day: 6},
		Type: Major | Holiday,
	}

	SeventeenOfTamuz = Event{
		Name: "Shiva Asar Be'Tamuz",
		Date: hebcal.HebrewDate{Month: hebcal.Tamuz, Day: 17},
		Type: Fast,
	}

	NinthOfAv = Event{
		Name: "Tisha Be'Av",
		Date: hebcal.HebrewDate{Month: hebcal.Av, Day: 9},
		Type: Fast,
	}

	All = []Event{
		// tishrei
		RoshHashana,
		TzomGedalia,
		YomKippur,
		Sukkot,
		SheminiAtzeret,
		// kislev, tevet, shevat
		Chanukah,
		TenthOfTevet,
		FifteenOfShevat,
		// adar // TODO: adar i/ii handling
		TaanitEsther,
		Purim,
		ShushanPurim, // TODO: start, diaspora?
		// nissan
		TaanitBechorot,
		Pesach,
		// PesachEnd, TODO: make last one/two major
		// sivan, tamuz, av
		Shavuot,
		SeventeenOfTamuz,
		NinthOfAv,
	}
)

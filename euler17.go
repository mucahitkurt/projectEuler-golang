package euler

const (
	one       = "ONE"
	two       = "TWO"
	three     = "THREE"
	four      = "FOUR"
	five      = "FIVE"
	six       = "SIX"
	seven     = "SEVEN"
	eight     = "EIGHT"
	nine      = "NINE"
	ten       = "TEN"
	eleven    = "ELEVEN"
	twelve    = "TWELVE"
	thirteen  = "THIRTEEN"
	fourteen  = "FOURTEEN"
	fifteen   = "FIFTEEN"
	sixteen   = "SIXTEEN"
	seventeen = "SEVENTEEN"
	eighteen  = "EIGHTEEN"
	nineteen  = "NINETEEN"

	twenty  = "TWENTY"
	thirty  = "THIRTY"
	fourty  = "FORTY"
	fifty   = "FIFTY"
	sixty   = "SIXTY"
	seventy = "SEVENTY"
	eighty  = "EIGHTY"
	ninety  = "NINETY"

	hundred = "HUNDRED"

	thousand = "THOUSAND"
)

var ones []string = []string{one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen}
var tens []string = []string{twenty, thirty, fourty, fifty, sixty, seventy, eighty, ninety}

func Euler17() int {

	/*sum := 0
	for i := 1; i <= 1000; i++ {
		strArr := readNum(i)
		//fmt.Println(strArr)
		for _, str := range strArr {
			sum += len(str)
		}
	}

	return sum*/return letterCount()
}

func readNum(i int) []string {
	var res []string

	rem := i / 1000
	if rem > 0 {
		res = append(res, ones[rem-1])
		res = append(res, thousand)
	}
	i = i % 1000

	rem = i / 100
	if rem > 0 {
		res = append(res, ones[rem-1])
		res = append(res, hundred)
	}
	i = i % 100

	if i > 0 {
		if rem > 0 {
			res = append(res, "AND")
		}
		if i >= 20 {
			res = append(res, tens[(i/10)-2])
			rem = i % 10
			if rem > 0 {
				res = append(res, ones[(rem-1)])
			}
		} else {
			res = append(res, ones[(i-1)])
		}
	}

	return res
}

//alternative approach
func letterCount() int {
	onetonine := len("onetwothreefourfivesixseveneightnine")
	tentonineteen := len("teneleventwelvethirteenfourteenfifteensixteenseventeeneighteennineteen")
	and := len("and")
	twentytoninety := len("twentythirtyfortyfiftysixtyseventyeightyninety")
	hundred := len("hundred")
	thousand := len("thousand")
	count := len("one") + thousand + 900*hundred + 100*onetonine + 100*twentytoninety + 891*and + 80*onetonine + 10*(onetonine+tentonineteen)

	return count
}

package pointer

type Address1 struct {
	City, Province, Country string
}

func ChangeCountryIndonesia(address *Address1) {
	address.Country = "Indonesia"
}

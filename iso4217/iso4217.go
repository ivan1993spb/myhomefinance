package iso4217

const (
	defaultCode = "UNKNOWN"
	defaultName = "Unknown"
)

const (
	defaultCurrency Currency = iota
	AED
	AFN
	ALL
	AMD
	ANG
	AOA
	ARS
	AUD
	AWG
	AZN
	BAM
	BBD
	BDT
	BGN
	BHD
	BIF
	BMD
	BND
	BOB
	BRL
	BSD
	BTN
	BWP
	BYN
	BZD
	CAD
	CDF
	CHF
	CLP
	CNY
	COP
	CRC
	CUC
	CUP
	CVE
	CZK
	DJF
	DKK
	DOP
	DZD
	EGP
	ERN
	ETB
	EUR
	FJD
	FKP
	GBP
	GEL
	GGP
	GHS
	GIP
	GMD
	GNF
	GTQ
	GYD
	HKD
	HNL
	HRK
	HTG
	HUF
	IDR
	ILS
	IMP
	INR
	IQD
	IRR
	ISK
	JEP
	JMD
	JOD
	JPY
	KES
	KGS
	KHR
	KMF
	KPW
	KRW
	KWD
	KYD
	KZT
	LAK
	LBP
	LKR
	LRD
	LSL
	LYD
	MAD
	MDL
	MGA
	MKD
	MMK
	MNT
	MOP
	MRO
	MUR
	MVR
	MWK
	MXN
	MYR
	MZN
	NAD
	NGN
	NIO
	NOK
	NPR
	NZD
	OMR
	PAB
	PEN
	PGK
	PHP
	PKR
	PLN
	PYG
	QAR
	RON
	RSD
	RUB
	RWF
	SAR
	SBD
	SCR
	SDG
	SEK
	SGD
	SHP
	SLL
	SOS
	SPL
	SRD
	STD
	SVC
	SYP
	SZL
	THB
	TJS
	TMT
	TND
	TOP
	TRY
	TTD
	TVD
	TWD
	TZS
	UAH
	UGX
	USD
	UYU
	UZS
	VEF
	VND
	VUV
	WST
	XAF
	XCD
	XDR
	XOF
	XPF
	YER
	ZAR
	ZMW
	ZWD
)

var codes = map[Currency]string{
	AED: "AED",
	AFN: "AFN",
	ALL: "ALL",
	AMD: "AMD",
	ANG: "ANG",
	AOA: "AOA",
	ARS: "ARS",
	AUD: "AUD",
	AWG: "AWG",
	AZN: "AZN",
	BAM: "BAM",
	BBD: "BBD",
	BDT: "BDT",
	BGN: "BGN",
	BHD: "BHD",
	BIF: "BIF",
	BMD: "BMD",
	BND: "BND",
	BOB: "BOB",
	BRL: "BRL",
	BSD: "BSD",
	BTN: "BTN",
	BWP: "BWP",
	BYN: "BYN",
	BZD: "BZD",
	CAD: "CAD",
	CDF: "CDF",
	CHF: "CHF",
	CLP: "CLP",
	CNY: "CNY",
	COP: "COP",
	CRC: "CRC",
	CUC: "CUC",
	CUP: "CUP",
	CVE: "CVE",
	CZK: "CZK",
	DJF: "DJF",
	DKK: "DKK",
	DOP: "DOP",
	DZD: "DZD",
	EGP: "EGP",
	ERN: "ERN",
	ETB: "ETB",
	EUR: "EUR",
	FJD: "FJD",
	FKP: "FKP",
	GBP: "GBP",
	GEL: "GEL",
	GGP: "GGP",
	GHS: "GHS",
	GIP: "GIP",
	GMD: "GMD",
	GNF: "GNF",
	GTQ: "GTQ",
	GYD: "GYD",
	HKD: "HKD",
	HNL: "HNL",
	HRK: "HRK",
	HTG: "HTG",
	HUF: "HUF",
	IDR: "IDR",
	ILS: "ILS",
	IMP: "IMP",
	INR: "INR",
	IQD: "IQD",
	IRR: "IRR",
	ISK: "ISK",
	JEP: "JEP",
	JMD: "JMD",
	JOD: "JOD",
	JPY: "JPY",
	KES: "KES",
	KGS: "KGS",
	KHR: "KHR",
	KMF: "KMF",
	KPW: "KPW",
	KRW: "KRW",
	KWD: "KWD",
	KYD: "KYD",
	KZT: "KZT",
	LAK: "LAK",
	LBP: "LBP",
	LKR: "LKR",
	LRD: "LRD",
	LSL: "LSL",
	LYD: "LYD",
	MAD: "MAD",
	MDL: "MDL",
	MGA: "MGA",
	MKD: "MKD",
	MMK: "MMK",
	MNT: "MNT",
	MOP: "MOP",
	MRO: "MRO",
	MUR: "MUR",
	MVR: "MVR",
	MWK: "MWK",
	MXN: "MXN",
	MYR: "MYR",
	MZN: "MZN",
	NAD: "NAD",
	NGN: "NGN",
	NIO: "NIO",
	NOK: "NOK",
	NPR: "NPR",
	NZD: "NZD",
	OMR: "OMR",
	PAB: "PAB",
	PEN: "PEN",
	PGK: "PGK",
	PHP: "PHP",
	PKR: "PKR",
	PLN: "PLN",
	PYG: "PYG",
	QAR: "QAR",
	RON: "RON",
	RSD: "RSD",
	RUB: "RUB",
	RWF: "RWF",
	SAR: "SAR",
	SBD: "SBD",
	SCR: "SCR",
	SDG: "SDG",
	SEK: "SEK",
	SGD: "SGD",
	SHP: "SHP",
	SLL: "SLL",
	SOS: "SOS",
	SPL: "SPL",
	SRD: "SRD",
	STD: "STD",
	SVC: "SVC",
	SYP: "SYP",
	SZL: "SZL",
	THB: "THB",
	TJS: "TJS",
	TMT: "TMT",
	TND: "TND",
	TOP: "TOP",
	TRY: "TRY",
	TTD: "TTD",
	TVD: "TVD",
	TWD: "TWD",
	TZS: "TZS",
	UAH: "UAH",
	UGX: "UGX",
	USD: "USD",
	UYU: "UYU",
	UZS: "UZS",
	VEF: "VEF",
	VND: "VND",
	VUV: "VUV",
	WST: "WST",
	XAF: "XAF",
	XCD: "XCD",
	XDR: "XDR",
	XOF: "XOF",
	XPF: "XPF",
	YER: "YER",
	ZAR: "ZAR",
	ZMW: "ZMW",
	ZWD: "ZWD",
}

var names = map[Currency]string{
	AED: "United Arab Emirates Dirham",
	AFN: "Afghanistan Afghani",
	ALL: "Albania Lek",
	AMD: "Armenia Dram",
	ANG: "Netherlands Antilles Guilder",
	AOA: "Angola Kwanza",
	ARS: "Argentina Peso",
	AUD: "Australia Dollar",
	AWG: "Aruba Guilder",
	AZN: "Azerbaijan New Manat",
	BAM: "Bosnia and Herzegovina Convertible Marka",
	BBD: "Barbados Dollar",
	BDT: "Bangladesh Taka",
	BGN: "Bulgaria Lev",
	BHD: "Bahrain Dinar",
	BIF: "Burundi Franc",
	BMD: "Bermuda Dollar",
	BND: "Brunei Darussalam Dollar",
	BOB: "Bolivia Bolíviano",
	BRL: "Brazil Real",
	BSD: "Bahamas Dollar",
	BTN: "Bhutan Ngultrum",
	BWP: "Botswana Pula",
	BYN: "Belarus Ruble",
	BZD: "Belize Dollar",
	CAD: "Canada Dollar",
	CDF: "Congo/Kinshasa Franc",
	CHF: "Switzerland Franc",
	CLP: "Chile Peso",
	CNY: "China Yuan Renminbi",
	COP: "Colombia Peso",
	CRC: "Costa Rica Colon",
	CUC: "Cuba Convertible Peso",
	CUP: "Cuba Peso",
	CVE: "Cape Verde Escudo",
	CZK: "Czech Republic Koruna",
	DJF: "Djibouti Franc",
	DKK: "Denmark Krone",
	DOP: "Dominican Republic Peso",
	DZD: "Algeria Dinar",
	EGP: "Egypt Pound",
	ERN: "Eritrea Nakfa",
	ETB: "Ethiopia Birr",
	EUR: "Euro Member Countries",
	FJD: "Fiji Dollar",
	FKP: "Falkland Islands (Malvinas) Pound",
	GBP: "United Kingdom Pound",
	GEL: "Georgia Lari",
	GGP: "Guernsey Pound",
	GHS: "Ghana Cedi",
	GIP: "Gibraltar Pound",
	GMD: "Gambia Dalasi",
	GNF: "Guinea Franc",
	GTQ: "Guatemala Quetzal",
	GYD: "Guyana Dollar",
	HKD: "Hong Kong Dollar",
	HNL: "Honduras Lempira",
	HRK: "Croatia Kuna",
	HTG: "Haiti Gourde",
	HUF: "Hungary Forint",
	IDR: "Indonesia Rupiah",
	ILS: "Israel Shekel",
	IMP: "Isle of Man Pound",
	INR: "India Rupee",
	IQD: "Iraq Dinar",
	IRR: "Iran Rial",
	ISK: "Iceland Krona",
	JEP: "Jersey Pound",
	JMD: "Jamaica Dollar",
	JOD: "Jordan Dinar",
	JPY: "Japan Yen",
	KES: "Kenya Shilling",
	KGS: "Kyrgyzstan Som",
	KHR: "Cambodia Riel",
	KMF: "Comoros Franc",
	KPW: "Korea (North) Won",
	KRW: "Korea (South) Won",
	KWD: "Kuwait Dinar",
	KYD: "Cayman Islands Dollar",
	KZT: "Kazakhstan Tenge",
	LAK: "Laos Kip",
	LBP: "Lebanon Pound",
	LKR: "Sri Lanka Rupee",
	LRD: "Liberia Dollar",
	LSL: "Lesotho Loti",
	LYD: "Libya Dinar",
	MAD: "Morocco Dirham",
	MDL: "Moldova Leu",
	MGA: "Madagascar Ariary",
	MKD: "Macedonia Denar",
	MMK: "Myanmar (Burma) Kyat",
	MNT: "Mongolia Tughrik",
	MOP: "Macau Pataca",
	MRO: "Mauritania Ouguiya",
	MUR: "Mauritius Rupee",
	MVR: "Maldives (Maldive Islands) Rufiyaa",
	MWK: "Malawi Kwacha",
	MXN: "Mexico Peso",
	MYR: "Malaysia Ringgit",
	MZN: "Mozambique Metical",
	NAD: "Namibia Dollar",
	NGN: "Nigeria Naira",
	NIO: "Nicaragua Cordoba",
	NOK: "Norway Krone",
	NPR: "Nepal Rupee",
	NZD: "New Zealand Dollar",
	OMR: "Oman Rial",
	PAB: "Panama Balboa",
	PEN: "Peru Sol",
	PGK: "Papua New Guinea Kina",
	PHP: "Philippines Peso",
	PKR: "Pakistan Rupee",
	PLN: "Poland Zloty",
	PYG: "Paraguay Guarani",
	QAR: "Qatar Riyal",
	RON: "Romania New Leu",
	RSD: "Serbia Dinar",
	RUB: "Russia Ruble",
	RWF: "Rwanda Franc",
	SAR: "Saudi Arabia Riyal",
	SBD: "Solomon Islands Dollar",
	SCR: "Seychelles Rupee",
	SDG: "Sudan Pound",
	SEK: "Sweden Krona",
	SGD: "Singapore Dollar",
	SHP: "Saint Helena Pound",
	SLL: "Sierra Leone Leone",
	SOS: "Somalia Shilling",
	SPL: "Seborga Luigino",
	SRD: "Suriname Dollar",
	STD: "São Tomé and Príncipe Dobra",
	SVC: "El Salvador Colon",
	SYP: "Syria Pound",
	SZL: "Swaziland Lilangeni",
	THB: "Thailand Baht",
	TJS: "Tajikistan Somoni",
	TMT: "Turkmenistan Manat",
	TND: "Tunisia Dinar",
	TOP: "Tonga Pa'anga",
	TRY: "Turkey Lira",
	TTD: "Trinidad and Tobago Dollar",
	TVD: "Tuvalu Dollar",
	TWD: "Taiwan New Dollar",
	TZS: "Tanzania Shilling",
	UAH: "Ukraine Hryvnia",
	UGX: "Uganda Shilling",
	USD: "United States Dollar",
	UYU: "Uruguay Peso",
	UZS: "Uzbekistan Som",
	VEF: "Venezuela Bolivar",
	VND: "Viet Nam Dong",
	VUV: "Vanuatu Vatu",
	WST: "Samoa Tala",
	XAF: "Communauté Financière Africaine (BEAC) CFA Franc BEAC",
	XCD: "East Caribbean Dollar",
	XDR: "International Monetary Fund (IMF) Special Drawing Rights",
	XOF: "Communauté Financière Africaine (BCEAO) Franc",
	XPF: "Comptoirs Français du Pacifique (CFP) Franc",
	YER: "Yemen Rial",
	ZAR: "South Africa Rand",
	ZMW: "Zambia Kwacha",
	ZWD: "Zimbabwe Dollar",
}

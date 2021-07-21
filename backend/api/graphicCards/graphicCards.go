package graphiccards

type GraphicCard struct {
	Id				int
	Img				string
	Name			string
	Manufacturer	string
	Brand			string
	RamType 		string
	VideoOutput		string
	Price 			float64
}

var Card = []GraphicCard{
	{
		Id: 0,
		Img: "https://thumb.pccomponentes.com/w-530-530/articles/36/364293/1107-gigabyte-amd-radeon-rx-6700-xt-gaming-oc-12gb-gddr6.jpg",
		Name: "AMD Radeon RX 6700 XT",
		Manufacturer: "AMD",
		Brand:	"PowerColor",
		RamType: "12GB GDDR6",
		VideoOutput: "HDMI 2.1, DisplayPort",
		Price: 904.60,
	},
	{
		Id: 1,
		Img: "https://m.media-amazon.com/images/I/71FHpthlS3L._AC_SL1500_.jpg",
		Name: "MSI Gaming GeForce RTX 3070 Gaming X Trio",
		Manufacturer: "‎MSI COMPUTER",
		Brand:	"MSI",
		RamType: "8GB GDDR6",
		VideoOutput: "HDMI 2.1, DisplayPort",
		Price: 800.00,
	},
	{
		Id: 2,
		Img: "https://m.media-amazon.com/images/I/617763dVzSL._AC_SL1200_.jpg",
		Name: "Asrock RX6800XT PGD 16GO",
		Manufacturer: "Asrock",
		Brand:	"AsRock",
		RamType: "61GB GDDR6",
		VideoOutput: "HDMI",
		Price: 1280.40,
	},
	{
		Id: 3,
		Img: "https://images-na.ssl-images-amazon.com/images/I/81-70VBUexL._AC_SL1500_.jpg",
		Name: "XFX Speedster MERC319 AMD Radeon RX 6800 XT Core",
		Manufacturer: "XFX",
		Brand:	"XFX",
		RamType: "16GB GDDR6",
		VideoOutput: "HDMI 3",
		Price: 1550.00,
	},
}
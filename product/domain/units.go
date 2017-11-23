package domain

// Unit describes a unit on an attribute
type Unit struct {
	Code   string
	Symbol string
}

// Constants of the unit codes
// See also the PIM unit codes on which these are based (and no other will get into this system)
const (
	// Area Units
	SQUARE_MILLIMETER = "SQUARE_MILLIMETER"
	SQUARE_CENTIMETER = "SQUARE_CENTIMETER"
	SQUARE_DECIMETER  = "SQUARE_DECIMETER"
	SQUARE_METER      = "SQUARE_METER"
	CENTIARE          = "CENTIARE"
	SQUARE_DEKAMETER  = "SQUARE_DEKAMETER"
	ARE               = "ARE"
	SQUARE_HECTOMETER = "SQUARE_HECTOMETER"
	HECTARE           = "HECTARE"
	SQUARE_KILOMETER  = "SQUARE_KILOMETER"
	SQUARE_MIL        = "SQUARE_MIL"
	SQUARE_INCH       = "SQUARE_INCH"
	SQUARE_FOOT       = "SQUARE_FOOT"
	SQUARE_YARD       = "SQUARE_YARD"
	ARPENT            = "ARPENT"
	ACRE              = "ACRE"
	SQUARE_FURLONG    = "SQUARE_FURLONG"
	SQUARE_MILE       = "SQUARE_MILE"

	// Binary Units
	BIT      = "BIT"
	BYTE     = "BYTE"
	KILOBYTE = "KILOBYTE"
	MEGABYTE = "MEGABYTE"
	GIGABYTE = "GIGABYTE"
	TERABYTE = "TERABYTE"

	// Loudness Units
	DECIBEL = "DECIBEL"

	// Frequency Units
	HERTZ     = "HERTZ"
	KILOHERTZ = "KILOHERTZ"
	MEGAHERTZ = "MEGAHERTZ"
	GIGAHERTZ = "GIGAHERTZ"
	TERAHERTZ = "TERAHERTZ"

	// Length Unit
	MILLIMETER = "MILLIMETER"
	CENTIMETER = "CENTIMETER"
	DECIMETER  = "DECIMETER"
	METER      = "METER"
	DEKAMETER  = "DEKAMETER"
	HECTOMETER = "HECTOMETER"
	KILOMETER  = "KILOMETER"
	MIL        = "MIL"
	INCH       = "INCH"
	FEET       = "FEET"
	YARD       = "YARD"
	CHAIN      = "CHAIN"
	FURLONG    = "FURLONG"
	MILE       = "MILE"

	// Power Units
	WATT     = "WATT"
	KILOWATT = "KILOWATT"
	MEGAWATT = "MEGAWATT"
	GIGAWATT = "GIGAWATT"
	TERAWATT = "TERAWATT"

	// Voltage Units
	MILLIVOLT = "MILLIVOLT"
	CENTIVOLT = "CENTIVOLT"
	DECIVOLT  = "DECIVOLT"
	VOLT      = "VOLT"
	DEKAVOLT  = "DEKAVOLT"
	HECTOVOLT = "HECTOVOLT"
	KILOVOLT  = "KILOVOLT"

	// Intensity
	MILLIAMPERE = "MILLIAMPERE"
	CENTIAMPERE = "CENTIAMPERE"
	DECIAMPERE  = "DECIAMPERE"
	AMPERE      = "AMPERE"
	DEKAMPERE   = "DEKAMPERE"
	HECTOAMPERE = "HECTOAMPERE"
	KILOAMPERE  = "KILOAMPERE"

	// Resistance
	MILLIOHM = "MILLIOHM"
	CENTIOHM = "CENTIOHM"
	DECIOHM  = "DECIOHM"
	OHM      = "OHM"
	DEKAOHM  = "DEKAOHM"
	HECTOHM  = "HECTOHM"
	KILOHM   = "KILOHM"
	MEGOHM   = "MEGOHM"

	// Speed
	METER_PER_SECOND   = "METER_PER_SECOND"
	METER_PER_MINUTE   = "METER_PER_MINUTE"
	METER_PER_HOUR     = "METER_PER_HOUR"
	KILOMETER_PER_HOUR = "KILOMETER_PER_HOUR"
	FOOT_PER_SECOND    = "FOOT_PER_SECOND"
	FOOT_PER_HOUR      = "FOOT_PER_HOUR"
	YARD_PER_HOUR      = "YARD_PER_HOUR"
	MILE_PER_HOUR      = "MILE_PER_HOUR"

	// Electric Charge
	MILLIAMPEREHOUR = "MILLIAMPEREHOUR"
	AMPEREHOUR      = "AMPEREHOUR"
	MILLICOULOMB    = "MILLICOULOMB"
	CENTICOULOMB    = "CENTICOULOMB"
	DECICOULOMB     = "DECICOULOMB"
	COULOMB         = "COULOMB"
	DEKACOULOMB     = "DEKACOULOMB"
	HECTOCOULOMB    = "HECTOCOULOMB"
	KILOCOULOMB     = "KILOCOULOMB"

	// Duration
	MILLISECOND = "MILLISECOND"
	SECOND      = "SECOND"
	MINUTE      = "MINUTE"
	HOUR        = "HOUR"
	DAY         = "DAY"
	WEEK        = "WEEK"
	MONTH       = "MONTH"
	YEAR        = "YEAR"

	// Temperature
	CELSIUS    = "CELSIUS"
	FAHRENHEIT = "FAHRENHEIT"
	KELVIN     = "KELVIN"
	RANKINE    = "RANKINE"
	REAUMUR    = "REAUMUR"

	// Volume Units
	CUBIC_MILLIMETER = "CUBIC_MILLIMETER"
	CUBIC_CENTIMETER = "CUBIC_CENTIMETER"
	MILLILITER       = "MILLILITER"
	CENTILITER       = "CENTILITER"
	DECILITER        = "DECILITER"
	CUBIC_DECIMETER  = "CUBIC_DECIMETER"
	LITER            = "LITER"
	CUBIC_METER      = "CUBIC_METER"
	OUNCE            = "OUNCE"
	PINT             = "PINT"
	BARREL           = "BARREL"
	GALLON           = "GALLON"
	CUBIC_FOOT       = "CUBIC_FOOT"
	CUBIC_INCH       = "CUBIC_INCH"
	CUBIC_YARD       = "CUBIC_YARD"

	// Weigt Units
	MILLIGRAM = "MILLIGRAM"
	GRAM      = "GRAM"
	KILOGRAM  = "KILOGRAM"
	TON       = "TON"
	GRAIN     = "GRAIN"
	DENIER    = "DENIER"
	POUND     = "POUND"
	MARC      = "MARC"
	LIVRE     = "LIVRE"

	// Pressure
	BAR         = "BAR"
	PASCAL      = "PASCAL"
	HECTOPASCAL = "HECTOPASCAL"
	MILLIBAR    = "MILLIBAR"
	ATM         = "ATM"
	PSI         = "PSI"
	TORR        = "TORR"
	MMHG        = "MMHG"

	// Energy
	JOULE       = "JOULE"
	CALORIE     = "CALORIE"
	KILOCALORIE = "KILOCALORIE"
	KILOJOULE   = "KILOJOULE"

	// Piece / CaseBox Units
	PCS   = "PCS"
	PIECE = "PIECE"
	DOZEN = "DOZEN"
)

// Units provides the unit map
var Units = map[string]Unit{
	// Area Units
	SQUARE_MILLIMETER: {
		Code:   SQUARE_MILLIMETER,
		Symbol: "mm²",
	},

	SQUARE_CENTIMETER: {
		Code:   SQUARE_CENTIMETER,
		Symbol: "cm²",
	},

	SQUARE_DECIMETER: {
		Code:   SQUARE_DECIMETER,
		Symbol: "dm²",
	},

	SQUARE_METER: {
		Code:   SQUARE_METER,
		Symbol: "m²",
	},

	CENTIARE: {
		Code:   CENTIARE,
		Symbol: "ca",
	},

	SQUARE_DEKAMETER: {
		Code:   SQUARE_DEKAMETER,
		Symbol: "dam²",
	},

	ARE: {
		Code:   ARE,
		Symbol: "a",
	},

	SQUARE_HECTOMETER: {
		Code:   SQUARE_HECTOMETER,
		Symbol: "hm²",
	},

	HECTARE: {
		Code:   HECTARE,
		Symbol: "ha",
	},

	SQUARE_KILOMETER: {
		Code:   SQUARE_KILOMETER,
		Symbol: "km²",
	},

	SQUARE_MIL: {
		Code:   "SQUARE_MIL",
		Symbol: "sq mil",
	},

	SQUARE_INCH: {
		Code:   "SQUARE_INCH",
		Symbol: "in²",
	},

	SQUARE_FOOT: {
		Code:   "SQUARE_FOOT",
		Symbol: "ft²",
	},

	SQUARE_YARD: {
		Code:   "SQUARE_YARD",
		Symbol: "yd²",
	},

	ARPENT: {
		Code:   "ARPENT",
		Symbol: "arpent",
	},

	ACRE: {
		Code:   "ACRE",
		Symbol: "A",
	},

	SQUARE_FURLONG: {
		Code:   "SQUARE_FURLONG",
		Symbol: "fur²",
	},

	SQUARE_MILE: {
		Code:   "SQUARE_MILE",
		Symbol: "mi²",
	},

	// Binary Units
	BIT: {
		Code:   BIT,
		Symbol: "b",
	},
	BYTE: {
		Code:   BYTE,
		Symbol: "B",
	},
	KILOBYTE: {
		Code:   KILOBYTE,
		Symbol: "kB",
	},
	MEGABYTE: {
		Code:   MEGABYTE,
		Symbol: "MB",
	},
	GIGABYTE: {
		Code:   GIGABYTE,
		Symbol: "GB",
	},
	TERABYTE: {
		Code:   TERABYTE,
		Symbol: "TB",
	},

	// Loudness Units
	DECIBEL: {
		Code:   DECIBEL,
		Symbol: "dB",
	},

	// Frequency Units
	HERTZ: {
		Code:   HERTZ,
		Symbol: "Hz",
	},
	KILOHERTZ: {
		Code:   KILOHERTZ,
		Symbol: "kHz",
	},
	MEGAHERTZ: {
		Code:   MEGAHERTZ,
		Symbol: "MHz",
	},
	GIGAHERTZ: {
		Code:   GIGAHERTZ,
		Symbol: "GHz",
	},
	TERAHERTZ: {
		Code:   TERAHERTZ,
		Symbol: "THz",
	},

	// Length Units
	MILLIMETER: {
		Code:   MILLIMETER,
		Symbol: "mm",
	},
	CENTIMETER: {
		Code:   CENTIMETER,
		Symbol: "cm",
	},
	DECIMETER: {
		Code:   DECIMETER,
		Symbol: "dm",
	},
	METER: {
		Code:   METER,
		Symbol: "m",
	},
	DEKAMETER: {
		Code:   DEKAMETER,
		Symbol: "dam",
	},
	HECTOMETER: {
		Code:   HECTOMETER,
		Symbol: "hm",
	},
	KILOMETER: {
		Code:   KILOMETER,
		Symbol: "km",
	},
	MIL: {
		Code:   MIL,
		Symbol: "mil",
	},
	INCH: {
		Code:   INCH,
		Symbol: "in",
	},
	FEET: {
		Code:   FEET,
		Symbol: "ft",
	},
	YARD: {
		Code:   YARD,
		Symbol: "yd",
	},
	CHAIN: {
		Code:   CHAIN,
		Symbol: "ch",
	},
	FURLONG: {
		Code:   FURLONG,
		Symbol: "fur",
	},
	MILE: {
		Code:   MILE,
		Symbol: "mi",
	},

	// Power Units
	WATT: {
		Code:   WATT,
		Symbol: "W",
	},
	KILOWATT: {
		Code:   KILOWATT,
		Symbol: "mW",
	},
	MEGAWATT: {
		Code:   MEGAWATT,
		Symbol: "MW",
	},
	GIGAWATT: {
		Code:   GIGAWATT,
		Symbol: "GW",
	},
	TERAWATT: {
		Code:   TERAWATT,
		Symbol: "TW",
	},

	// Voltage Units
	MILLIVOLT: {
		Code:   MILLIVOLT,
		Symbol: "mV",
	},
	CENTIVOLT: {
		Code:   CENTIVOLT,
		Symbol: "cV",
	},
	DECIVOLT: {
		Code:   DECIVOLT,
		Symbol: "dV",
	},
	VOLT: {
		Code:   VOLT,
		Symbol: "V",
	},
	DEKAVOLT: {
		Code:   DEKAVOLT,
		Symbol: "daV",
	},
	HECTOVOLT: {
		Code:   HECTOVOLT,
		Symbol: "hV",
	},
	KILOVOLT: {
		Code:   KILOVOLT,
		Symbol: "kV",
	},

	// Intensity
	MILLIAMPERE: {
		Code:   MILLIAMPERE,
		Symbol: "mA",
	},
	CENTIAMPERE: {
		Code:   CENTIAMPERE,
		Symbol: "cA",
	},
	DECIAMPERE: {
		Code:   DECIAMPERE,
		Symbol: "dA",
	},
	AMPERE: {
		Code:   AMPERE,
		Symbol: "A",
	},
	DEKAMPERE: {
		Code:   "DEKAMPERE",
		Symbol: "daA",
	},
	HECTOAMPERE: {
		Code:   "HECTOAMPERE",
		Symbol: "hA",
	},
	KILOAMPERE: {
		Code:   "KILOAMPERE",
		Symbol: "kA",
	},

	// Resistance
	MILLIOHM: {
		Code:   MILLIOHM,
		Symbol: "mΩ",
	},
	CENTIOHM: {
		Code:   CENTIOHM,
		Symbol: "cΩ",
	},
	DECIOHM: {
		Code:   DECIOHM,
		Symbol: "dΩ",
	},
	OHM: {
		Code:   OHM,
		Symbol: "Ω",
	},
	DEKAOHM: {
		Code:   DEKAOHM,
		Symbol: "daΩ",
	},
	HECTOHM: {
		Code:   HECTOHM,
		Symbol: "hΩ",
	},
	KILOHM: {
		Code:   KILOHM,
		Symbol: "kΩ",
	},
	MEGOHM: {
		Code:   MEGOHM,
		Symbol: "MΩ",
	},

	// Speed
	METER_PER_SECOND: {
		Code:   METER_PER_SECOND,
		Symbol: "mdivs",
	},
	METER_PER_MINUTE: {
		Code:   METER_PER_MINUTE,
		Symbol: "mdivm",
	},
	METER_PER_HOUR: {
		Code:   METER_PER_HOUR,
		Symbol: "mdivh",
	},
	KILOMETER_PER_HOUR: {
		Code:   KILOMETER_PER_HOUR,
		Symbol: "kmdivh",
	},
	FOOT_PER_SECOND: {
		Code:   FOOT_PER_SECOND,
		Symbol: "ftdivs",
	},
	FOOT_PER_HOUR: {
		Code:   FOOT_PER_HOUR,
		Symbol: "ftdivh",
	},
	YARD_PER_HOUR: {
		Code:   YARD_PER_HOUR,
		Symbol: "yddivh",
	},
	MILE_PER_HOUR: {
		Code:   MILE_PER_HOUR,
		Symbol: "midivh",
	},

	// Electric Charge
	MILLIAMPEREHOUR: {
		Code:   MILLIAMPEREHOUR,
		Symbol: "mAh",
	},
	AMPEREHOUR: {
		Code:   AMPEREHOUR,
		Symbol: "Ah",
	},
	MILLICOULOMB: {
		Code:   MILLICOULOMB,
		Symbol: "mC",
	},
	CENTICOULOMB: {
		Code:   CENTICOULOMB,
		Symbol: "cC",
	},
	DECICOULOMB: {
		Code:   DECICOULOMB,
		Symbol: "dC",
	},
	COULOMB: {
		Code:   COULOMB,
		Symbol: "C",
	},
	DEKACOULOMB: {
		Code:   DEKACOULOMB,
		Symbol: "daC",
	},
	HECTOCOULOMB: {
		Code:   HECTOCOULOMB,
		Symbol: "hC",
	},
	KILOCOULOMB: {
		Code:   KILOCOULOMB,
		Symbol: "kC",
	},

	// Duration
	MILLISECOND: {
		Code:   MILLISECOND,
		Symbol: "ms",
	},
	SECOND: {
		Code:   SECOND,
		Symbol: "s",
	},
	MINUTE: {
		Code:   MINUTE,
		Symbol: "m",
	},
	HOUR: {
		Code:   HOUR,
		Symbol: "h",
	},
	DAY: {
		Code:   DAY,
		Symbol: "d",
	},
	WEEK: {
		Code:   WEEK,
		Symbol: "week",
	},
	MONTH: {
		Code:   MONTH,
		Symbol: "month",
	},
	YEAR: {
		Code:   YEAR,
		Symbol: "year",
	},

	// Temperature
	CELSIUS: {
		Code:   CELSIUS,
		Symbol: "°C",
	},
	FAHRENHEIT: {
		Code:   FAHRENHEIT,
		Symbol: "°F",
	},
	KELVIN: {
		Code:   KELVIN,
		Symbol: "°K",
	},
	RANKINE: {
		Code:   RANKINE,
		Symbol: "°R",
	},
	REAUMUR: {
		Code:   REAUMUR,
		Symbol: "°r",
	},

	// Volume Units
	CUBIC_MILLIMETER: {
		Code:   CUBIC_MILLIMETER,
		Symbol: "mm³",
	},
	CUBIC_CENTIMETER: {
		Code:   CUBIC_CENTIMETER,
		Symbol: "cm³",
	},
	MILLILITER: {
		Code:   MILLILITER,
		Symbol: "ml",
	},
	CENTILITER: {
		Code:   CENTILITER,
		Symbol: "cl",
	},
	DECILITER: {
		Code:   DECILITER,
		Symbol: "dl",
	},
	CUBIC_DECIMETER: {
		Code:   CUBIC_DECIMETER,
		Symbol: "dm³",
	},
	LITER: {
		Code:   LITER,
		Symbol: "l",
	},
	CUBIC_METER: {
		Code:   CUBIC_METER,
		Symbol: "m³",
	},
	OUNCE: {
		Code:   OUNCE,
		Symbol: "oz",
	},
	PINT: {
		Code:   PINT,
		Symbol: "pt",
	},
	BARREL: {
		Code:   BARREL,
		Symbol: "bbl",
	},
	GALLON: {
		Code:   GALLON,
		Symbol: "gal",
	},
	CUBIC_FOOT: {
		Code:   CUBIC_FOOT,
		Symbol: "ft³",
	},
	CUBIC_INCH: {
		Code:   CUBIC_INCH,
		Symbol: "in³",
	},
	CUBIC_YARD: {
		Code:   CUBIC_YARD,
		Symbol: "yd³",
	},

	// Weight Units
	MILLIGRAM: {
		Code:   MILLIGRAM,
		Symbol: "mg",
	},
	GRAM: {
		Code:   GRAM,
		Symbol: "g",
	},
	KILOGRAM: {
		Code:   KILOGRAM,
		Symbol: "kg",
	},
	TON: {
		Code:   TON,
		Symbol: "t",
	},
	GRAIN: {
		Code:   GRAIN,
		Symbol: "gr",
	},
	DENIER: {
		Code:   DENIER,
		Symbol: "dernier",
	},
	POUND: {
		Code:   POUND,
		Symbol: "lb",
	},
	MARC: {
		Code:   MARC,
		Symbol: "marc",
	},
	LIVRE: {
		Code:   LIVRE,
		Symbol: "livre",
	},

	// Pressure
	BAR: {
		Code:   BAR,
		Symbol: "Bar",
	},
	PASCAL: {
		Code:   PASCAL,
		Symbol: "Pa",
	},
	HECTOPASCAL: {
		Code:   HECTOPASCAL,
		Symbol: "hPa",
	},
	MILLIBAR: {
		Code:   MILLIBAR,
		Symbol: "mBar",
	},
	ATM: {
		Code:   ATM,
		Symbol: "atm",
	},
	PSI: {
		Code:   PSI,
		Symbol: "PSI",
	},
	TORR: {
		Code:   TORR,
		Symbol: "Torr",
	},
	MMHG: {
		Code:   MMHG,
		Symbol: "mmHg",
	},

	// Energy
	JOULE: {
		Code:   JOULE,
		Symbol: "J",
	},
	CALORIE: {
		Code:   CALORIE,
		Symbol: "cal",
	},
	KILOCALORIE: {
		Code:   KILOCALORIE,
		Symbol: "kcal",
	},
	KILOJOULE: {
		Code:   KILOJOULE,
		Symbol: "kJ",
	},

	// Piece / CaseBox Units
	PCS: {
		Code:   PCS,
		Symbol: "pcs",
	},
	PIECE: {
		Code:   PIECE,
		Symbol: "Pc",
	},
	DOZEN: {
		Code:   DOZEN,
		Symbol: "Dz",
	},
}
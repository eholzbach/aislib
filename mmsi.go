// Copyright (c) 2015, Marios Andreopoulos.
//
// This file is part of aislib.
//
//  Aislib is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
//  Aislib is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
// along with aislib.  If not, see <http://www.gnu.org/licenses/>.

package aislib

// Contains MMSI owners' descriptions. Currently not used anywhere
var MmsiCodes = [...]string{
	"Ship", "Coastal Station", "Group of ships", "SAR —Search and Rescue Aircraft",
	"Diver's radio", "Aids to navigation", "Auxiliary craft associated with parent ship",
	"AIS SART —Search and Rescue Transmitter", "MOB —Man Overboard Device",
	"EPIRB —Emergency Position Indicating Radio Beacon", "Invalid MMSI",
}

// DecodeMMSI returns a string with the type of the owner of the MMSI and its country
// Some MMSIs aren't valid. There is some more information in some MMSIs (the satellite
// equipment of the ship). We may add them in the future.
// Have a look at http://en.wikipedia.org/wiki/Maritime_Mobile_Service_Identity
func DecodeMMSI(m uint32) (string, string) {
	owner := ""
	country := ""
	mid := uint32(1000)

	// Current intervals:
	// [0 00999999][010000000 099999999][100000000 199999999][200000000 799999999]
	// [800000000 899999999]...[970000000 970999999]...[972000000 972999999]...
	// [974000000 974999999]...[98000000 98999999][99000000 999999999]
	switch {
	case m >= 200000000 && m < 800000000:
		mid = m / 1000000
		owner = "Ship"
	case m <= 9999999:
		mid = m / 10000
		owner = "Coastal Station"
	case m <= 99999999:
		mid = m / 100000
		owner = "Group of ships"
	case m <= 199999999:
		mid = m/1000 - 111000
		owner = "SAR —Search and Rescue Aircraft"
	case m < 900000000:
		mid = m/100000 - 8000
		owner = "Diver's radio"
	case m >= 990000000 && m < 1000000000:
		mid = m/10000 - 99000
		owner = "Aids to navigation"
	case m >= 980000000 && m < 990000000:
		mid = m/10000 - 98000
		owner = "Auxiliary craft associated with parent ship"
	case m >= 970000000 && m < 970999999:
		mid = m/1000 - 970000
		owner = "AIS SART —Search and Rescue Transmitter"
	case m >= 972000000 && m < 972999999:
		owner = "MOB —Man Overboard Device"
	case m >= 974000000 && m < 974999999:
		owner = "EPIRB —Emergency Position Indicating Radio Beacon"
	default:
		owner = "Invalid MMSI"
	}

	if mid < 1000 {
		country = Mid[int(mid)]
		if country == "" {
			country = "Unknown Country ID"
		}
		return owner, country
	}
	return owner, country
}

// Maritime Identification Digits, have a look at http://www.itu.int/online/mms/glad/cga_mids.sh?lang=en
var Mid = map[int]string{
	201: "Albania",
	202: "Andorra",
	203: "Austria",
	204: "Azores",
	205: "Belgium",
	206: "Belarus",
	207: "Bulgaria",
	208: "Vatican",
	209: "Cyprus",
	210: "Cyprus",
	211: "Germany",
	212: "Cyprus",
	213: "Georgia",
	214: "Moldova",
	215: "Malta",
	216: "Armenia",
	218: "Germany",
	219: "Denmark",
	220: "Denmark",
	224: "Spain",
	225: "Spain",
	226: "France",
	227: "France",
	228: "France",
	229: "Malta",
	230: "Finland",
	231: "Faroe Islands",
	232: "United Kingdom",
	233: "United Kingdom",
	234: "United Kingdom",
	235: "United Kingdom",
	236: "Gibraltar",
	237: "Greece",
	238: "Croatia",
	239: "Greece",
	240: "Greece",
	241: "Greece",
	242: "Morocco",
	243: "Hungary",
	244: "Netherlands",
	245: "Netherlands",
	246: "Netherlands",
	247: "Italy",
	248: "Malta",
	249: "Malta",
	250: "Ireland",
	251: "Iceland",
	252: "Liechtenstein",
	253: "Luxembourg",
	254: "Monaco",
	255: "Madeira",
	256: "Malta",
	257: "Norway",
	258: "Norway",
	259: "Norway",
	261: "Poland",
	262: "Montenegro",
	263: "Portugal",
	264: "Romania",
	265: "Sweden",
	266: "Sweden",
	267: "Slovak Republic",
	268: "San Marino",
	269: "Switzerland",
	270: "Czech Republic",
	271: "Turkey",
	272: "Ukraine",
	273: "Russian Federation",
	274: "The Former Yugoslav Republic of Macedonia",
	275: "Latvia",
	276: "Estonia",
	277: "Lithuania",
	278: "Slovenia",
	279: "Serbia",
	301: "Anguilla",
	303: "Alaska",
	304: "Antigua and Barbuda",
	305: "Antigua and Barbuda",
	306: "Curacao Sint Maarten",
	307: "Aruba Netherlands",
	308: "Bahamas",
	309: "Bahamas",
	310: "Bermuda",
	311: "Bahamas",
	312: "Belize",
	314: "Barbados",
	316: "Canada",
	319: "Cayman Islands",
	321: "Costa Rica",
	323: "Cuba",
	325: "Dominica",
	327: "Dominican Republic",
	329: "Guadeloupe",
	330: "Grenada",
	331: "Greenland",
	332: "Guatemala",
	334: "Honduras",
	336: "Haiti",
	338: "United States of America",
	339: "Jamaica",
	341: "Saint Kitts and Nevis",
	343: "Saint Lucia",
	345: "Mexico",
	347: "Martiniquee",
	348: "Montserrat",
	350: "Nicaragua",
	351: "Panama",
	352: "Panama",
	353: "Panama",
	354: "Panama",
	355: " - ",
	356: " - ",
	357: " - ",
	358: "Puerto Rico",
	359: "El Salvador",
	361: "Saint Pierre and Miquelon",
	362: "Trinidad and Tobago",
	364: "Turks and Caicos Islands",
	366: "United States of America",
	367: "United States of America",
	368: "United States of America",
	369: "United States of America",
	370: "Panama",
	371: "Panama",
	372: "Panama",
	373: "Panama",
	375: "Saint Vincent and the Grenadines",
	376: "Saint Vincent and the Grenadines",
	377: "Saint Vincent and the Grenadines",
	378: "British Virgin Islands",
	379: "United States Virgin Islands",
	401: "Afghanistan",
	403: "Saudi Arabia",
	405: "Bangladesh",
	408: "Bahrain",
	410: "Bhutan",
	412: "China",
	413: "China",
	414: "China",
	416: "Taiwan)",
	417: "Sri Lanka",
	419: "India",
	422: "Iran",
	423: "Azerbaijan",
	425: "Iraq",
	428: "Israel",
	431: "Japan",
	432: "Japan",
	434: "Turkmenistan",
	436: "Kazakhstan",
	437: "Uzbekistan",
	438: "Jordan",
	440: "Korea",
	441: "Korea",
	443: "State of Palestine",
	445: "Democratic Peoples Republic of Korea",
	447: "Kuwait",
	450: "Lebanon",
	451: "Kyrgyz Republic",
	453: "Macao",
	455: "Maldives",
	457: "Mongolia",
	459: "Nepal",
	461: "Oman",
	463: "Pakistan",
	466: "Qatar",
	468: "Syrian Arab Republic",
	470: "United Arab Emirates",
	472: "Tajikistan",
	473: "Yemen",
	475: "Yemen",
	477: "Hong Kong",
	478: "Bosnia and Herzegovina",
	501: "Adelie Land",
	503: "Australia",
	506: "Myanmar",
	508: "Brunei Darussalam",
	510: "Micronesia",
	511: "Palau",
	512: "New Zealand",
	514: "Cambodia",
	515: "Cambodia",
	516: "Christmas Island",
	518: "Cook Islands",
	520: "Fiji",
	523: "Cocos Islands",
	525: "Indonesia",
	529: "Kiribati",
	531: "Lao Peoples Democratic Republic",
	533: "Malaysia",
	536: "Northern Mariana Islands",
	538: "Marshall Islands",
	540: "New Caledonia",
	542: "Niue",
	544: "Nauru",
	546: "French Polynesia",
	548: "Philippines",
	553: "Papua New Guinea",
	555: "Pitcairn Island",
	557: "Solomon Islands",
	559: "American Samoa",
	561: "Samoa",
	563: "Singapore",
	564: "Singapore",
	565: "Singapore",
	566: "Singapore",
	567: "Thailand",
	570: "Tonga",
	572: "Tuvalu",
	574: "Viet Nam",
	576: "Vanuatu",
	577: "Vanuatu",
	578: "Wallis and Futuna Islands",
	601: "South Africa",
	603: "Angola",
	605: "Algeria",
	607: "Saint Paul and Amsterdam Islands",
	608: "Ascension Island",
	609: "Burundi",
	610: "Benin",
	611: "Botswana",
	612: "Central African Republic",
	613: "Cameroon",
	615: "Congo",
	616: "Comoros",
	617: "Cabo Verde",
	618: "Crozet Archipelago",
	619: "Cote dIvoire",
	620: "Comoros",
	621: "Djibouti",
	622: "Egypt",
	624: "Ethiopia",
	625: "Eritrea",
	626: "Gabonese Republic",
	627: "Ghana",
	629: "Gambia",
	630: "Guinea-Bissau",
	631: "Equatorial Guinea",
	632: "Guinea",
	633: "Burkina Faso",
	634: "Kenya",
	635: "Kerguelen Islands",
	636: "Liberia",
	637: "Liberia",
	638: "South Sudan",
	642: "Libya",
	644: "Lesotho",
	645: "Mauritius",
	647: "Madagascar",
	649: "Mali",
	650: "Mozambique",
	654: "Mauritania",
	655: "Malawi",
	656: "Niger",
	657: "Nigeria",
	659: "Namibia",
	660: "Reunion",
	661: "Rwanda",
	662: "Sudan",
	663: "Senegal",
	664: "Seychelles",
	665: "Saint Helena",
	666: "Somalia",
	667: "Sierra Leone",
	668: "Sao Tome and Principe",
	669: "Swaziland",
	670: "Chad",
	671: "Togolese Republic",
	672: "Tunisia",
	674: "Tanzania",
	675: "Uganda",
	676: "Democratic Republic of the Congo",
	677: "Tanzania",
	678: "Zambia",
	679: "Zimbabwe",
	701: "Argentine Republic",
	710: "Brazil",
	720: "Bolivia",
	725: "Chile",
	730: "Colombia",
	735: "Ecuador",
	740: "Falkland Islands",
	745: "Guiana",
	750: "Guyana",
	755: "Paraguay",
	760: "Peru",
	765: "Suriname",
	770: "Uruguay",
}

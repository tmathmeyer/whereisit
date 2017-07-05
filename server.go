package main

import (
	"encoding/json"
	"html/template"
	"log"
    "os"
	"net/http"
	"database/sql"
	"strings"
	"math/rand"
	"time"
	_ "github.com/lib/pq"
)

type Map struct {
	Map  string
}

type Guess struct {
	Lat float64
	Lon float64
}

const serverPort string = ":8081"
var db *sql.DB


var countries = [...]string{
	"Afghanistan",
	"Albania",
	"Algeria",
	"American_Samoa",
	"Andorra",
	"Angola",
	"Anguilla",
	"Antarctica",
	"Antigua_and_Barbuda",
	"Argentina",
	"Armenia",
	"Aruba",
	"Australia",
	"Austria",
	"Azerbaijan",
	"Bahamas",
	"Bahrain",
	"Bangladesh",
	"Barbados",
	"Belarus",
	"Belgium",
	"Belize",
	"Benin",
	"Bermuda",
	"Bhutan",
	"Bolivia",
	"Bosnia_and_Herzegovina",
	"Botswana",
	"Bouvet_Island",
	"Brazil",
	"British_Antarctic_Territory",
	"British_Indian_Ocean_Territory",
	"British_Virgin_Islands",
	"Brunei",
	"Bulgaria",
	"Burkina_Faso",
	"Burundi",
	"Cambodia",
	"Cameroon",
	"Canada",
	"Canton_and_Enderbury_Islands",
	"Cape_Verde",
	"Cayman_Islands",
	"Central_African_Republic",
	"Chad",
	"Chile",
	"China",
	"Christmas_Island",
	"Cocos_[Keeling]_Islands",
	"Colombia",
	"Comoros",
	"Congo_-_Brazzaville",
	"Congo_-_Kinshasa",
	"Cook_Islands",
	"Costa_Rica",
	"Croatia",
	"Cuba",
	"Cyprus",
	"Czech_Republic",
	"Côte_d’Ivoire",
	"Denmark",
	"Djibouti",
	"Dominica",
	"Dominican_Republic",
	"Dronning_Maud_Land",
	"East_Germany",
	"Ecuador",
	"Egypt",
	"El_Salvador",
	"Equatorial_Guinea",
	"Eritrea",
	"Estonia",
	"Ethiopia",
	"Falkland_Islands",
	"Faroe_Islands",
	"Fiji",
	"Finland",
	"France",
	"French_Guiana",
	"French_Polynesia",
	"French_Southern_Territories",
	"French_Southern_and_Antarctic_Territories",
	"Gabon",
	"Gambia",
	"Georgia",
	"Germany",
	"Ghana",
	"Gibraltar",
	"Greece",
	"Greenland",
	"Grenada",
	"Guadeloupe",
	"Guam",
	"Guatemala",
	"Guernsey",
	"Guinea",
	"Guinea-Bissau",
	"Guyana",
	"Haiti",
	"Heard_Island_and_McDonald_Islands",
	"Honduras",
	"Hong_Kong_SAR_China",
	"Hungary",
	"Iceland",
	"India",
	"Indonesia",
	"Iran",
	"Iraq",
	"Ireland",
	"Isle_of_Man",
	"Israel",
	"Italy",
	"Jamaica",
	"Japan",
	"Jersey",
	"Johnston_Island",
	"Jordan",
	"Kazakhstan",
	"Kenya",
	"Kiribati",
	"Kuwait",
	"Kyrgyzstan",
	"Laos",
	"Latvia",
	"Lebanon",
	"Lesotho",
	"Liberia",
	"Libya",
	"Liechtenstein",
	"Lithuania",
	"Luxembourg",
	"Macau_SAR_China",
	"Macedonia",
	"Madagascar",
	"Malawi",
	"Malaysia",
	"Maldives",
	"Mali",
	"Malta",
	"Marshall_Islands",
	"Martinique",
	"Mauritania",
	"Mauritius",
	"Mayotte",
	"Metropolitan_France",
	"Mexico",
	"Micronesia",
	"Midway_Islands",
	"Moldova",
	"Monaco",
	"Mongolia",
	"Montenegro",
	"Montserrat",
	"Morocco",
	"Mozambique",
	"Myanmar_[Burma]",
	"Namibia",
	"Nauru",
	"Nepal",
	"Netherlands",
	"Netherlands_Antilles",
	"Neutral_Zone",
	"New_Caledonia",
	"New_Zealand",
	"Nicaragua",
	"Niger",
	"Nigeria",
	"Niue",
	"Norfolk_Island",
	"North_Korea",
	"North_Vietnam",
	"Northern_Mariana_Islands",
	"Norway",
	"Oman",
	"Pacific_Islands_Trust_Territory",
	"Pakistan",
	"Palau",
	"Palestinian_Territories",
	"Panama",
	"Panama_Canal_Zone",
	"Papua_New_Guinea",
	"Paraguay",
	"People's_Democratic_Republic_of_Yemen",
	"Peru",
	"Philippines",
	"Pitcairn_Islands",
	"Poland",
	"Portugal",
	"Puerto_Rico",
	"Qatar",
	"Romania",
	"Russia",
	"Rwanda",
	"Réunion",
	"Saint_Barthélemy",
	"Saint_Helena",
	"Saint_Kitts_and_Nevis",
	"Saint_Lucia",
	"Saint_Martin",
	"Saint_Pierre_and_Miquelon",
	"Saint_Vincent_and_the_Grenadines",
	"Samoa",
	"San_Marino",
	"Saudi_Arabia",
	"Senegal",
	"Serbia",
	"Serbia",
	"Montenegro",
	"Seychelles",
	"Sierra_Leone",
	"Singapore",
	"Slovakia",
	"Slovenia",
	"Solomon_Islands",
	"Somalia",
	"South_Africa",
	"South_Georgia_and_the_South_Sandwich_Islands",
	"South_Korea",
	"Spain",
	"Sri_Lanka",
	"Sudan",
	"Suriname",
	"Svalbard_and_Jan_Mayen",
	"Swaziland",
	"Sweden",
	"Switzerland",
	"Syria",
	"São_Tomé_and_Príncipe",
	"Taiwan",
	"Tajikistan",
	"Tanzania",
	"Thailand",
	"Timor-Leste",
	"Togo",
	"Tokelau",
	"Tonga",
	"Trinidad_and_Tobago",
	"Tunisia",
	"Turkey",
	"Turkmenistan",
	"Turks_and_Caicos_Islands",
	"Tuvalu",
	"U.S._Minor_Outlying_Islands",
	"U.S._Miscellaneous_Pacific_Islands",
	"U.S._Virgin_Islands",
	"Uganda",
	"Ukraine",
	"Union_of_Soviet_Socialist_Republics",
	"United_Arab_Emirates",
	"United_Kingdom",
	"United_States",
	"Unknown_or_Invalid_Region",
	"Uruguay",
	"Uzbekistan",
	"Vanuatu",
	"Vatican_City",
	"Venezuela",
	"Vietnam",
	"Wake_Island",
	"Wallis_and_Futuna",
	"Western_Sahara",
	"Yemen",
	"Zambia",
	"Zimbabwe",
	"Aland_Islands",
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/guessmap/"+countries[rand.Intn(len(countries))], 307)
}

func getPathParts(r *http.Request, str string) []string {
	restPath := r.URL.Path[len(str):]
	return strings.Split(restPath, "/")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, countries)
}

func vizHandler(w http.ResponseWriter, r *http.Request) {
	var Map = Map {
		Map: getPathParts(r, "/viz/")[0],
	}
	t, _ := template.ParseFiles("static/viz.html")
	t.Execute(w, Map)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	var Map = Map {
		Map: getPathParts(r, "/guessmap/")[0],
	}
	t, _ := template.ParseFiles("static/makeguess.html")
	t.Execute(w, Map)
}

func guessRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	loc := getPathParts(r, "/api/guess/")[0]
    var t Guess
    for key, _ := range r.Form {
        err := json.Unmarshal([]byte(key), &t)
        if err != nil {
        	log.Println("--guessRequest-- Unmarshall Error")
            log.Println(err.Error())
        } else {
        	_, err = db.Exec("INSERT INTO guesses VALUES ($1,$2,$3)", loc, t.Lat, t.Lon)
        	if err != nil {
        		log.Println("--guessRequest-- Exec Error")
            	log.Println(err.Error())
        	}
        	return
        }
    }
}

func resultServer(w http.ResponseWriter, r *http.Request) {
	var guesses []Guess
	loc := getPathParts(r, "/api/results/")[0]
	rows, err := db.Query("SELECT lat,lon FROM guesses WHERE map = $1", loc)
	if err != nil {
		log.Println("--resultServer-- Query Error")
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	} else {
		defer rows.Close()
		for rows.Next() {
			g := new(Guess)
			if err := rows.Scan(&g.Lat, &g.Lon); err != nil {
				log.Println("--resultServer-- Scan Error")
				log.Println(err)
				http.Error(w, err.Error(), 500)
			}
			guesses = append(guesses, *g)
		}
	}
	js, err := json.Marshal(guesses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	data, err := sql.Open("postgres", "user=whereisit dbname=whereIsItGuesses password="+os.Args[1]+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
		return
	}

	db = data
	rand.Seed(time.Now().Unix())

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/guess/", redirect)
	http.HandleFunc("/api/guess/", guessRequest)
	http.HandleFunc("/api/results/", resultServer)
	http.HandleFunc("/viz/", vizHandler)
	http.HandleFunc("/guessmap/", guessHandler)
	http.ListenAndServe(serverPort, nil)
}

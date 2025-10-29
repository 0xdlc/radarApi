package flightRadar

import (
	"encoding/json"
	"fmt"
	"os"
	//"net/http"
	"io/ioutil"
	"net/url"
	"sync"
	"strconv"
	//"crypto/tls"
	//"reflect"
	//"time"
	//"golang.org/x/net/http2"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"path"
	// "path/filepath"
	// "go.mongodb.org/mongo-driver/bson"
)

type Bound struct {
	TLX float64 `json:"tl_x"`
	TLY float64 `json:"tl_y"`
	BRX float64 `json:"br_x"`
	BRY float64 `json:"br_y"`
}
type FlightBounds struct {
	Bounds []Bound `json:"bounds"`
}

type FlightDetails struct {
	Identification struct {
		ID     string `json:"id"`
		Row    int64  `json:"row"`
		Number struct {
			Default     interface{} `json:"default"`
			Alternative interface{} `json:"alternative"`
		} `json:"number"`
		Callsign string `json:"callsign"`
	} `json:"identification"`
	Status struct {
		Live      bool        `json:"live"`
		Text      string      `json:"text"`
		Icon      interface{} `json:"icon"`
		Estimated interface{} `json:"estimated"`
		Ambiguous bool        `json:"ambiguous"`
		Generic   struct {
			Status struct {
				Text  string `json:"text"`
				Color string `json:"color"`
				Type  string `json:"type"`
			} `json:"status"`
		} `json:"generic"`
	} `json:"status"`
	Level    string `json:"level"`
	Promote  bool   `json:"promote"`
	Aircraft struct {
		Model struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"model"`
		CountryID    int         `json:"countryId"`
		Registration string      `json:"registration"`
		Age          interface{} `json:"age"`
		Msn          interface{} `json:"msn"`
		Images       struct {
			Thumbnails []struct {
				Src       string `json:"src"`
				Link      string `json:"link"`
				Copyright string `json:"copyright"`
				Source    string `json:"source"`
			} `json:"thumbnails"`
			Medium []struct {
				Src       string `json:"src"`
				Link      string `json:"link"`
				Copyright string `json:"copyright"`
				Source    string `json:"source"`
			} `json:"medium"`
			Large []struct {
				Src       string `json:"src"`
				Link      string `json:"link"`
				Copyright string `json:"copyright"`
				Source    string `json:"source"`
			} `json:"large"`
		} `json:"images"`
		Hex string `json:"hex"`
	} `json:"aircraft"`
	Airline struct {
		Name  string `json:"name"`
		Short string `json:"short"`
		Code  struct {
			Iata interface{} `json:"iata"`
			Icao string      `json:"icao"`
		} `json:"code"`
		URL string `json:"url"`
	} `json:"airline"`
	Owner    interface{} `json:"owner"`
	Airspace interface{} `json:"airspace"`
	Airport  struct {
		Origin struct {
			Name string `json:"name"`
			Code struct {
				Iata string `json:"iata"`
				Icao string `json:"icao"`
			} `json:"code"`
			Position struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Altitude  int     `json:"altitude"`
				Country   struct {
					ID   interface{} `json:"id"`
					Name string      `json:"name"`
					Code string      `json:"code"`
				} `json:"country"`
				Region struct {
					City string `json:"city"`
				} `json:"region"`
			} `json:"position"`
			Timezone struct {
				Name        string `json:"name"`
				Offset      int    `json:"offset"`
				OffsetHours string `json:"offsetHours"`
				Abbr        string `json:"abbr"`
				AbbrName    string `json:"abbrName"`
				IsDst       bool   `json:"isDst"`
			} `json:"timezone"`
			Visible bool        `json:"visible"`
			Website interface{} `json:"website"`
			Info    struct {
				Terminal interface{} `json:"terminal"`
				Baggage  interface{} `json:"baggage"`
				Gate     interface{} `json:"gate"`
			} `json:"info"`
		} `json:"origin"`
		Destination interface{} `json:"destination"`
		Real        interface{} `json:"real"`
	} `json:"airport"`
	FlightHistory struct {
		Aircraft []struct {
			Identification struct {
				ID     string `json:"id"`
				Number struct {
					Default interface{} `json:"default"`
				} `json:"number"`
			} `json:"identification"`
			Airport struct {
				Origin struct {
					Name string `json:"name"`
					Code struct {
						Iata string `json:"iata"`
						Icao string `json:"icao"`
					} `json:"code"`
					Position struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Altitude  int     `json:"altitude"`
						Country   struct {
							ID   interface{} `json:"id"`
							Name string      `json:"name"`
							Code string      `json:"code"`
						} `json:"country"`
						Region struct {
							City string `json:"city"`
						} `json:"region"`
					} `json:"position"`
					Timezone struct {
						Name        string `json:"name"`
						Offset      int    `json:"offset"`
						OffsetHours string `json:"offsetHours"`
						Abbr        string `json:"abbr"`
						AbbrName    string `json:"abbrName"`
						IsDst       bool   `json:"isDst"`
					} `json:"timezone"`
					Visible bool        `json:"visible"`
					Website interface{} `json:"website"`
				} `json:"origin"`
				Destination struct {
					Name string `json:"name"`
					Code struct {
						Iata string `json:"iata"`
						Icao string `json:"icao"`
					} `json:"code"`
					Position struct {
						Latitude  float64 `json:"latitude"`
						Longitude float64 `json:"longitude"`
						Altitude  int     `json:"altitude"`
						Country   struct {
							ID   interface{} `json:"id"`
							Name string      `json:"name"`
							Code string      `json:"code"`
						} `json:"country"`
						Region struct {
							City string `json:"city"`
						} `json:"region"`
					} `json:"position"`
					Timezone struct {
						Name        string `json:"name"`
						Offset      int    `json:"offset"`
						OffsetHours string `json:"offsetHours"`
						Abbr        string `json:"abbr"`
						AbbrName    string `json:"abbrName"`
						IsDst       bool   `json:"isDst"`
					} `json:"timezone"`
					Visible bool        `json:"visible"`
					Website interface{} `json:"website"`
				} `json:"destination"`
			} `json:"airport"`
			Time struct {
				Real struct {
					Departure int `json:"departure"`
				} `json:"real"`
			} `json:"time"`
		} `json:"aircraft"`
	} `json:"flightHistory"`
	Ems          interface{} `json:"ems"`
	Availability []string    `json:"availability"`
	Time         struct {
		Scheduled struct {
			Departure int `json:"departure"`
			Arrival   int `json:"arrival"`
		} `json:"scheduled"`
		Real struct {
			Departure int         `json:"departure"`
			Arrival   interface{} `json:"arrival"`
		} `json:"real"`
		Estimated struct {
			Departure interface{} `json:"departure"`
			Arrival   interface{} `json:"arrival"`
		} `json:"estimated"`
		Other struct {
			Eta     int `json:"eta"`
			Updated int `json:"updated"`
		} `json:"other"`
		Historical interface{} `json:"historical"`
	} `json:"time"`
	Trail []struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
		Alt int     `json:"alt"`
		Spd int     `json:"spd"`
		Ts  int     `json:"ts"`
		Hd  int     `json:"hd"`
	} `json:"trail"`
	FirstTimestamp int    `json:"firstTimestamp"`
	S              string `json:"s"`
}

type ApiStruct struct {
	FullCount int `json:"full_count"`
	Version int `json:"version"`
	Data map[string]interface{} `json:"-"`
}

var (
	flightIDs []string
)
func Start() {

	currentDir, err := os.Getwd()
	if _, e := os.Stat(path.Join(currentDir, "Data")); os.IsNotExist(e) {
		os.Mkdir(path.Join(currentDir, "Data"), 0777)
	}

	file, err := os.Open("flightRadar/flightBounds.json")
	if err != nil {
		fmt.Println("There was an error reading the fligh bounds file (flightBounds.json)")
		return
	}
	defer file.Close()

	var flightBounds FlightBounds
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&flightBounds)
	if err != nil {
		fmt.Println("There was an error")
		return
	}
	
	jar := tls_client.NewCookieJar()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(profiles.Chrome_120),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar), // create cookieJar instance and pass it as argument
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
			fmt.Println(err)
			return
		}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 4)
	for _, bound := range flightBounds.Bounds {
		wg.Add(1)
		go getFlights(bound, client, &wg, sem)
	}
	wg.Wait()

	// client := new(http.Client(tr))
	//disable certificate check
	// b := fmt.Sprintf("%s,%s,%s,%s", flightBounds.Bounds[0].TLX, flightBounds[0].TLY, flightBounds[0].BRX, flightBounds[0].BRY)
	// params := url.Values{}
	// params.Add("bounds",b)
}


func getFlights(bound Bound, client tls_client.HttpClient, wg *sync.WaitGroup, sem chan struct{}){
	defer wg.Done()
	sem <- struct{}{} // Acquire a token
	defer func() { <-sem }() // Release the token when done
	//Http client:
	//note: using tls-client package to imporsonate TLS fingerprinting to bypass cloudflare's restrictions
	
	realTimeFlightDataURl := "https://data-cloud.flightradar24.com/zones/fcgi/feed.js"
	//uri := fmt.Sprintf("%.2f,%.2f,%.2f,%.2f", bound.TLY, bound.BRY, bound.TLX, bound.BRX)
	uri := url.QueryEscape(fmt.Sprintf("%.2f,%.2f,%.2f,%.2f", bound.TLY, bound.BRY, bound.TLX, bound.BRX))
	//uri := url.QueryEscape("74.0,70.0,20.0,28.0")
	reqURL := fmt.Sprintf("%s?bounds=%s", realTimeFlightDataURl, uri)
	req, err := http.NewRequest("GET", reqURL, nil)
	//TODO: change this to a setter so you set the headers only once
	req.Header.Set("accept-encoding", "gzip, br")
	req.Header.Set("accept-language", "pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("origin", "https://www.flightradar24.com")
	req.Header.Set("referer", "https://www.flightradar24.com/")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	res, err := client.Do(req)
	
	if err!=nil {
		fmt.Println("There was an error sending your request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	fmt.Println("Error reading response body:", err)
	return
	}
	
	var JsonResponse ApiStruct
	if err = json.Unmarshal(body, &JsonResponse); err != nil{
		fmt.Println("THere was an error encoding the json", err)
		return
	}

	JsonResponse.Data = make(map[string]interface{})
	if err := json.Unmarshal(body, &JsonResponse.Data); err != nil {
		fmt.Println("Error:", err)
		return
	}
	// var FlightData Flights
	outerloop:
		for key, _ := range JsonResponse.Data {
			if key == "version" || key == "full_count"{
				continue
			}
			for _, f := range flightIDs{
				if key == f{
					fmt.Println("found same id", key)
					continue outerloop
				}
			}
			getFlightDetail(key, client)
			flightIDs = append(flightIDs, key)
				// if assertedInfo, ok := info.([]interface{}); ok {
				// 	FlightData.ID = key
				// 	//FlightData.ICAO = assertedInfo[0].(string)
				// 	//FlightData.Heading = assertedInfo[3].(string)
				// 	FlightData.Altitude = assertedInfo[4].(float64)
				// 	FlightData.GroundSpeed = assertedInfo[5].(float64)
				// 	//FlightData.Squawk = assertedInfo[6].(string)
				// 	FlightData.AircraftCode = assertedInfo[8].(string)
				// 	FlightData.Registration = assertedInfo[9].(string)
				// 	FlightData.OriginAirportIATA = assertedInfo[11].(string)
				// 	FlightData.DestinationAirportIATA = assertedInfo[12].(string)
				// 	FlightData.FlightNumber = assertedInfo[13].(string)
				// 	FlightData.Onground = assertedInfo[14].(float64)
				// 	FlightData.VerticalSpeed = assertedInfo[15].(float64)
				// 	FlightData.Callsign = assertedInfo[16].(string)
				// 	FlightData.AirlineICAO = assertedInfo[18].(string)
				// }
				//fmt.Println(FlightData)
		}
	fmt.Println(flightIDs)
		
		
		 		//checking if the folder with directory name of the flightnumber(key) exist or not
		 		// currentDir, _ := os.Getwd()
		 		// if assertedInfo, ok := info.([]interface{}); ok {
		 		// 	if assertedTimestamp, ok := assertedInfo[0].(string); ok {
		 		// 		FilePath := path.Join(currentDir, assertedTimestamp)
					
		 		// //info[8] is the aircraftID which is uniq for every aircraft
		 		// 		if _ ,err := os.Stat(FilePath); os.IsNotExist(err){
		 		// 			os.Mkdir(FilePath, 0777)
		 		// 		}

		// 				bsondata, err := bson.Marshal(JsonResponse)
		// 				if err!=nil{
		// 					fmt.Println("[Error] There was an error Writing data to File", err)
		// 					return
		// 				}
		// 				err = os.WriteFile(filepath.Join(FilePath,info[10]),bsondata,0777)
		// 				if err!=nil{
		// 					fmt.Println("[Error] There was an error Writing data to File", err)
		// 					return
		// 				}
		// 				fmt.Println(len(JsonResponse))
		// 			}
		// 		}
		// 	}
		// }
		//time.Sleep(1 * time.Second)
	}




func getFlightDetail(flightNumber string, client tls_client.HttpClient){
	url := fmt.Sprintf("https://data-live.flightradar24.com/clickhandler/?flight=%s", flightNumber)

	req, err := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err!=nil {
	fmt.Println("There was an error sending your request", err)
	return
	}
	defer res.Body.Close()


	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	fmt.Println("Error reading response body:", err)
	return
	}

	var JsonResponse FlightDetails
	if err = json.Unmarshal(body, &JsonResponse); err != nil{
		fmt.Println("THere was an error encoding the json", err)
		return
	}

	currentDir, err := os.Getwd()
	flightRegNum := JsonResponse.Aircraft.Registration
	flightDir := path.Join(currentDir, "Data", flightRegNum)
	if _, e := os.Stat(flightDir); os.IsNotExist(e) {
		os.Mkdir(flightDir, 0777)
	}
	
	output, err := json.Marshal(JsonResponse)
	file, err := os.Create(path.Join(flightDir, strconv.Itoa(JsonResponse.FlightHistory.Aircraft[0].Time.Real.Departure)+".json"))
	defer file.Close()

	_, err = file.Write(output)
	if err!=nil{fmt.Println("[Err] There was an Error Writing Into file")}
	return

}


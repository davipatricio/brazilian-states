package brazilianstates

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type StateInfo struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Capital string `json:"capital"`
	Region  string `json:"region"`
	Cities  int    `json:"cities"`
}

// Returns a slice of all states
func GetAll() ([]StateInfo, error) {
	// Create a variable to Unmarshal the JSON data into
	var states []StateInfo

	// Open the JSON file
	jsonFile, err := ioutil.ReadFile("./data/states.json")
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into the variable
	err = json.Unmarshal(jsonFile, &states)
	if err != nil {
		return nil, err
	}

	return states, nil
}

// Returns a object of a state by its code
// If empty, returns an error
//  state, err := brazilianstates.Get("AC")
//  if err != nil {
//   panic(err)
//  }
//  fmt.Println(state)
//  fmt.Println(state.Code)
func Get(uf string) (StateInfo, error) {
	if uf == "" {
		return StateInfo{}, errors.New("uf cannot be empty")
	}

	allStates, err := GetAll()
	if err != nil {
		return StateInfo{}, err
	}

	for _, state := range allStates {
		if state.Code == uf {
			return state, nil
		}
	}

	return StateInfo{}, errors.New("uf not found")
}

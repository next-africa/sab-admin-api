package graphql

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sab.com/domain/university"
	"strconv"
	"strings"
)

type UniversityNode struct {
	Id         string                 `json:"id"`
	Properties *university.University `json:"properties"`
}

func newUniversityNodeFromUniversity(aUniversity *university.University, countryCode string) UniversityNode {
	id := fmt.Sprintf("%s:%s:%s", "University", countryCode, aUniversity.Id)
	id = base64.StdEncoding.EncodeToString([]byte(id))

	return UniversityNode{id, aUniversity}
}

func mapUniversitiesToUniversityNodes(universities []university.University, countryCode string) []UniversityNode {
	universitiesMap := make([]UniversityNode, len(universities))
	for i, v := range universities {
		universitiesMap[i] = newUniversityNodeFromUniversity(&v, countryCode)
	}
	return universitiesMap
}

func getUniversityByGlobalId(encodedGlobalId string, universityService *university.UniversityService) (UniversityNode, error) {
	if decoded, err := base64.StdEncoding.DecodeString(encodedGlobalId); err != nil {
		return UniversityNode{}, err
	} else {
		idParts := strings.Split(string(decoded), ":")

		if len(idParts) != 3 {
			return UniversityNode{}, errors.New("Invalid global university Id, the country relay Id should be of the form University:{countryCode}:{universityId}")
		}

		countryCode := idParts[1]
		universityId := idParts[2]

		return getUniversityNodeByCountryCodeAndUniversityId(universityId, countryCode, universityService)
	}
}

func getUniversityNodeByCountryCodeAndUniversityId(universityIdString string, countryCode string, universityService *university.UniversityService) (UniversityNode, error) {
	universityId, err := strconv.ParseInt(universityIdString, 10, 64)
	if err != nil {
		return UniversityNode{}, errors.New("Invalid university Id, university Id should be an Integer")
	}

	theUniversity, err := universityService.GetUniversityByIdAndCountryCode(universityId, countryCode)

	if err != nil {
		return UniversityNode{}, err
	}

	return newUniversityNodeFromUniversity(&theUniversity, countryCode), nil
}

func getAllUniversities(countryCode string, universityService *university.UniversityService) ([]UniversityNode, error) {
	if universities, err := universityService.GetAllUniversitiesForCountryCode(countryCode); err != nil {
		return []UniversityNode{}, nil
	} else {
		return mapUniversitiesToUniversityNodes(universities, countryCode), nil
	}
}

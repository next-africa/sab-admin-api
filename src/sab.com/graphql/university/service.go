package graphql

import (
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

type UniversityGraphService struct {
	universityService *university.UniversityService
}

func NewUniversityGraphqlService(universityService *university.UniversityService) UniversityGraphService {
	return UniversityGraphService{universityService: universityService}
}

func ComputeUniversityGlobalId(countryCode string, universityId int64) (globalId string) {
	globalId = fmt.Sprintf("%s:%v", countryCode, universityId)
	return
}

func (graphqlService *UniversityGraphService) NewUniversityNodeFromUniversity(aUniversity *university.University, countryCode string) UniversityNode {
	return UniversityNode{ComputeUniversityGlobalId(countryCode, aUniversity.Id), aUniversity}
}

func (graphqlService *UniversityGraphService) mapUniversitiesToUniversityNodes(universities []university.University, countryCode string) []UniversityNode {
	universitiesMap := make([]UniversityNode, len(universities))
	for i, v := range universities {
		universitiesMap[i] = graphqlService.NewUniversityNodeFromUniversity(&v, countryCode)
	}
	return universitiesMap
}

func (graphqlService *UniversityGraphService) GetUniversityByGlobalId(encodedGlobalId string) (UniversityNode, error) {

	idParts := strings.Split(encodedGlobalId, ":")

	if len(idParts) != 2 {
		return UniversityNode{}, errors.New("Invalid global university Id, the country relay Id should be of the form University:{countryCode}:{universityId}")
	}

	countryCode := idParts[0]
	universityId := idParts[1]

	return graphqlService.GetUniversityNodeByCountryCodeAndUniversityId(universityId, countryCode)

}

func (graphqlService *UniversityGraphService) GetUniversityNodeByCountryCodeAndUniversityId(universityIdString string, countryCode string) (UniversityNode, error) {
	universityId, err := strconv.ParseInt(universityIdString, 10, 64)
	if err != nil {
		return UniversityNode{}, errors.New("Invalid university Id, university Id should be an Integer")
	}

	theUniversity, err := graphqlService.universityService.GetUniversityByIdAndCountryCode(universityId, countryCode)

	if err != nil {
		return UniversityNode{}, err
	}

	return graphqlService.NewUniversityNodeFromUniversity(&theUniversity, countryCode), nil
}

func (graphqlService *UniversityGraphService) GetAllUniversities(countryCode string) ([]UniversityNode, error) {
	if universities, err := graphqlService.universityService.GetAllUniversitiesForCountryCode(countryCode); err != nil {
		return []UniversityNode{}, nil
	} else {
		return graphqlService.mapUniversitiesToUniversityNodes(universities, countryCode), nil
	}
}

func (graphqlService *UniversityGraphService) CreateUniversity(university *university.University, countryCode string) error {
	return graphqlService.universityService.CreateUniversity(university, countryCode)
}

func (graphqlService *UniversityGraphService) UpdateUniversity(university *university.University, countryCode string) error {
	return graphqlService.universityService.UpdateUniversity(university, countryCode)
}

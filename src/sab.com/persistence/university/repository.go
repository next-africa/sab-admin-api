package university

import (
	"google.golang.org/appengine/datastore"
	"sab.com/domain/university"
	"sab.com/persistence"
	"sab.com/persistence/country"
	"strings"
)

const UNIVERSITY_KIND = "University"

type DatastoreUniversityRepository struct {
	contextStore *persistence.ContextStore
}

func NewUniversityRepository(contextStore *persistence.ContextStore) DatastoreUniversityRepository {
	return DatastoreUniversityRepository{contextStore}
}

func (repository DatastoreUniversityRepository) Save(universityToSave *university.University, countryCode string) error {
	gaeContext := repository.contextStore.GetContext()

	key := repository.getUniversityKey(universityToSave.Id, countryCode)

	if completeKey, err := datastore.Put(gaeContext, key, universityToSave); err != nil {
		return err
	} else {
		universityToSave.Id = completeKey.IntID()
		return nil
	}
}

func (repository DatastoreUniversityRepository) GetAll(countryCode string) ([]university.University, error) {
	gaeContext := repository.contextStore.GetContext()

	universities := make([]university.University, 0)

	countryKey := datastore.NewKey(gaeContext, country.COUNTRY_KIND, strings.ToLower(countryCode), 0, nil)

	keys, err := datastore.NewQuery(UNIVERSITY_KIND).Ancestor(countryKey).GetAll(gaeContext, &universities)

	if err != nil {
		return universities, err
	}

	for i := range universities {
		universities[i].Id = keys[i].IntID()
	}

	return universities, nil
}

func (repository DatastoreUniversityRepository) GetById(id int64, countryCode string) (university.University, error) {
	gaeContext := repository.contextStore.GetContext()

	var universityToReturn university.University

	universityKey := repository.getUniversityKey(id, countryCode)

	if err := datastore.Get(gaeContext, universityKey, &universityToReturn); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return universityToReturn, university.UniversityNotFoundError
		}
		return universityToReturn, err
	}

	universityToReturn.Id = id

	return universityToReturn, nil
}

func (repository DatastoreUniversityRepository) getUniversityKey(universityId int64, countryCode string) *datastore.Key {
	gaeContext := repository.contextStore.GetContext()
	countryKey := datastore.NewKey(gaeContext, country.COUNTRY_KIND, strings.ToLower(countryCode), 0, nil)

	return datastore.NewKey(gaeContext, UNIVERSITY_KIND, "", universityId, countryKey)
}

func (repository DatastoreUniversityRepository) HasUniversity(id int64, countryCode string) (bool, error) {
	gaeContext := repository.contextStore.GetContext()

	universityKey := repository.getUniversityKey(id, countryCode)

	var dst []university.University

	q, err := datastore.NewQuery(UNIVERSITY_KIND).Filter("__key__ =", universityKey).KeysOnly().GetAll(gaeContext, dst)

	return q != nil, err
}

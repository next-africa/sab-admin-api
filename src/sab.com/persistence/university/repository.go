package university

import (
	"google.golang.org/appengine/datastore"
	"sab.com/domain/university"
	"sab.com/persistence"
	"sab.com/persistence/country"
)

const UNIVERSITY_KIND = "University"

type universityRepository struct {
	contextStore *persistence.ContextStore
}

func NewUniversityRepository(contextStore *persistence.ContextStore) universityRepository {
	return universityRepository{contextStore}
}

func (repository universityRepository) Save(universityToSave *university.University, countryCode string) error {
	gaeContext := repository.contextStore.GetContext()
	countryKey := datastore.NewKey(gaeContext, country.COUNTRY_KIND, countryCode, 0, nil)

	key := datastore.NewKey(gaeContext, UNIVERSITY_KIND, "", universityToSave.Id, countryKey)

	if completeKey, err := datastore.Put(gaeContext, key, universityToSave); err != nil {
		return err
	} else {
		universityToSave.Id = completeKey.IntID()
		return nil
	}
}

func (repository universityRepository) GetAll(countryCode string) ([]university.University, error) {
	gaeContext := repository.contextStore.GetContext()

	universities := make([]university.University, 0)

	countryKey := datastore.NewKey(gaeContext, country.COUNTRY_KIND, countryCode, 0, nil)

	keys, err := datastore.NewQuery(UNIVERSITY_KIND).Ancestor(countryKey).GetAll(gaeContext, &universities)

	if err != nil {
		return universities, err
	}

	for i := range universities {
		universities[i].Id = keys[i].IntID()
	}

	return universities, nil
}

func (repository universityRepository) GetById(id int64, countryCode string) (university.University, error) {
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

func (repository universityRepository) getUniversityKey(universityId int64, countryCode string) *datastore.Key {
	gaeContext := repository.contextStore.GetContext()
	countryKey := datastore.NewKey(gaeContext, country.COUNTRY_KIND, countryCode, 0, nil)

	return datastore.NewKey(gaeContext, UNIVERSITY_KIND, "", universityId, countryKey)
}

func (repository universityRepository) HasUniversity(id int64, countryCode string) (bool, error) {
	gaeContext := repository.contextStore.GetContext()

	universityKey := repository.getUniversityKey(id, countryCode)

	var dst []university.University

	q, err := datastore.NewQuery(UNIVERSITY_KIND).Filter("__key__ =", universityKey).KeysOnly().GetAll(gaeContext, dst)

	return q != nil, err
}

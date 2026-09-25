package search

// personOrOrg is the Watchman types that OpenSanctions LegalEntity (and similar
// "unknown kind of legal person") records get encoded as. Vessel/aircraft stay
// hard-partitioned: a person query must not score a ship.
func personOrOrg(t EntityType) bool {
	switch t {
	case EntityPerson, EntityBusiness, EntityOrganization, EntityUnknown, emptyEntityType:
		return true
	default:
		return false
	}
}

// recastToType projects a person/business/organization query onto the index
// entity's type so Similarity can compare names and IDs instead of returning 0.
// Vessel and aircraft mismatches still fail. Same-type and empty-type queries
// are returned unchanged (no alloc).
func recastToType[Q any](e Entity[Q], typ EntityType) (Entity[Q], bool) {
	if e.Type == typ || e.Type == emptyEntityType || e.Type == EntityUnknown {
		return e, true
	}
	if !personOrOrg(e.Type) || !personOrOrg(typ) {
		return e, false
	}

	name, alts, ids := identityFrom(e)
	out := e
	out.Type = typ
	switch typ {
	case EntityPerson:
		p := Person{Name: name, AltNames: alts, GovernmentIDs: ids}
		if e.Person != nil {
			if p.Name == "" {
				p.Name = e.Person.Name
			}
			p.BirthDate = e.Person.BirthDate
			p.DeathDate = e.Person.DeathDate
			p.Gender = e.Person.Gender
			p.PlaceOfBirth = e.Person.PlaceOfBirth
			p.Titles = e.Person.Titles
			if len(p.GovernmentIDs) == 0 {
				p.GovernmentIDs = e.Person.GovernmentIDs
			}
		}
		out.Person = &p
	case EntityBusiness:
		out.Business = &Business{Name: name, AltNames: alts, GovernmentIDs: ids}
	case EntityOrganization:
		out.Organization = &Organization{Name: name, AltNames: alts, GovernmentIDs: ids}
	default:
		return e, false
	}
	return out, true
}

func identityFrom[T any](e Entity[T]) (name string, alts []string, ids []GovernmentID) {
	name = e.Name
	switch {
	case e.Person != nil:
		if name == "" {
			name = e.Person.Name
		}
		alts = e.Person.AltNames
		ids = e.Person.GovernmentIDs
	case e.Business != nil:
		if name == "" {
			name = e.Business.Name
		}
		alts = e.Business.AltNames
		ids = e.Business.GovernmentIDs
	case e.Organization != nil:
		if name == "" {
			name = e.Organization.Name
		}
		alts = e.Organization.AltNames
		ids = e.Organization.GovernmentIDs
	}
	return name, alts, ids
}

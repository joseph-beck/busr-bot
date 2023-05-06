package racing

import "log"

type Circuit int

const (
	UndefinedTrack Circuit = iota
	Bahrain
	SaudiArabia
	Australia
	Imola
	Miami
	Spain
	Monaco
	Azerbaijan
	Canada
	GreatBritain
	Austria
	France
	Hungary
	Belgium
	Netherlands
	Monza
	Singapore
	Japan
	Austin
	Mexico
	Brazil
	AbuDhabi
	Portugal
	China
)

func (c Circuit) Str() string {
	switch c {
	case UndefinedTrack:
		log.Panic("Undefined Track")
		return "err"
	case Bahrain:
		return "Bahrain"
	case SaudiArabia:
		return "Saudi Arabia"
	case Australia:
		return "Australia"
	case Imola:
		return "Imola"
	case Miami:
		return "Miami"
	case Spain:
		return "Spain"
	case Monaco:
		return "Monaco"
	case Azerbaijan:
		return "Azerbaijan"
	case Canada:
		return "Canada"
	case GreatBritain:
		return "Great Britain"
	case Austria:
		return "Austria"
	case France:
		return "France"
	case Hungary:
		return "Hungary"
	case Belgium:
		return "Belgium"
	case Netherlands:
		return "Netherlands"
	case Monza:
		return "Monza"
	case Singapore:
		return "Singapore"
	case Japan:
		return "Japan"
	case Austin:
		return "Austin"
	case Mexico:
		return "Mexico"
	case Brazil:
		return "Brazil"
	case AbuDhabi:
		return "Abu Dhabi"
	case Portugal:
		return "Portugal"
	case China:
		return "China"
	}
	return "unknown"
}

type Seasons int

const (
	UndefinedSeason Seasons = iota
	Spring
	Summer
	Winter
)

func (s Seasons) Str() string {
	switch s {
	case UndefinedSeason:
		log.Panic("Undefined Track")
		return "err"
	case Spring:
		return "Spring"
	case Summer:
		return "Summer"
	case Winter:
		return "Winter"
	}
	return "unknown"
}
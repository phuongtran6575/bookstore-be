package models

// ActorType defines who performed an action in the system
type ActorType string

const (
	ActorUser   ActorType = "user"
	ActorAdmin  ActorType = "admin"
	ActorSystem ActorType = "system"
)

// FeeType defines how fees are calculated
type FeeType string

const (
	FeeNone    FeeType = "none"
	FeePercent FeeType = "percent"
	FeeFixed   FeeType = "fixed"
)

package hero

type RoleType string

type IHero interface {
	GetHeroAbility() (attack, hp, defense int)
}

type Hero struct {
	Name                                   string
	Role                                   RoleType
	AttackPoint, HealthPoint, DefensePoint int
}

// Hero Role
const (
	RoleFigters  RoleType = "Figters"
	RoleAssassin RoleType = "Assassin"
	RoleMage     RoleType = "Mage"
	RoleMarksman RoleType = "Marksman"
	RoleSupport  RoleType = "Support"
	RoleTank     RoleType = "Tank"
)

func (hero Hero) GetHeroAbility() (attack, hp, defense int) {
	return hero.AttackPoint, hero.HealthPoint, hero.DefensePoint
}

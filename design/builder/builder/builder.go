package builder

const (
	FOOD  = "food"
	DRINK = "drink"
)

type Item interface {
	Price() float32
	Name() string
	Category() string
}

type Food struct {
}

func (Food) Price() float32 {
	return 0.0
}

func (Food) Name() string {
	return ""
}

func (Food) Category() string {
	return FOOD
}

type Drink struct {
}

func (Drink) Price() float32 {
	return 0.0
}

func (Drink) Name() string {
	return ""
}

func (Drink) Category() string {
	return DRINK
}

type Hamburger struct {
	Food
}

func (Hamburger) Price() float32 {
	return 12.00
}

func (Hamburger) Name() string {
	return "Hamburger"
}

type FriedChicken struct {
	Food
}

func (FriedChicken) Price() float32 {
	return 18.00
}

func (FriedChicken) Name() string {
	return "FriedChicken"
}

type Cola struct {
	Drink
}

func (Cola) Price() float32 {
	return 3.00
}

func (Cola) Name() string {
	return "Cola"
}

type Beer struct {
	Drink
}

func (Beer) Price() float32 {
	return 5.00
}

func (Beer) Name() string {
	return "Beer"
}

type Meal []Item

func (m *Meal) AddItem(item ...Item) {
	*m = append(*m, item...)
}

func (m Meal) Cost() (cost float32) {
	for _, val := range m {
		cost += val.Price()
	}
	return
}

func (m Meal) ShowItems() (msg string) {
	for _, val := range m {
		msg += "Category: " + val.Category() + "Name: " + val.Name() + "\n"
	}
	return
}

type MealBuilder struct {
}

func (MealBuilder) MealOne() (meal *Meal) {
	meal = new(Meal)
	meal.AddItem(new(FriedChicken), new(Beer))
	return
}

func (mb MealBuilder) MealTwo(food1, food2, drink Item) *Meal {
	meal := new(Meal)
	meal.AddItem(food1, food2, drink)
	return meal
}

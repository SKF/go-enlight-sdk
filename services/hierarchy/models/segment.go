package models

type IndustrySegment string

const (
	Agriculture           IndustrySegment = "agriculture"
	Construction          IndustrySegment = "construction"
	FoodAndBeverage       IndustrySegment = "food_and_beverage"
	HydrocarbonProcessing IndustrySegment = "hydrocarbon_processing"
	MachineTool           IndustrySegment = "machine_tool"
	Marine                IndustrySegment = "marine"
	Metal                 IndustrySegment = "metal"
	Mining                IndustrySegment = "mining"
	PowerGeneration       IndustrySegment = "power_generation"
	PulpAndPaper          IndustrySegment = "pulp_and_paper"
	Renewable             IndustrySegment = "renewable"
	Undefined             IndustrySegment = "undefined"
)

func (seg IndustrySegment) String() string {
	return string(seg)
}

func (seg IndustrySegment) Title() string {
	switch seg {
	case Agriculture:
		return "Agriculture"
	case Construction:
		return "Construction"
	case FoodAndBeverage:
		return "Food & Beverage"
	case HydrocarbonProcessing:
		return "Hydrocarbon Processing"
	case MachineTool:
		return "Machine Tool"
	case Marine:
		return "Marine"
	case Metal:
		return "Metal"
	case Mining:
		return "Mining"
	case PowerGeneration:
		return "Power Generation"
	case PulpAndPaper:
		return "Pulp & Paper"
	case Renewable:
		return "Renewable"
	case Undefined:
		return "Undefined"
	}
	return "Invalid"
}

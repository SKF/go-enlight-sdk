package models

import "fmt"

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

var allSegments = []IndustrySegment{
	Agriculture, Construction, FoodAndBeverage, HydrocarbonProcessing, MachineTool,
	Marine, Metal, Mining, PowerGeneration, PulpAndPaper, Renewable, Undefined,
}

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

func (seg IndustrySegment) Validate() error {
	for _, segment := range allSegments {
		if seg == segment {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid industry segment", seg)
}

package vtex

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type SelectedFacetInput struct {
	Key   string
	Value string
}

type SimulationBehavior string

const (
	SimulationBehaviorDefault SimulationBehavior = "default"
	SimulationBehaviorOnly1P  SimulationBehavior = "only1P"
	SimulationBehaviorSkip    SimulationBehavior = "skip"
	SimulationBehaviorAsync   SimulationBehavior = "async"
)

type ProductSearchInput struct {
	From                 int
	To                   int
	FullText             string
	OrderBy              string
	PriceRange           string
	SalesChannel         string
	HideUnavailableItems bool
	SimulationBehavior   SimulationBehavior
	SelectedFacets       []SelectedFacetInput
}

type ProductSearchQuery struct {
	ProductSearch struct {
		RecordsFiltered int
		Products        []ProductFragment
	}
}

var GraphQLClient = graphql.NewClient("https://agenciam3.myvtex.com/_v/private/graphql/v1")

func ProductSearch(input ProductSearchInput) ProductSearchQuery {

	response := ProductSearchQuery{}

	req := graphql.NewRequest(fmt.Sprintf(`
	%s
	query ProductSearch(
		$from: Int = 0
		$to: Int = 9
		$fullText: String
		$orderBy: String = "OrderByScoreDESC"
		$priceRange: String
		$salesChannel: String
		$hideUnavailableItems: Boolean = false
		$simulationBehavior: SimulationBehavior = default
		$selectedFacets: [SelectedFacetInput]
	) {
		productSearch(
			from: $from
			fullText: $fullText
			hideUnavailableItems: $hideUnavailableItems
			orderBy: $orderBy
			priceRange: $priceRange
			salesChannel: $salesChannel
			selectedFacets: $selectedFacets
			simulationBehavior: $simulationBehavior
			to: $to
		) @context(provider: "vtex.search-graphql") {
			recordsFiltered
			products {
				...ProductFragment
			}
		}
	}
	`, PRODUCT_FRAGMENT))

	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	if err := GraphQLClient.Run(ctx, req, &response); err != nil {
		log.Fatal(err)
	}

	return response
}

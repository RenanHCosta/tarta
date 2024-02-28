package vtex

const PRODUCT_FRAGMENT = `
fragment ProductFragment on Product {
	productId
	productName
	items {
		itemId
		name
		images {
			imageUrl
		}
		sellers {
			sellerId
			commertialOffer {
				ListPrice
				Price
				Installments {
					NumberOfInstallments
					Value
				}
				AvailableQuantity
				discountHighlights {
					name
				}
			}
		}
		variations {
			name
			values
		}
	}
	brand
	categoryTree {
		href
		name
	}
	productReference
	description
	properties {
		name
		values
	}
	clusterHighlights {
		id
		name
	}
}
`

type (
	Image struct {
		ImageUrl string
	}
	Installment struct {
		NumberOfInstallments int     `graphql:"NumberOfInstallments"`
		Value                float32 `graphql:"Value"`
	}
	Discount struct {
		Name string
	}
	CommertialOffer struct {
		ListPrice          float32       `graphql:"ListPrice"`
		Price              float32       `graphql:"Price"`
		AvailableQuantity  int           `graphql:"AvailableQuantity"`
		Installments       []Installment `graphql:"Installments"`
		DiscountHighlights []Discount
	}
	Variation struct {
		Name   string
		Values []string
	}
	Seller struct {
		SellerId        string
		CommertialOffer CommertialOffer
	}
	ItemFragment struct {
		ItemId     string
		Name       string
		Images     []Image
		Variations []Variation
		Sellers    []Seller
	}
	Category struct {
		Href string
		Name string
	}
	Property struct {
		Name   string
		Values []string
	}
	ClusterHighlight struct {
		Id   string
		Name string
	}
	ProductFragment struct {
		ProductId         string
		ProductName       string
		Brand             string
		ProductReference  string
		Description       string
		Items             []ItemFragment
		CategoryTree      []Category
		Properties        []Property
		ClusterHighlights []ClusterHighlight
	}
)

func (product ProductFragment) GetSelectedItem() ItemFragment {
	return product.Items[0]
}

func (product ProductFragment) GetSelectedSeller() Seller {
	return product.GetSelectedItem().Sellers[0]
}

func (product ProductFragment) GetHighestInstallment() Installment {

	highestInstallment := product.GetSelectedSeller().CommertialOffer.Installments[0]

	for i := range product.GetSelectedSeller().CommertialOffer.Installments {
		if product.GetSelectedSeller().CommertialOffer.Installments[i].NumberOfInstallments > highestInstallment.NumberOfInstallments {
			highestInstallment = product.GetSelectedSeller().CommertialOffer.Installments[i]
		}
	}

	return highestInstallment
}

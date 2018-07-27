package cart

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sort"

	"flamingo.me/flamingo-commerce/product/domain"
	"flamingo.me/flamingo/framework/flamingo"
)

type (
	// DecoratedCartFactory - Factory to be injected: If you need to create a new Decorator then get the factory injected and use the factory
	DecoratedCartFactory struct {
		ProductService domain.ProductService `inject:""`
		Logger         flamingo.Logger       `inject:""`
	}

	// DecoratedCart Decorates Access To a Cart
	DecoratedCart struct {
		Cart           Cart
		DecoratedItems []DecoratedCartItem
		Ctx            context.Context `json:"-"`
	}

	// DecoratedCartItem Decorates a CartItem with its Product
	GroupedDecoratedCartItem struct {
		DecoratedItems []DecoratedCartItem
		Group          string
	}

	// DecoratedCartItem Decorates a CartItem with its Product
	DecoratedCartItem struct {
		Item    Item
		Product domain.BasicProduct
	}
)

// Create Factory method to get Decorated Cart
func (df *DecoratedCartFactory) Create(ctx context.Context, Cart Cart) *DecoratedCart {
	decoratedCart := DecoratedCart{Cart: Cart}
	decoratedCart.DecoratedItems = df.CreateDecorateCartItems(ctx, Cart.Cartitems)
	decoratedCart.Ctx = ctx
	return &decoratedCart
}

// CreateDecorateCartItems Factory method to get Decorated Cart
func (df *DecoratedCartFactory) CreateDecorateCartItems(ctx context.Context, items []Item) []DecoratedCartItem {
	var decoratedItems []DecoratedCartItem
	for _, cartitem := range items {
		decoratedItem := df.decorateCartItem(ctx, cartitem)
		decoratedItems = append(decoratedItems, decoratedItem)
	}
	return decoratedItems
}

//decorateCartItem factory method
func (df *DecoratedCartFactory) decorateCartItem(ctx context.Context, cartitem Item) DecoratedCartItem {
	decorateditem := DecoratedCartItem{Item: cartitem}
	product, e := df.ProductService.Get(ctx, cartitem.MarketplaceCode)
	if e != nil {
		df.Logger.Error("cart.decorator - no product for item", e)
		if product == nil {
			//To avoid errors if consumers want to access the product data
			product = domain.SimpleProduct{
				BasicProductData: domain.BasicProductData{
					Title: cartitem.ProductName + "[outdated]",
				},
			}
		}
		return decorateditem
	}
	if product.Type() == domain.TYPECONFIGURABLE {
		if configureable, ok := product.(domain.ConfigurableProduct); ok {
			configurableWithVariant, err := configureable.GetConfigurableWithActiveVariant(cartitem.VariantMarketPlaceCode)
			if err != nil {
				product = domain.SimpleProduct{
					BasicProductData: domain.BasicProductData{
						Title: cartitem.ProductName + "[outdated]",
					},
				}
			} else {
				product = configurableWithVariant
			}
		}
	}
	decorateditem.Product = product
	return decorateditem
}

// IsConfigurable - checks if current CartItem is a Configurable Product
func (dci DecoratedCartItem) IsConfigurable() bool {
	return dci.Product.Type() == domain.TYPECONFIGURABLE_WITH_ACTIVE_VARIANT
}

// GetVariant getter
func (dci DecoratedCartItem) GetVariant() (*domain.Variant, error) {
	return dci.Product.(domain.ConfigurableProductWithActiveVariant).Variant(dci.Item.VariantMarketPlaceCode)
}

// GetDisplayTitle getter
func (dci DecoratedCartItem) GetDisplayTitle() string {
	if dci.IsConfigurable() {
		variant, e := dci.GetVariant()
		if e != nil {
			return "Error Getting Variant"
		}
		return variant.Title
	}
	return dci.Product.BaseData().Title
}

// GetDisplayMarketplaceCode getter
func (dci DecoratedCartItem) GetDisplayMarketplaceCode() string {
	if dci.IsConfigurable() {
		variant, e := dci.GetVariant()
		if e != nil {
			return "Error Getting Variant"
		}
		return variant.MarketPlaceCode
	}
	return dci.Product.BaseData().MarketPlaceCode
}

// GetVariantsVariationAttribute getter
func (dci DecoratedCartItem) GetVariantsVariationAttributes() domain.Attributes {
	attributes := domain.Attributes{}
	if dci.IsConfigurable() {
		variant, _ := dci.GetVariant()

		for _, attributeName := range dci.Product.(domain.ConfigurableProductWithActiveVariant).VariantVariationAttributes {
			attributes[attributeName] = variant.BaseData().Attributes[attributeName]
		}
	}
	log.Println(attributes)
	return attributes
}

// GetVariantsVariationAttribute getter
func (dci DecoratedCartItem) GetVariantsVariationAttributeCodes() []string {
	if dci.Product.Type() == domain.TYPECONFIGURABLE_WITH_ACTIVE_VARIANT {
		return dci.Product.(domain.ConfigurableProductWithActiveVariant).VariantVariationAttributes
	}
	return nil
}

// GetGroupedBy getter
func (dc DecoratedCart) GetGroupedBy(group string, sortGroup bool) []*GroupedDecoratedCartItem {
	groupedItemsCollection := make(map[string]*GroupedDecoratedCartItem)
	var groupedItemsCollectionKeys []string

	var groupKey string
	for _, item := range dc.DecoratedItems {
		switch group {
		case "retailer_code":
			groupKey = item.Product.BaseData().RetailerCode
		case "deliveryIntent":
			groupKey = item.Item.OriginalDeliveryIntent.String()
		default:
			groupKey = "default"
		}
		if _, ok := groupedItemsCollection[groupKey]; !ok {
			groupedItemsCollection[groupKey] = &GroupedDecoratedCartItem{
				Group: groupKey,
			}
			groupedItemsCollectionKeys = append(groupedItemsCollectionKeys, groupKey)
		}

		if groupedItemsEntry, ok := groupedItemsCollection[groupKey]; ok {
			groupedItemsEntry.DecoratedItems = append(groupedItemsEntry.DecoratedItems, item)
		}
	}

	//sort before return
	if sortGroup {
		sort.Strings(groupedItemsCollectionKeys)
	}

	var groupedItemsCollectionSorted []*GroupedDecoratedCartItem
	for _, keyName := range groupedItemsCollectionKeys {
		if groupedItemsEntry, ok := groupedItemsCollection[keyName]; ok {
			groupedItemsCollectionSorted = append(groupedItemsCollectionSorted, groupedItemsEntry)
		}
	}
	return groupedItemsCollectionSorted
}

func (dc DecoratedCart) GetByItemId(itemId string) (*DecoratedCartItem, error) {
	for _, item := range dc.DecoratedItems {
		if item.Item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("itemId %v not existend not existent in decorated cart", itemId))
}

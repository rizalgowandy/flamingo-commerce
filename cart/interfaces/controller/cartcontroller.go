package controller

import (
	"context"
	"encoding/gob"
	"log"
	"strconv"

	"flamingo.me/flamingo-commerce/cart/application"
	cartDomain "flamingo.me/flamingo-commerce/cart/domain/cart"
	productDomain "flamingo.me/flamingo-commerce/product/domain"
	"flamingo.me/flamingo/framework/router"
	"flamingo.me/flamingo/framework/web"
	"flamingo.me/flamingo/framework/web/responder"
)

type (
	// CartViewData is used for cart views/templates
	CartViewData struct {
		DecoratedCart         cartDomain.DecoratedCart
		CartValidationResult  cartDomain.CartValidationResult
		AddToCartProductsData []productDomain.BasicProductData
	}

	// CartViewController for carts
	CartViewController struct {
		responder.RenderAware   `inject:""`
		responder.RedirectAware `inject:""`

		ApplicationCartService         *application.CartService         `inject:""`
		ApplicationCartReceiverService *application.CartReceiverService `inject:""`
		Router                         *router.Router                   `inject:""`

		ShowEmptyCartPageIfNoItems bool `inject:"config:cart.showEmptyCartPageIfNoItems,optional"`
	}

	CartViewActionData struct {
		AddToCartProductsData []productDomain.BasicProductData
	}
)

func init() {
	gob.Register(CartViewActionData{})
}

// ViewAction the DecoratedCart View ( / cart)
func (cc *CartViewController) ViewAction(ctx context.Context, r *web.Request) web.Response {
	decoratedCart, err := cc.ApplicationCartReceiverService.ViewDecoratedCart(ctx, r.Session().G())
	if err != nil {
		log.Printf("cart.cartcontroller.viewaction: Error %v", err)
		return cc.Render(ctx, "checkout/carterror", nil)
	}

	if cc.ShowEmptyCartPageIfNoItems && decoratedCart.Cart.ItemCount() == 0 {
		return cc.Render(ctx, "checkout/emptycart", nil)
	}

	cartViewData := CartViewData{
		DecoratedCart:        *decoratedCart,
		CartValidationResult: cc.ApplicationCartService.ValidateCart(ctx, r.Session().G(), decoratedCart),
	}

	flashes := r.Session().Flashes("cart.view.data")
	if len(flashes) > 0 {
		if cartViewActionData, ok := flashes[0].(CartViewActionData); ok {
			cartViewData.AddToCartProductsData = cartViewActionData.AddToCartProductsData
		}
	}

	return cc.Render(ctx, "checkout/cart", cartViewData)

}

// AddAndViewAction the DecoratedCart View ( / cart)
func (cc *CartViewController) AddAndViewAction(ctx context.Context, r *web.Request) web.Response {
	variantMarketplaceCode, _ := r.Param1("variantMarketplaceCode")

	qty, ok := r.Param1("qty")
	if !ok {
		qty = "1"
	}

	qtyInt, _ := strconv.Atoi(qty)
	deliveryCode, _ := r.Param1("deliveryCode")

	addRequest := cc.ApplicationCartService.BuildAddRequest(ctx, r.MustParam1("marketplaceCode"), variantMarketplaceCode, qtyInt)

	err, product := cc.ApplicationCartService.AddProduct(ctx, r.Session().G(), deliveryCode, addRequest)
	if notAllowedErr, ok := err.(*cartDomain.AddToCartNotAllowed); ok {
		if notAllowedErr.RedirectHandlerName != "" {
			return cc.Redirect(notAllowedErr.RedirectHandlerName, notAllowedErr.RedirectParams)
		}
	}
	if err != nil {
		log.Printf("cart.cartcontroller.addandviewaction: Error %v", err)
		return cc.Render(ctx, "checkout/carterror", nil)
	}

	return cc.Redirect("cart.view", nil).With("cart.view.data", CartViewActionData{
		AddToCartProductsData: []productDomain.BasicProductData{product.BaseData()},
	})
}

// UpdateQtyAndViewAction the DecoratedCart View ( / cart)
func (cc *CartViewController) UpdateQtyAndViewAction(ctx context.Context, r *web.Request) web.Response {

	id, ok := r.Param1("id")
	if !ok {
		log.Printf("cart.cartcontroller.UpdateQtyAndViewAction: param id not found")
		return cc.Redirect("cart.view", nil)
	}

	qty, ok := r.Param1("qty")
	if !ok {
		qty = "1"
	}

	qtyInt, err := strconv.Atoi(qty)
	if err != nil {
		qtyInt = 1
	}
	deliveryCode, _ := r.Param1("deliveryCode")

	err = cc.ApplicationCartService.UpdateItemQty(ctx, r.Session().G(), id, deliveryCode, qtyInt)

	if err != nil {
		log.Printf("cart.cartcontroller.UpdateAndViewAction: Error %v", err)
	}

	return cc.Redirect("cart.view", nil)
}

// DeleteAndViewAction the DecoratedCart View ( / cart)
func (cc *CartViewController) DeleteAndViewAction(ctx context.Context, r *web.Request) web.Response {

	id, ok := r.Param1("id")
	if !ok {
		log.Printf("cart.cartcontroller.deleteaction: param id not found")
		return cc.Redirect("cart.view", nil)
	}
	deliveryCode, _ := r.Param1("deliveryCode")
	err := cc.ApplicationCartService.DeleteItem(ctx, r.Session().G(), id, deliveryCode)
	if err != nil {
		log.Printf("cart.cartcontroller.deleteaction: Error %v", err)
	}

	return cc.Redirect("cart.view", nil)
}

// DeleteAllAndViewAction the empty
func (cc *CartViewController) DeleteAllAndViewAction(ctx context.Context, r *web.Request) web.Response {
	err := cc.ApplicationCartService.DeleteAllItems(ctx, r.Session().G())
	if err != nil {
		log.Printf("cart.cartcontroller.deleteaction: Error %v", err)
	}

	return cc.Redirect("cart.view", nil)
}

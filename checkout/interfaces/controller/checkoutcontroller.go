package controller

import (
	"log"

	"go.aoe.com/flamingo/core/breadcrumbs"
	"go.aoe.com/flamingo/core/cart/application"
	"go.aoe.com/flamingo/core/cart/domain/cart"

	"errors"

	application2 "go.aoe.com/flamingo/core/checkout/application"
	"go.aoe.com/flamingo/core/checkout/interfaces/controller/formDto"
	formApplicationService "go.aoe.com/flamingo/core/form/application"
	"go.aoe.com/flamingo/core/form/domain"
	"go.aoe.com/flamingo/framework/router"
	"go.aoe.com/flamingo/framework/web"
	"go.aoe.com/flamingo/framework/web/responder"
)

type (
	CheckoutViewData struct {
		DecoratedCart cart.DecoratedCart
		Form          domain.Form
		ErrorMessage  string
		HasError      bool
	}

	SuccessViewData struct {
		OrderId string
		Email   string
	}

	CheckoutController struct {
		responder.RenderAware   `inject:""`
		responder.RedirectAware `inject:""`
		ApplicationCartService  application.CartService     `inject:""`
		PaymentService          application2.PaymentService `inject:""`
		Router                  *router.Router              `inject:""`
	}
)

func (cc *CheckoutController) SubmitAction(ctx web.Context) web.Response {

	decoratedCart, e := cc.ApplicationCartService.GetDecoratedCart(ctx)
	if e != nil {
		log.Printf("cart.checkoutcontroller.viewaction: Error %v", e)
		return cc.Render(ctx, "checkout/carterror", nil)
	}

	breadcrumbs.Add(ctx, breadcrumbs.Crumb{
		Title: "Shopping Bag",
		Url:   cc.Router.URL("cart.view", nil).String(),
	})
	breadcrumbs.Add(ctx, breadcrumbs.Crumb{
		Title: "Reserve & Collect",
		Url:   cc.Router.URL("checkout.start", nil).String(),
	})

	form, e := formApplicationService.ProcessFormRequest(ctx, new(formDto.CheckoutFormService))
	log.Printf("lkjlklkj %#v %#v", form, e)
	log.Printf("lkjlklkj %#v %#v", form, e)
	// return on error (template need to handle error display)
	if e != nil {
		return cc.Render(ctx, "checkout/checkout", CheckoutViewData{
			DecoratedCart: decoratedCart,
			Form:          form,
		})
	}
	if form.IsValidAndSubmitted() {
		if checkoutFormData, ok := form.Data.(formDto.CheckoutFormData); ok {
			orderId, err := cc.placeOrder(ctx, checkoutFormData, decoratedCart.Cart)
			if err != nil {
				return cc.Render(ctx, "checkout/checkout", CheckoutViewData{
					DecoratedCart: decoratedCart,
					HasError:      true,
					Form:          form,
					ErrorMessage:  err.Error(),
				})
			}
			shippingEmail := checkoutFormData.ShippingAddress.Email
			if shippingEmail == "" {
				shippingEmail = checkoutFormData.BillingAddress.Email
			}
			return cc.Redirect("checkout.success", nil).With("checkout.success.data", SuccessViewData{
				OrderId: orderId,
				Email:   shippingEmail,
			})
		}

	}

	return cc.Render(ctx, "checkout/checkout", CheckoutViewData{
		DecoratedCart: decoratedCart,
		HasError:      false,
		Form:          form,
	})
}

func (cc *CheckoutController) SuccessAction(ctx web.Context) web.Response {
	flashes := ctx.Session().Flashes("checkout.success.data")
	if len(flashes) > 0 {
		return cc.Render(ctx, "checkout/success", flashes[0].(SuccessViewData))
	}

	return cc.Render(ctx, "checkout/expired", nil)
}

func (cc *CheckoutController) placeOrder(ctx web.Context, checkoutFormData formDto.CheckoutFormData, cart cart.Cart) (string, error) {

	billingAddress, shippingAddress := formDto.MapAddresses(checkoutFormData)
	log.Printf("Checkoutcontroller.submit - Info: billingAddress: %#v", billingAddress)
	err := cart.SetShippingInformation(ctx, shippingAddress, billingAddress, "flatrate", "flatrate")
	if err != nil {
		log.Printf("Error during place Order: %v", err)
		return "", errors.New("Error while setting shipping informations.")
	}
	orderid, err := cart.PlaceOrder(ctx, cc.PaymentService.GetPayment())

	if err != nil {
		log.Printf("Error during place Order: %v", err)
		return "", errors.New("Error while placing the order.")
	}
	return orderid, nil
}

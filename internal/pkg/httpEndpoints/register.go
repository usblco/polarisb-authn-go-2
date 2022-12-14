package httpEndpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/usblco/polarisb-authn-go-2/pkg"
	"github.com/usblco/polarisb-authn-go-2/pkg/contracts"
)

func (endpointApp *Endpoints) RegisterEndpoint(c *gin.Context) {
	// check if user if authorized to create users
	if checkIfAuthorizedByRole("admin", endpointApp.Actions, c) {
		// Bind the request to the struct
		request := &contracts.RegisterEndpointReceiveUserInfoDto{}
		// todo: handle error
		c.BindJSON(request)

		// Validate the request first.
		if request.Email == "" || request.Password == "" {
			// Return json response
			c.JSON(400, contracts.RegisterEndpointReturnDto{
				Err: "email and password are required",
			})
			return
		}

		// create the user
		_, state, _ := endpointApp.Actions.UserCreate(request.Email, request.Password)
		if state != pkg.UserCreated {
			if state == pkg.UserAlreadyExists {
				// Return json response
				c.JSON(409, contracts.RegisterEndpointReturnDto{
					Err: "user already exists",
				})
				return
			}
			// Return json response
			c.JSON(500, contracts.RegisterEndpointReturnDto{
				Err: "user could not be created",
			})
			return
		}

		// Return json response
		c.JSON(200, contracts.RegisterEndpointReturnDto{
			Message: "user created",
		})
	} else {
		// Return json response
		c.JSON(401, contracts.RegisterEndpointReturnDto{
			Err: "not authorized",
		})
	}
}

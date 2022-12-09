/*
Package hack contains a frontend that provides a /metadata endpoint for the rootio hub action.
It is not intended to be long lived and will be removed as we migrate to exposing Hardware
data to Tinkerbell templates. In doing so, we can convert the rootio action to accept its inputs
via parameters instead of retrieing them from Hegel and subsequently delete this frontend.
*/
package hack

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinkerbell/hegel/internal/http/request"
)

// Client is a backend for retrieving hack instance data.
type Client interface {
	GetHackInstance(ctx context.Context, ip string) (Instance, error)
}

// Instance is a representation of the instance metadata. Its based on the rooitio hub action
// and should have just enough information for it to work.
type Instance struct {
	Metadata struct {
		Instance struct {
			Storage struct {
				Disks []struct {
					Device     string `json:"device"`
					Partitions []struct {
						Label  string `json:"label"`
						Number int    `json:"number"`
						Size   uint64 `json:"size"`
					} `json:"partitions"`
					WipeTable bool `json:"wipe_table"`
				} `json:"disks"`
				Filesystems []struct {
					Mount struct {
						Create struct {
							Options []string `json:"options"`
						} `json:"create"`
						Device string `json:"device"`
						Format string `json:"format"`
						Point  string `json:"point"`
					} `json:"mount"`
				} `json:"filesystems"`
			} `json:"storage"`
		} `json:"instance"`
	} `json:"metadata"`
}

// Configure configures router with a `/metadata` endpoint using client to retrieve instance data.
func Configure(router gin.IRouter, client Client) {
	router.GET("/metadata", func(ctx *gin.Context) {
		ip, err := request.RemoteAddrIP(ctx.Request)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid remote address"))
		}

		instance, err := client.GetHackInstance(ctx, ip)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(200, instance)
	})
}

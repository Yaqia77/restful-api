package http

import (
	"net/http"
	"restful-api/apps/host"
	"restful-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()

	if err := c.Bind(ins); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, ins)
}

func (h *Handler) queryHost(c *gin.Context) {
	req := host.NewQueryHostFromHTTP(c.Request)

	if err := c.BindQuery(req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, result)
}

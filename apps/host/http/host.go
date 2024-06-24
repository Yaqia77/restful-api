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

func (h *Handler) describeHost(c *gin.Context) {
	id := host.NewDescribeHostRequestWithId(c.Param("id"))
	if err := c.Bind(id); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.svc.DescribeHost(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, result)
}

func (h *Handler) deleteHost(c *gin.Context) {
	id := host.NewDeleteHostRequestWithId(c.Param("id"))
	if err := c.Bind(id); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.svc.DeleteHost(c.Request.Context(), id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, nil)
}
func (h *Handler) putHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewPutUpdateHostRequest(c.Param("id"))

	// 解析Body里面的数据
	if err := c.Bind(req.Host); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Id = c.Param("id")

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, set)
}

func (h *Handler) patchHost(c *gin.Context) {
	req := host.NewPatchUpdateHostRequest(c.Param("id"))

	// 解析Body里面的数据
	if err := c.Bind(req.Host); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Id = c.Param("id")

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, set)

}

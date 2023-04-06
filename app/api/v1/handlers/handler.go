package handlers

import (
	"TT-Micro-Backend-Destiny/pkg/dto"
	"net/http"

	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Retrieve(ctx *gin.Context) {
	ID := ctx.Param("id")
	// make redis key to store and retrieve
	redisKey := fmt.Sprintf("%s:%s", "Destiny", ID)

	// get the response in cache
	response, rediserr := h.redisCache.Get(ctx, redisKey)
	if rediserr != nil {
		ctx.JSON(http.StatusInternalServerError, rediserr.Error())
		return
	}
	if response != nil {
		// cache exist
		var newRetrieveDestinyResponse dto.RetrieveDestinyResponse
		jsonerr := json.Unmarshal(response, &newRetrieveDestinyResponse)
		if jsonerr != nil {
			ctx.JSON(http.StatusInternalServerError, jsonerr.Error())
			return
		}
		ctx.IndentedJSON(http.StatusOK, newRetrieveDestinyResponse)
		return
	}
	// cache does not exist and retrieve Destiny data from db
	newRetrieveDestinyResponse, err := h.DestinyService.RetrieveDestiny(ctx, ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	data, jsonerr := json.Marshal(newRetrieveDestinyResponse)
	if jsonerr != nil {
		ctx.JSON(http.StatusInternalServerError, jsonerr.Error())
		return
	}

	rediserr = h.redisCache.Set(ctx, redisKey, string(data), 10*time.Second)
	if rediserr != nil {
		ctx.JSON(http.StatusInternalServerError, rediserr.Error())
		return
	}

	ctx.JSON(http.StatusOK, newRetrieveDestinyResponse)
}

func (h *Handler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var newUpdateDestinyRequest dto.UpdateDestinyRequest

	jsonerr := ctx.BindJSON(&newUpdateDestinyRequest)
	if jsonerr != nil {
		ctx.JSON(http.StatusBadRequest, jsonerr.Error())
		return
	}
	newUpdateDestinyResponse, err := h.DestinyService.UpdateDestiny(ctx, newUpdateDestinyRequest, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, newUpdateDestinyResponse)
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	newDeleteDestinyResponse, err := h.DestinyService.DeleteDestiny(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, newDeleteDestinyResponse)
}

func (h *Handler) Create(ctx *gin.Context) {
	var newAddDestinyRequest dto.AddDestinyRequest

	jsonerr := ctx.BindJSON(&newAddDestinyRequest)
	if jsonerr != nil {
		ctx.JSON(http.StatusBadRequest, jsonerr.Error())
		return
	}

	newAddDestinyResponse, err := h.DestinyService.AddDestiny(ctx, newAddDestinyRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, newAddDestinyResponse)
}

func (h *Handler) List(ctx *gin.Context) {
	// get the response in cache
	response, rediserr := h.redisCache.Get(ctx, "ListDestiny")
	if rediserr != nil {
		ctx.JSON(http.StatusInternalServerError, rediserr.Error())
		return
	}
	if response != nil {
		// cache exist
		var newListDestinyResponse dto.ListDestinyResponse
		jsonerr := json.Unmarshal(response, &newListDestinyResponse)
		if jsonerr != nil {
			ctx.JSON(http.StatusInternalServerError, jsonerr.Error())
			return
		}
		ctx.IndentedJSON(http.StatusOK, newListDestinyResponse)
		return
	}
	newListDestinyResponse, err := h.DestinyService.ListDestiny(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	data, jsonerr := json.Marshal(newListDestinyResponse)
	if jsonerr != nil {
		ctx.JSON(http.StatusInternalServerError, jsonerr.Error())
		return
	}

	rediserr = h.redisCache.Set(ctx, "ListDestiny", string(data), 10*time.Second)
	if rediserr != nil {
		ctx.JSON(http.StatusInternalServerError, rediserr.Error())
		return
	}
	ctx.JSON(http.StatusOK, newListDestinyResponse)
}

package handler

import (
	"ApiGateway/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func AdminProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var client http.Client
		var data entities.AdminResponse
		var baseURL = "https://abbiyudha.cloud.okteto.net/admin"

		request, err := http.NewRequest("GET", baseURL, nil)
		jwtToken := c.Get("user").(*jwt.Token)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken.Raw))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}

		return c.JSON(http.StatusOK, data)
	}

}

func CreateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			client  http.Client
			data    CreateAdminRequest
			baseURL = "https://abbiyudha.cloud.okteto.net/create/admin"
		)
		c.Bind(&data)

		var bodyJson, _ = json.Marshal(data)
		var body = bytes.NewBufferString(string(bodyJson))

		request, err := http.NewRequest("POST", baseURL, body)
		request.Header.Set("Content-Type", "application/json;")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		var resultResponse CreateAdminResponse
		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &resultResponse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultResponse)
		}

		return c.JSON(http.StatusOK, resultResponse)
	}

}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			client  http.Client
			data    LoginAdminRequest
			baseURL = "https://abbiyudha.cloud.okteto.net/login"
		)
		c.Bind(&data)

		var bodyJson, _ = json.Marshal(data)
		var body = bytes.NewBufferString(string(bodyJson))

		request, err := http.NewRequest("POST", baseURL, body)
		request.Header.Set("Content-Type", "application/json;")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		var resultResponse LoginAdminResponse
		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &resultResponse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultResponse)
		}

		return c.JSON(http.StatusOK, resultResponse)
	}

}

func UpdateAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			client  http.Client
			data    UpdateAdminRequest
			baseURL = "https://abbiyudha.cloud.okteto.net/update"
		)
		c.Bind(&data)

		var bodyJson, _ = json.Marshal(data)
		var body = bytes.NewBufferString(string(bodyJson))

		request, err := http.NewRequest("PUT", baseURL, body)
		jwtToken := c.Get("user").(*jwt.Token)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken.Raw))
		request.Header.Set("Content-Type", "application/json;")
		request.Header.Set("Content-Type", "application/json;")

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error 1")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		var resultResponse UpdateAdminResponse
		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &resultResponse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultResponse)
		}

		return c.JSON(http.StatusOK, resultResponse)
	}

}

func DeleteAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var client http.Client
		var data entities.Admin
		var baseURL = "https://abbiyudha.cloud.okteto.net/delete"

		request, err := http.NewRequest("DELETE", baseURL, nil)
		jwtToken := c.Get("user").(*jwt.Token)
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken.Raw))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}
		response, err := client.Do(request)
		if err != nil {
			return c.JSON(http.StatusOK, err)
		}
		defer response.Body.Close()

		result, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(result, &data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "server error")
		}

		return c.JSON(http.StatusOK, "succses")
	}

}

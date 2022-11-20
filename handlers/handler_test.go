package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ahmedkhaeld/recipes-api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes", ts.URL))
	defer resp.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	//assert.Equal(t, len(recipes), 10)
}

func TestNewRecipesHandler(t *testing.T) {
	r := SetupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	recipe := models.Recipe{
		Name: "New York Pizza",
	}
	jsonValue, _ := json.Marshal(recipe)
	req, _ := http.NewRequest("POST", "/recipes", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateRecipeHandler(t *testing.T) {
	r := SetupRouter()
	//r.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)

	ID := "c0283p3d0cvuglq85lpg"
	objID, _ := primitive.ObjectIDFromHex(ID)
	recipe := models.Recipe{
		ID:   objID,
		Name: "Gnocchi",
		Ingredients: []string{
			"5 large Idaho potatoes",
			"2 egges",
			"3/4 cup grated Parmesan",
			"3 1/2 cup all-purpose flour",
		},
	}
	jsonValue, _ := json.Marshal(recipe)

	stringObjectID := recipe.ID.Hex()
	reqFound, _ := http.NewRequest("PUT", "/recipes/:"+stringObjectID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("PUT", "/recipes/", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteRecipeHandler(t *testing.T) {
	r := SetupRouter()
	ts := httptest.NewServer(r)
	defer ts.Close()

	ID := "c0283p3d0cvuglq85lpg"
	objID, _ := primitive.ObjectIDFromHex(ID)
	stringObjectID := objID.Hex()

	msg := Msg{
		Message: "Recipe has been deleted",
	}
	jsonValue, _ := json.Marshal(msg)

	reqFound, err := http.NewRequest("DELETE", "/recipes/"+stringObjectID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	data, _ := ioutil.ReadAll(w.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Recipe has been deleted")
}

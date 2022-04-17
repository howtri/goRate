package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/howtri/goRate/skills"
)

// AddTodoHandler adds a new todo to the todo list
func AddSkillHandler(c *gin.Context) {
	skillItem, statusCode, err := convertHTTPBodyToSkill(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	c.JSON(statusCode, gin.H{"id": skills.AddSkill(skillItem)})
}

func RankSkillsHandler(c *gin.Context) {
	rankingItem, statusCode, err := convertHTTPBodyToRanking(c.Request.Body)
	if err != nil {
		c.JSON(statusCode, err)
		return
	}
	err = skills.RankSkill(rankingItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "")
}

func GetAllSkillsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, skills.GetAll())
}

func convertHTTPBodyToSkill(httpBody io.ReadCloser) (skills.Skill, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return skills.Skill{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToSkill(body)
}

func convertJSONBodyToSkill(jsonBody []byte) (skills.Skill, int, error) {
	var skillItem skills.Skill
	err := json.Unmarshal(jsonBody, &skillItem)
	if err != nil {
		return skills.Skill{}, http.StatusBadRequest, err
	}
	return skillItem, http.StatusOK, nil
}

func convertHTTPBodyToRanking(httpBody io.ReadCloser) (skills.Ranking, int, error) {
	body, err := ioutil.ReadAll(httpBody)
	if err != nil {
		return skills.Ranking{}, http.StatusInternalServerError, err
	}
	defer httpBody.Close()
	return convertJSONBodyToRanking(body)
}

func convertJSONBodyToRanking(jsonBody []byte) (skills.Ranking, int, error) {
	var rankingItem skills.Ranking
	err := json.Unmarshal(jsonBody, &rankingItem)
	if err != nil {
		return skills.Ranking{}, http.StatusBadRequest, err
	}
	return rankingItem, http.StatusOK, nil
}

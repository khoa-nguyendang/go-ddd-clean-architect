package controllers

import (
	mds "app/core/models"
	"app/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *RestfulServer) searchJob(c *gin.Context) {
	index := c.DefaultQuery("index", "")
	term := c.DefaultQuery("term", "")
	pageIndex := utils.GetInteger(c.DefaultQuery("pageIndex", "0"))
	pageAmount := utils.GetInteger(c.DefaultQuery("pageAmount", "20"))
	res, err := s.SearchJob(c, index, term, pageIndex, pageAmount)
	// jobs, err := json.Marshal(res)
	c.JSON(http.StatusOK, gin.H{
		"term":       term,
		"pageIndex":  pageIndex,
		"pageAmount": pageAmount,
		"data":       res,
		"err":        err,
	})
}

func (s *RestfulServer) searchJobDatabase(c *gin.Context) {
	term := c.DefaultQuery("term", "")
	pageIndex := utils.GetInteger(c.DefaultQuery("pageIndex", "0"))
	pageAmount := utils.GetInteger(c.DefaultQuery("pageAmount", "20"))
	res, err := s.SearchJobDatabase(c, term, pageIndex, pageAmount)
	// jobs, err := json.Marshal(res)
	c.JSON(http.StatusOK, gin.H{
		"term":       term,
		"pageIndex":  pageIndex,
		"pageAmount": pageAmount,
		"data":       res,
		"err":        err,
	})
}

func (s *RestfulServer) getJob(c *gin.Context) {
	jobId := c.Param("id")
	job, err := s.GetJob(c, jobId)
	c.JSON(http.StatusOK, gin.H{
		"jobId": jobId,
		"data":  job,
		"err":   err,
	})
}

func (s *RestfulServer) addJob(c *gin.Context) {
	var payload mds.Job

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"payload": payload,
			"err":     err,
		})
	}

	job, err := s.AddJob(c, payload)
	c.JSON(http.StatusOK, gin.H{
		"payload": payload,
		"job":     job,
		"err":     err,
	})
}

func (s *RestfulServer) addTestJob(c *gin.Context) {
	err := s.AddTestJob(c)
	c.JSON(http.StatusOK, gin.H{
		"payload": "",
		"job":     "",
		"err":     err,
	})
}

func (s *RestfulServer) updateJob(c *gin.Context) {
	var payload mds.Job

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"payload": payload,
			"err":     err,
		})
	}

	job, err := s.UpdateJob(c, payload)
	c.JSON(http.StatusOK, gin.H{
		"payload": payload,
		"job":     job,
		"err":     err,
	})
}

func (s *RestfulServer) deleteJob(c *gin.Context) {
	jobId := c.Param("id")
	job, err := s.DeleteJob(c, jobId)
	c.JSON(http.StatusOK, gin.H{
		"jobId": jobId,
		"data":  job,
		"err":   err,
	})
}

func (s *RestfulServer) patchJob(c *gin.Context) {
	var payload mds.Job

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"payload": payload,
			"err":     err,
		})
	}

	job, err := s.PatchJob(c, payload)
	c.JSON(http.StatusOK, gin.H{
		"payload": payload,
		"job":     job,
		"err":     err,
	})
}

func (s *RestfulServer) headJob(c *gin.Context) {
	jobId := c.Param("id")
	job, err := s.GetJob(c, jobId)
	c.JSON(http.StatusOK, gin.H{
		"jobId": jobId,
		"data":  job,
		"err":   err,
	})
}

func (s *RestfulServer) createIndex(c *gin.Context) {
	index := c.Param("index")
	err := s.CreateIndex(c, index)
	c.JSON(http.StatusOK, gin.H{
		"index": index,
		"err":   err,
	})
}

func (s *RestfulServer) pushDocuments(c *gin.Context) {
	index := c.Param("index")
	skips := utils.GetInteger(c.DefaultQuery("skips", "0"))
	takes := utils.GetInteger(c.DefaultQuery("takes", "10000"))
	sucess, fail, err := s.PushDocuments(c, index, takes, skips)

	c.JSON(http.StatusOK, gin.H{
		"index":  index,
		"sucess": sucess,
		"fail":   fail,
		"err":    err,
	})
}

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/yangpeng-chn/go-web-framework/middlewares"
	"github.com/yangpeng-chn/go-web-framework/models"
	"github.com/yangpeng-chn/go-web-framework/utils/formaterror"
	"github.com/yangpeng-chn/go-web-framework/utils/logger"
)

// CreatePost creats post
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	var err error
	var formattedError error
	var body []byte
	var code = http.StatusBadRequest
	var createdPost *models.Post
	var uid uint32
	post := models.Post{}

	if body, err = server.ParseRequest(w, r); err != nil {
		code = http.StatusUnprocessableEntity
		goto Error
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		middlewares.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post.Prepare()
	err = post.Validate()
	if err != nil {
		middlewares.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err = middlewares.ExtractTokenID(r)
	if err != nil {
		err = errors.New("Unauthorized")
		code = http.StatusUnauthorized
		goto Error
	}
	if uid != post.AuthorID {
		middlewares.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	createdPost, err = post.SavePost(server.DB)
	if err != nil {
		formattedError = formaterror.FormatError(err.Error())
		middlewares.ERROR(w, http.StatusInternalServerError, formattedError)
		logger.WriteLog(r, http.StatusInternalServerError, formattedError, server.GetCurrentFuncName())
		return
	}
	code = http.StatusCreated
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, createdPost.ID))
	middlewares.JSON(w, code, createdPost)
	logger.WriteLog(r, code, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, code, err)
	logger.WriteLog(r, code, err, server.GetCurrentFuncName())
}

// GetPosts gets posts
func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {
	var code = http.StatusBadRequest
	post := models.Post{}
	posts, err := post.FindAllPosts(server.DB)
	if err != nil {
		code = http.StatusInternalServerError
		goto Error
	}
	code = http.StatusOK
	middlewares.JSON(w, code, posts)
	logger.WriteLog(r, code, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, code, err)
	logger.WriteLog(r, code, err, server.GetCurrentFuncName())
}

// GetPost gets post
func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	var code = http.StatusBadRequest
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	post := models.Post{}
	var postReceived *models.Post
	if err != nil {
		goto Error
	}
	postReceived, err = post.FindPostByID(server.DB, pid)
	if err != nil {
		code = http.StatusInternalServerError
		goto Error
	}
	code = http.StatusOK
	middlewares.JSON(w, code, postReceived)
	logger.WriteLog(r, code, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, code, err)
	logger.WriteLog(r, code, err, server.GetCurrentFuncName())
}

// UpdatePost update post by id
func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var code = http.StatusBadRequest
	var uid uint32
	var body []byte
	vars := mux.Vars(r)
	post := models.Post{}
	postUpdate := models.Post{}
	var postUpdated *models.Post
	// Check if the post id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		goto Error
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err = middlewares.ExtractTokenID(r)
	if err != nil {
		err = errors.New("Unauthorized")
		code = http.StatusUnauthorized
		goto Error
	}

	// Check if the post exist
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		err = errors.New("Post not found")
		code = http.StatusNotFound
		goto Error
	}

	// If a user attempt to update a post not belonging to him
	if uid != post.AuthorID {
		err = errors.New("Unauthorized")
		code = http.StatusUnauthorized
		goto Error
	}

	if body, err = server.ParseRequest(w, r); err != nil {
		code = http.StatusUnprocessableEntity
		goto Error
	}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		code = http.StatusUnprocessableEntity
		goto Error
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != postUpdate.AuthorID {
		err = errors.New("Unauthorized")
		code = http.StatusUnauthorized
		goto Error
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		code = http.StatusUnprocessableEntity
		goto Error
	}

	postUpdate.ID = post.ID //this is important to tell the model the post id to update, the other update field are set above
	postUpdated, err = postUpdate.UpdateAPost(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		middlewares.ERROR(w, http.StatusInternalServerError, formattedError)
		logger.WriteLog(r, http.StatusInternalServerError, formattedError, server.GetCurrentFuncName())
		return
	}
	code = http.StatusOK
	middlewares.JSON(w, code, postUpdated)
	logger.WriteLog(r, code, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, code, err)
	logger.WriteLog(r, code, err, server.GetCurrentFuncName())
}

// DeletePost delete post by id
func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	var code = http.StatusBadRequest
	vars := mux.Vars(r)
	post := models.Post{}
	var uid uint32
	// Is a valid post id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		goto Error
	}

	// Is this user authenticated?
	uid, err = middlewares.ExtractTokenID(r)
	if err != nil {
		err = errors.New("Unauthorized")
		code = http.StatusUnauthorized
		goto Error
	}

	// Check if the post exist
	err = server.DB.Debug().Model(models.Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		middlewares.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != post.AuthorID {
		middlewares.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = post.DeleteAPost(server.DB, pid, uid)
	if err != nil {
		middlewares.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	code = http.StatusOK
	middlewares.JSON(w, code, middlewares.Result{Code: 200, Msg: "OK"})
	logger.WriteLog(r, code, nil, server.GetCurrentFuncName())
	return
Error:
	middlewares.ERROR(w, code, err)
	logger.WriteLog(r, code, err, server.GetCurrentFuncName())
}

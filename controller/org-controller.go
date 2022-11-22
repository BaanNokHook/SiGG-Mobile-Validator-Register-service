package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nextclan/user-service-go/model"
	"github.com/nextclan/user-service-go/service"
)

/*
 *	Org controller layer to accept request from exposed API and pass it org service layer
**/

var (
	orgSVC service.OrgService = service.NewOrgService()  
)  

func GetOrgById(w http.ResponseWriter, r *http.Request) {  
	params := mux.Vars(r)  
	orgId := params["id"]  
	var org *model.org  
	var err error  

	if org, err = orgSVC.GetOrgById(orgId); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())   
		return  
	}  
	RespondWithJSON(w, http.StatusOK, org)   

}  

func RespondWithJSON(w http.ResponseWriter, i int, org *invalid type) {  
	panic("unimplemented")   
}  


func RespondWithError(w http.ResponseWriter, i int, s string) {
	panic("unimplemented")
}



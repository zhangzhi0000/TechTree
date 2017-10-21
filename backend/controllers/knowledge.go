package controllers

import (
	"encoding/json"

	"github.com/Jsharkc/TechTree/backend/general"
	"github.com/Jsharkc/TechTree/backend/models"
	"github.com/Jsharkc/TechTree/backend/utils"
	"github.com/Jsharkc/TechTree/lib/log"
)

type KnowledgeController struct {
	BaseController
}

func (kc *KnowledgeController) GetKnowledge() {
	var (
		err    error
		k      []models.Knowledge
		flag   bool
		nid struct{
			NID string `json:"nid" valid:"Required"`
		}
	)

	err = json.Unmarshal(kc.Ctx.Input.RequestBody, &nid)
	if err != nil {
		log.Logger.Error("Get knowledge json unmarshal err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrInvalidParams}
		goto finish
	}

	flag, err = utils.GlobalValid.Valid(&nid)
	if !flag {
		for _, err := range utils.GlobalValid.Errors {
			log.Logger.Error("The knowledge key "+err.Key+" has err:", err)
		}

		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrInvalidParams}
		goto finish
	}

	k, err = models.KnowledgeService.GetKnowledgeByNode(nid.NID)
	if err != nil {
		log.Logger.Error("get knowledge err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrMysql}
		goto finish
	}

	kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrSucceed, general.RespKeyData: k}
finish:
	kc.ServeJSON(true)
}

func (kc *KnowledgeController) AdminAddKnowledge() {
	var (
		err  error
		k    models.Knowledge
		flag bool
	)

	err = json.Unmarshal(kc.Ctx.Input.RequestBody, &k)
	if err != nil {
		log.Logger.Error("AdminAddKnowledge json unmarshal err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrInvalidParams}
		goto finish
	}

	flag, err = utils.GlobalValid.Valid(&k)
	if !flag {
		for _, err := range utils.GlobalValid.Errors {
			log.Logger.Error("The knowledge key "+err.Key+" has err:", err)
		}

		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrInvalidParams}
		goto finish
	}

	err = models.KnowledgeService.AdminAddKnowledge(&k)
	if err != nil {
		log.Logger.Error("Add knowledge mysql err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrMysql}
		goto finish
	}

	kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrSucceed}
	log.Logger.Info("Add question success")
	finish:
	kc.ServeJSON(true)
}

func (kc *KnowledgeController) Delete() {
	var (
		err      error
		id       string
	)

	err = json.Unmarshal(kc.Ctx.Input.RequestBody, &id)
	if err != nil || id == "" {
		log.Logger.Error("Delete question json unmarshal err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrInvalidParams}
		goto finish
	}

	err = models.KnowledgeService.DeleteKnowledge(id)
	if err != nil {
		log.Logger.Error("Delete question mysql err:", err)
		kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrMysql}
		goto finish
	}

	kc.Data["json"] = map[string]interface{}{general.RespKeyStatus: general.ErrSucceed}
	log.Logger.Info("Delete question success")
	finish:
	kc.ServeJSON(true)
}
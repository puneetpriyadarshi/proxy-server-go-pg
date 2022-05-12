package controllers

import (
	"fmt"
	"log"
	"net/http"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type Tenant struct {
	ID   int
	NAME string
}

func CreateTenantTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Tenant{}, opts)
	if createError != nil {
		log.Printf("error while creating a table, Reason : %v\n", createError)
		return createError
	}
	log.Printf("Tenant table created")
	return nil
}

type TenantActivity struct {
	ID                int
	TENANT_ID         string
	CREATED_AT        time.Time
	API_END_POINT     string
	BYTES_TRANSFERRED int
	IS_ERROR          int
}

func CreateTenantActivityTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&TenantActivity{}, opts)
	if createError != nil {
		log.Printf("error while creating a table, Reason : %v\n", createError)
		return createError
	}
	log.Printf("tenant activity table created")
	return nil
}

var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func UpdateBytes(z int, error_count int) {
	user := &TenantActivity{
		ID:                2,
		BYTES_TRANSFERRED: z,
		IS_ERROR:          error_count,
	}
	_, err := dbConnect.Model(user).Set("bytes_transferred = bytes_transferred+?", z).Where("id = 2").Update()
	_, err1 := dbConnect.Model(user).Set("is_error = is_error+?", error_count).Where("id = 2").Update()
	if err != nil {
		fmt.Println("ERROR IS  :   ", err)
	}
	if err1 != nil {
		fmt.Println("ERROR IS  :   ", err)
	}
}
func GetAllTenants(c *gin.Context) {
	var tenants []Tenant
	err := dbConnect.Model(&tenants).Select()

	if err != nil {
		log.Printf("Error while getting all tenants, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Tenants",
		"data":    tenants,
	})
}
func GetSingleTenant(c *gin.Context) {
	s := c.Param("tenantId")
	tenantId, err1 := strconv.Atoi(s)
	if err1 != nil {
		log.Println("THE ERROR IS : =====>>>>>", err1)
	}
	tenant := &Tenant{ID: tenantId}
	err := dbConnect.Select(tenant)
	if err != nil {
		log.Printf("Error while getting a single tenant, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Tenant not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Tenant",
		"data":    tenant,
	})
}

func CreateTenant(c *gin.Context) {
	var tenant Tenant
	err := c.BindJSON(&tenant)
	if err != nil {
		log.Fatal(err)
	}
	NAME := tenant.NAME

	insertError := dbConnect.Insert(&Tenant{
		NAME: NAME,
	})
	if insertError != nil {
		log.Printf("Error while inserting new tenant into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Tenant created Successfully",
	})
}
func GetSingleTenantActivity(c *gin.Context) {
	s := c.Param("tenantactivityId")
	tenantactivityId, err1 := strconv.Atoi(s)
	if err1 != nil {
		log.Println("THE ERROR IS : =====>>>>>", err1)
	}
	tenantactivity := &TenantActivity{ID: tenantactivityId}
	err := dbConnect.Select(tenantactivity)
	if err != nil {
		log.Printf("Error while getting a single tenant activity, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Tenant Activity not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"bytes_transferred": tenantactivity.BYTES_TRANSFERRED,
		"error_count":       tenantactivity.IS_ERROR,
	})
}

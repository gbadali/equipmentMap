package handler

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gbadali/equipmentMap/db"
	"github.com/gbadali/equipmentMap/view/equipment"
	"github.com/labstack/echo/v4"
)

type EquipmentHandler struct {
	DB *sql.DB
	Q  *db.Queries
}

// NewEquipmentHandler creates a new EquipmentHandler with the given db and queries
func NewEquipmentHandler(db *sql.DB, q *db.Queries) *EquipmentHandler {
	return &EquipmentHandler{
		DB: db,
		Q:  q,
	}
}

// HandleShowEquipment handles the request to show all equipment in a table
func (e EquipmentHandler) HandleShowEquipment(c echo.Context) error {
	fmt.Println("Handling show all equipment request")
	equipmentList, err := e.Q.ListEquipmentAndParent((c.Request().Context()))
	if err != nil {
		return err
	}
	return render(c, equipment.EquipmentList(equipmentList))
}

// HandleListEquipment handles the request to show all equipment in a select and generates the options
func (e EquipmentHandler) HandleListEquipment(c echo.Context) error {
	fmt.Println("Handling list equipment request")
	equipmentList, err := e.Q.ListEquipment(c.Request().Context())
	fmt.Println(equipmentList)
	if err != nil {
		return err
	}
	return render(c, equipment.EquipmentSelectOptions(equipmentList))
}

// HandleAddEquipment handles the request to show the form to add a new equipment
func (e EquipmentHandler) HandleAddEquipment(c echo.Context) error {
	list, err := e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}
	fmt.Println("Handling add equipment request")
	return render(c, equipment.EquipmentForm(list))
}

// HandleSaveEquipment handles the post request to save a new equipment
func (e EquipmentHandler) HandleSaveEquipment(c echo.Context) error {
	name := c.FormValue("name")
	parent, err := strconv.ParseInt(c.FormValue("parent"), 10, 64)
	if err != nil {
		return err
	}
	equip := db.CreateEquipmentParams{
		Name: name,
		Parent: sql.NullInt64{
			Valid: true,
			Int64: parent,
		},
	}
	fmt.Printf("Handling save %v request", equip)

	list, err := e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}

	err = e.Q.CreateEquipment(c.Request().Context(), equip)
	if err != nil {
		return err
	}

	return render(c, equipment.EquipmentForm(list))
}

// HandleShowIndividualEquipment handles the request to show an individual equipment
// HandleShowIndividualEquipment handles the request to show individual equipment.
// It retrieves the equipment with the specified ID from the database and its parent and children.
// If the equipment has a parent, it also retrieves the parent equipment.
// Finally, it renders the equipment information using the render function.
// Parameters:
// - c: the echo.Context object representing the HTTP request and response.
// Returns:
// - error: an error if any occurred during the handling of the request, otherwise nil.
func (e EquipmentHandler) HandleShowIndividualEquipment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	fmt.Printf("Handling show individual equipment request for %v\n", id)

	breadcrumbs, err := e.Q.GetHierarchy(c.Request().Context(), id)
	if err != nil {
		return err
	}
	fmt.Println(breadcrumbs)

	// Create the equipment, parend and children variables
	var equip db.Equipment

	equip, err = e.Q.GetEquipment(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return render(c, equipment.Equipment(equip))
}

// HandleEditEquipment handles the request to show the form to edit an equipment
func (e EquipmentHandler) HandleEditEquipment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	fmt.Printf("Handling edit equipment request for %v\n", id)

	// Create the equipment, parend and children variables
	var equip db.Equipment
	var list []db.Equipment

	equip, err = e.Q.GetEquipment(c.Request().Context(), id)
	if err != nil {
		return err
	}

	list, err = e.Q.ListEquipment(c.Request().Context())
	if err != nil {
		return err
	}

	return render(c, equipment.EditEquipment(equip, list))

}

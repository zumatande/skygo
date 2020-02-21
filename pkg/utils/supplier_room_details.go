package utils

import (
	"fmt"

	"github.com/Zumata/v3-common/data"
)

// various supplier room details
var (
	srdDeluxe = data.SupplierRoomDetails{
		RatePlanCode: "drrbqb21",
		RoomType:     "Deluxe",
		RoomTypeCode: "drbqb21",
	}
	srdStandard = data.SupplierRoomDetails{
		RatePlanCode: "rbqb21",
		RoomType:     "Standard",
		RoomTypeCode: "rrbqb21",
	}

	commonSupplierRoomDetails = map[string]data.SupplierRoomDetails{
		"Deluxe":   srdDeluxe,
		"Standard": srdStandard,
	}
)

// NewSupplierRoomDetails ...
func NewSupplierRoomDetails(roomType string, bedCount int) data.SupplierRoomDetails {
	// TODO: parameterize type of bed
	desc := fmt.Sprintf("%v - %v Queen Bed", roomType, bedCount)

	if rp, ok := commonSupplierRoomDetails[roomType]; ok {
		rp.Description = desc
		return rp
	}

	return data.SupplierRoomDetails{
		Description:  desc,
		RatePlanCode: "ratePlan888",
		RoomType:     roomType,
		RoomTypeCode: "roomType888",
	}
}

// CommonSupplierRoomDetails ...
func CommonSupplierRoomDetails() map[string]data.SupplierRoomDetails {
	return commonSupplierRoomDetails
}

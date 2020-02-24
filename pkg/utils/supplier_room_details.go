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

// GetRoomSurcharge returns surcharge for common room types,
// if not in common suppliers, returns input surcharge
func GetRoomSurcharge(roomType string, surcharge ...float64) float64 {
	switch roomType {
	case "STANDARD":
		return 1.0
	case "DELUXE":
		return 1.5
	default:
		if surcharge == nil {
			return 1.0
		}
		return surcharge[0]
	}
}

// GetBedCount ...
func GetBedCount(req *data.SearchRequest) int {
	ceilBeds := req.Params.AdultCount + len(req.Params.Children)
	beds := ceilBeds / req.Params.RoomCount / 2
	if ceilBeds%2 != 0 {
		beds++
	}

	return beds
}

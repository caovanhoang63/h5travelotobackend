package hoteldetailmodel

import (
	"h5travelotobackend/common"
)

/*
	{
	  "_id" : "2131231231232123",
	  "hotel_id": 1,
	  "logo": {
		"id": 1,
		"url": "http://example.com/logo.jpg",
		"width": 100,
		"height": 100,
		"cloud_name": "s3",
		"extension": "jpg"
	  },
	  "covers": [],
	  "policies": {
		"check_in_time": "2018-01-01T12:00:00Z",
		"check_out_time": "2018-01-01T12:00:00Z",
		"required_documents": "ID Card",
		"deposit": 100,
		"min_age": 18
	  },
	  "facilities": [
		{
		  "name": "common",
		  "categories": {
			"wifi": true,
			"parking": true
		  }
		},
		{
		  "name": "general",
		  "categories": {
			"wifi": true,
			"parking": true
		  }
		}
	  ]
	}

*/

const EntityName = "HotelDetail"

type HotelDetail struct {
	common.MongoModel `bson:",inline" json:",inline"`
	HotelId           int            `json:"-" bson:"hotel_id"`
	Logo              *common.Image  `json:"logo,omitempty" bson:"logo,omitempty" `
	Covers            *common.Images `json:"covers,omitempty" bson:"covers,omitempty" `
	Policies          *Policies      `json:"policies,omitempty" bson:"policies,omitempty"`
	Facilities        *FacilityList  `json:"facilities,omitempty" bson:"facilities,omitempty"`
	Description       *Description   `bson:"description" json:"description"`
}

type Description struct {
	LocationDetails string `json:"location_details" bson:"location_details"`
	About           string `bson:"about" json:"about"`
}

func (HotelDetail) CollectionName() string {
	return "hotel_details"
}

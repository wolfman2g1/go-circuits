
import (
	"fmt"
  "database/sql"
  "github.com/jinzhu/gorm"
  )
type Carrier struct {
  ID int `sql:"AUTO_INCREMENT"`
  CarrierName string  `sql:"size:255"`
  SupportEmail string
  SupportNum string
  CarrierStreet string
  CarrierCity string
  CarrierState string
  CarrierPostal string
  CreatedAt time.Time
  UpdatedAt time.Time
  Circuits []Circuit // has many circuits

}
type Circuit struct {
  ID int `sql:"AUTO_INCREMENT"`
  CircuitID string `sql:"size:255"`
  ALoc string
  ZLoc string
  CarrierBlock string `sql:"size:255"`
  ExternalBlock string `sql:"size:255"`
  CircuitID // belongs to Carrier
  CreatedAt time.Time
  UpdatedAt time.Time
}

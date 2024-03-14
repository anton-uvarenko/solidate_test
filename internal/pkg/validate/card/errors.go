package card

import "github.com/anton-uvarenko/solidgate_test/internal/pkg/payload"

var ErrCardNotANumber = payload.NewError(001, "provided card number is not a number")
var ErrIssuingNetworkData = payload.NewError(002, "issuing network data is incorrect")
var ErrCheckSum = payload.NewError(003, "checksum is invalid")
var ErrInvalidDate = payload.NewError(004, "invalid card expiration date")

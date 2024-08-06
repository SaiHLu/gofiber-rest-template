package main

import (
	"github.com/SaiHLu/rest-template/common/logger"
	"github.com/google/uuid"
)

func main() {
	uId := "1d73a9d9-ceb7-420e-b82e-a64e7dc61ceb"

	logger.Debug(uuid.MustParse(uId).String())
}

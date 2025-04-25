package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"

	"github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/internal/infrastructure/db/postgres"
	pkgGorm "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/pkg/gorm/postgres"
	pkgZap "github.com/liwaisi-tech/iot-hub-center/backends/go-mqtt-subscriber/pkg/zap"
)

var logger = pkgZap.GetLogger()

const (
	HEADER = `
#                                     #######                     
#       # #    #   ##   #  ####  #       #    ######  ####  #    #
#       # #    #  #  #  # #      #       #    #      #    # #    #
#       # #    # #    # #  ####  #       #    #####  #      ######
#       # # ## # ###### #      # #       #    #      #      #    #
#       # ##  ## #    # # #    # #       #    #      #    # #    #
####### # #    # #    # #  ####  #       #    ######  ####  #    #
                                                                  
                           ###        #######                     
                            #   ####     #                        
                            #  #    #    #                        
                            #  #    #    #                        
                            #  #    #    #                        
                            #  #    #    #                        
                           ###  ####     #                        
                                                                  	`
)

func main() {
	fmt.Println(HEADER)
	logger.Infow("Initiating the Liwaisi IoT Hub MQTT Subscriber")
    err := runMigrations()
	if err != nil {
		logger.Errorw("Failed to run migrations", "error", err)
		return
	}
}

func runMigrations() error {
    logger.Info("Running migrations...")
    db, err := pkgGorm.GetPostgresDB()
	if err != nil {
		return err
	}
	return postgres.Migrate(db)
}

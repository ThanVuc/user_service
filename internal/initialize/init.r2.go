package initialize

import (
	"user_service/global"

	"github.com/thanvuc/go-core-lib/storage"
)

func initR2() {
	// R2 initialization logic goes here
	r2Congigs := global.Config.R2
	r2Clients, err := storage.NewClient(storage.Config{
		AccountID: r2Congigs.AccountID,
		Endpoint:  r2Congigs.Endpoint,
		AccessKey: r2Congigs.AccessKeyID,
		SecretKey: r2Congigs.SecrecAccessKey,
		Bucket:    r2Congigs.BucketName,
		UseSSL:    r2Congigs.UseSSL,
		PublicURL: r2Congigs.PublicURL,
	})
	if err != nil {
		global.Logger.Error("Failed to initialize R2 client", "")
		panic(err)
	}

	global.R2Client = r2Clients
}

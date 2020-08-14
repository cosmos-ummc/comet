package cmd

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/handlers"
	"comet/pkg/logger"
	model2 "comet/pkg/model"
	"comet/pkg/protocol/grpc"
	"comet/pkg/protocol/rest"
	"comet/pkg/utility"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"time"
)

// Config is configuration for Server
type Config struct {
	GRPCPort      string
	HTTPPort      string
	LogLevel      int
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	cfg := &Config{GRPCPort: "10001", HTTPPort: "10002", LogLevel: -1, LogTimeFormat: "02 Jan 2006 15:04:05 MST"}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// Load .env configuration
	err := godotenv.Load()
	if err != nil {
		logger.Log.Warn(".env file not found, using environment variables")
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error getting connect mongo client: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	// initialize model
	model := model2.InitModel(mongoClient)
	_, u, _ := model.QueryUsers(ctx, nil, nil, nil, false)
	for _, uu := range u {
		uu.BlockList = []string{}
		uu.Visible = false
		model.UpdateUser(ctx, uu)
	}

	_, c, _ := model.QueryChatRooms(ctx, nil, nil, nil)
	for _, cc := range c {
		model.DeleteChatRoom(ctx, cc.ID)
		//cc.Blocked = false
		//model.UpdateChatRoom(ctx, cc)
	}

	_, d, _ := model.QueryChatMessages(ctx, nil, nil, nil)
	for _, dd := range d {
		model.DeleteChatMessage(ctx, dd.ID)
	}

	_, dec, _ := model.QueryDeclarations(ctx, nil, nil, nil)
	for _, dd := range dec {
		model.DeleteDeclaration(ctx, dd.ID)
		//if dd.Category == "" {
		//	model.DeleteDeclaration(ctx, dd.ID)
		//}
	}

	_, ms, _ := model.QueryMeetings(ctx, nil, nil, nil)
	for _, m := range ms {
		p, _ := model.GetPatient(ctx, m.PatientID)
		m.PatientPhoneNumber = p.PhoneNumber
		m.PatientName = p.Name

		c, _ := model.GetConsultant(ctx, m.ConsultantID)
		m.ConsultantPhoneNumber = c.PhoneNumber
		m.ConsultantName = c.Name

		model.UpdateMeeting(ctx, m)
	}

	// telegram de-linking
	_, ppp, err := model.QueryPatients(ctx, nil, nil, nil)
	for _, p := range ppp {
		p.TelegramID = ""
		p.TutorialStage = 0
		model.UpdatePatient(ctx, p)
	}

	// report generator
	_, patients, err := model.QueryPatients(ctx, nil, nil, nil)
	for _, patient := range patients {

		var results []*dto.Question
		_, results, err := model.QueryQuestions(ctx, nil, nil, map[string]interface{}{constants.Category: constants.DASS})
		if err != nil {
			continue
		}

		// first declaration
		for _, r := range results {
			r.Score = int64(rand.Intn(2)) + 2
		}
		declaration := &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           results,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		}
		model.ClientCreateDeclaration(ctx, declaration)

		// second declaration
		for _, r := range results {
			r.Score = int64(rand.Intn(2))
		}
		declaration = &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           results,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())) + 100000,
		}
		model.ClientCreateDeclaration(ctx, declaration)

		_, resultsIesr, err := model.QueryQuestions(ctx, nil, nil, map[string]interface{}{constants.Category: constants.IESR})
		if err != nil {
			continue
		}

		// first declaration
		for _, r := range resultsIesr {
			r.Score = int64(rand.Intn(2)) + 3
		}
		declaration = &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           resultsIesr,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		}
		model.ClientCreateDeclaration(ctx, declaration)

		// second declaration
		for _, r := range resultsIesr {
			r.Score = int64(rand.Intn(2)) + 1
		}
		declaration = &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           resultsIesr,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())) + 100000,
		}
		model.ClientCreateDeclaration(ctx, declaration)

		_, resultsDaily, err := model.QueryQuestions(ctx, nil, nil, map[string]interface{}{constants.Category: constants.Daily})
		if err != nil {
			continue
		}

		// first declaration
		for _, r := range resultsDaily {
			r.Score = int64(rand.Intn(2)) + 1
		}
		declaration = &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           resultsDaily,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		}
		model.ClientCreateDeclaration(ctx, declaration)

		// second declaration
		for _, r := range resultsDaily {
			r.Score = int64(rand.Intn(2))
		}
		declaration = &dto.Declaration{
			ID:               uuid.NewV4().String(),
			PatientID:        patient.ID,
			Result:           resultsDaily,
			SubmittedAt:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())) + 100000,
		}
		model.ClientCreateDeclaration(ctx, declaration)
	}

	// initialize scheduler
	//go func() {
	//	it := utility.Scheduler{Enabled: true, Job: model.GenerateReport, RevokeJob: model.RevokeUserTokens}
	//	it.Start()
	//}()

	// initialize handlers
	handler := handlers.NewHandlers(model)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, handler, cfg.GRPCPort)
}

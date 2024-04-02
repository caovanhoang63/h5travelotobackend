package subcriber

import (
	"context"
	"fmt"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/asyncjob"
	"h5travelotobackend/component/pubsub"
	"log"
)

type consumerJob struct {
	Title   string
	Handler func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appContext.AppContext
}

func NewEngine(appCtx appContext.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	if err := engine.startSubTopic(common.TopicCreateNewRoom, true,
		IncreaseTotalRoomWhenCreateNewRoom(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicDeleteRoom, true,
		DecreaseTotalRoomWhenCreateNewRoom(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}
	if err := engine.startSubTopic(common.TopicCreateBooking, true,
		CreateBookingTracking(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicDeleteBooking, true,
		DeleteTrackingWhenBookingDeleted(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicCreateHotel, true,
		CreateOwnerWorker(engine.appCtx, context.Background()),
		CreateHotelFacilityDetails(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicConfirmBookingWhenSelectEnoughRoom, true,
		ConfirmBookingTracking(engine.appCtx, context.Background()),
	); err != nil {
		log.Println("Err:", err)
	}

	if err := engine.startSubTopic(common.TopicCreateRoomType, true,
		CreateRoomFacilityDetails(engine.appCtx, context.Background())); err != nil {
		log.Println("Err:", err)
	}

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic string, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)
	for _, item := range consumerJobs {
		log.Printf("Set up consumer for: %s", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Printf("running job for %s for. Value %s \n", job.Title, message.Data)
			return job.Handler(ctx, message)
		}
	}

	go func() {
		common.AppRecover()
		for {
			msg := <-c
			fmt.Println("Message: ", string(msg.Data))
			jobHdlArr := make([]asyncjob.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl)
			}
			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println("Err:", err)
			}
		}
	}()

	return nil
}

package queue

import (
	"context"
	"log"

	"github.com/malillog/testGolang/queue"
)

// New takes a processor and returns both
// a client and a worker. The client allows
// pushing jobs to the queue (with CreateJob)
// and the worker can run those jobs using
// the given Processor
func New(p Processor) (Client, Worker) {
	// return Client and Worker
	//p.Process()
	cli := ClientImpl{}
	wor := WorkerImpl{}

	return cli, wor
}

type JobImpl struct {
	ID    string
	Data  MarshalUnmarshaler
	State State
	Error string
}

type JobProcessingAccessImpl struct {
	JobImpl
}

func (jp JobProcessingAccessImpl) SetData(ctx context.Context, data MarshalUnmarshaler) error {
	jp.JobImpl.Data = data
	//jp.JobImpl.Error = ""
	//jp.State=""
	//jp.ID=""
	return nil
}

type ClientImpl struct {
}

func (cli ClientImpl) CreateJob(ctx context.Context, id string, initialData MarshalUnmarshaler) error {
	log.Print(queue.Queued)

	var job = new(JobImpl)
	job.ID = id
	job.Data = initialData
	job.State = "queued"
	job.Error = ""

	//ctx.Value(job.ID)
	//ctx.Value().WithValue(ctx, id, "secret")
	//var  State = queue.Queued
	//State est = queue.Queued
	//j := JobImpl{id, initialData, est, ""}	//ctx.Value()
	// agregar a queue
	//pushing jobs to the queue

	return nil
}

func (cli ClientImpl) GetJob(ctx context.Context, id string) (Job, error) {
	//return job para procesar
	log.Print(queue.Failed)

	//State est = queue.Queued
	//j := JobImpl{id, nil, est, ""}
	return nil, nil
}

type WorkerImpl struct {
}

func (wor WorkerImpl) Run(ctx context.Context, workers int) error {
	for index := 0; index < workers; index++ {
		log.Print("Inicio Run")
		log.Print("Fin Run")
	}

	return nil
}

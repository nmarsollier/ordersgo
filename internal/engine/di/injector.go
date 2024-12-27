package di

import (
	"github.com/nmarsollier/ordersgo/internal/engine/db"
	"github.com/nmarsollier/ordersgo/internal/engine/env"
	"github.com/nmarsollier/ordersgo/internal/engine/httpx"
	"github.com/nmarsollier/ordersgo/internal/engine/log"
	"github.com/nmarsollier/ordersgo/internal/events"
	"github.com/nmarsollier/ordersgo/internal/projections"
	"github.com/nmarsollier/ordersgo/internal/projections/order"
	"github.com/nmarsollier/ordersgo/internal/projections/status"
	"github.com/nmarsollier/ordersgo/internal/rabbit/consume"
	"github.com/nmarsollier/ordersgo/internal/rabbit/emit"
	"github.com/nmarsollier/ordersgo/internal/security"
	"github.com/nmarsollier/ordersgo/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// Singletons
var database *mongo.Database
var httpClient httpx.HTTPClient
var eventsCollection db.Collection
var ordersCollection db.Collection
var statusCollection db.Collection
var articleExistConsumer consume.ArticleExistConsumer
var logoutConsumer consume.LogoutConsumer
var orderPlacedConsumer consume.OrderPlacedConsumer

type Injector interface {
	Logger() log.LogRusEntry
	Database() *mongo.Database
	HttpClient() httpx.HTTPClient
	SecurityRepository() security.SecurityRepository
	SecurityService() security.SecurityService
	EventsCollection() db.Collection
	EventsRepository() events.EventsRepository
	EventService() events.EventService
	OrdersCollection() db.Collection
	OrdersRepository() order.OrderRepository
	OrderService() order.OrderService
	StatusCollection() db.Collection
	StatusRepository() status.StatusRepository
	StatusService() status.StatusService
	ProjectionsService() projections.ProjectionsService
	Service() services.Service
	ArticleExistConsumer() consume.ArticleExistConsumer
	LogoutConsumer() consume.LogoutConsumer
	OrderPlacedConsumer() consume.OrderPlacedConsumer
	RabbitChannel() emit.RabbitChannel
	RabbitEmit() emit.RabbitEmit
}

type Deps struct {
	CurrLog           log.LogRusEntry
	CurrHttpClient    httpx.HTTPClient
	CurrDatabase      *mongo.Database
	CurrSecRepo       security.SecurityRepository
	CurrSecSvc        security.SecurityService
	CurrEvtColl       db.Collection
	CurrOrdColl       db.Collection
	CurrStaColl       db.Collection
	CurrEvtRepo       events.EventsRepository
	CurrOrdRepo       order.OrderRepository
	CurrOrdSvc        order.OrderService
	CurrEvtSvc        events.EventService
	CurrStsRepo       status.StatusRepository
	CurrStsSvc        status.StatusService
	CurrPrjSvc        projections.ProjectionsService
	CurrSvc           services.Service
	CurrArtCon        consume.ArticleExistConsumer
	CurrLogCon        consume.LogoutConsumer
	CurrOrdPlaCon     consume.OrderPlacedConsumer
	CurrRabbitChannel emit.RabbitChannel
	CurrRabitEmit     emit.RabbitEmit
}

func NewInjector(log log.LogRusEntry) Injector {
	return &Deps{
		CurrLog: log,
	}
}

func (i *Deps) Logger() log.LogRusEntry {
	return i.CurrLog
}

func (i *Deps) Database() *mongo.Database {
	if i.CurrDatabase != nil {
		return i.CurrDatabase
	}

	if database != nil {
		return database
	}

	database, err := db.NewDatabase(env.Get().MongoURL, "catalog")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}

	return database
}

func (i *Deps) HttpClient() httpx.HTTPClient {
	if i.CurrHttpClient != nil {
		return i.CurrHttpClient
	}

	if httpClient != nil {
		return httpClient
	}

	httpClient = httpx.Get()
	return httpClient
}

func (i *Deps) SecurityRepository() security.SecurityRepository {
	if i.CurrSecRepo != nil {
		return i.CurrSecRepo
	}
	i.CurrSecRepo = security.NewSecurityRepository(i.Logger(), i.HttpClient())
	return i.CurrSecRepo
}

func (i *Deps) SecurityService() security.SecurityService {
	if i.CurrSecSvc != nil {
		return i.CurrSecSvc
	}
	i.CurrSecSvc = security.NewSecurityService(i.Logger(), i.SecurityRepository())
	return i.CurrSecSvc
}

func (i *Deps) EventsCollection() db.Collection {
	if i.CurrEvtColl != nil {
		return i.CurrEvtColl
	}

	if eventsCollection != nil {
		return eventsCollection
	}

	cartCollection, err := db.NewCollection(i.CurrLog, i.Database(), "events", "orderId")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}
	return cartCollection
}

func (i *Deps) EventsRepository() events.EventsRepository {
	if i.CurrEvtRepo != nil {
		return i.CurrEvtRepo
	}
	i.CurrEvtRepo = events.NewEventsRepository(i.Logger(), i.EventsCollection())
	return i.CurrEvtRepo
}

func (i *Deps) EventService() events.EventService {
	if i.CurrEvtSvc != nil {
		return i.CurrEvtSvc
	}
	i.CurrEvtSvc = events.NewEventService(i.Logger(), i.EventsRepository())
	return i.CurrEvtSvc
}

func (i *Deps) OrdersCollection() db.Collection {
	if i.CurrOrdColl != nil {
		return i.CurrOrdColl
	}

	if ordersCollection != nil {
		return ordersCollection
	}

	cartCollection, err := db.NewCollection(i.CurrLog, i.Database(), "order_projection", "orderId")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}
	return cartCollection
}

func (i *Deps) StatusCollection() db.Collection {
	if i.CurrStaColl != nil {
		return i.CurrStaColl
	}

	if statusCollection != nil {
		return statusCollection
	}

	cartCollection, err := db.NewCollection(i.CurrLog, i.Database(), "status_projection", "orderId")
	if err != nil {
		i.CurrLog.Fatal(err)
		return nil
	}
	return cartCollection
}

func (i *Deps) OrdersRepository() order.OrderRepository {
	if i.CurrOrdRepo != nil {
		return i.CurrOrdRepo
	}
	i.CurrOrdRepo = order.NewOrderRepository(i.Logger(), i.OrdersCollection())
	return i.CurrOrdRepo
}

func (i *Deps) StatusRepository() status.StatusRepository {
	if i.CurrStsRepo != nil {
		return i.CurrStsRepo
	}
	i.CurrStsRepo = status.NewStatusRepository(i.Logger(), i.StatusCollection())
	return i.CurrStsRepo
}

func (i *Deps) OrderService() order.OrderService {
	if i.CurrOrdSvc != nil {
		return i.CurrOrdSvc
	}
	i.CurrOrdSvc = order.NewOrderService(i.Logger(), i.OrdersRepository())
	return i.CurrOrdSvc
}

func (i *Deps) StatusService() status.StatusService {
	if i.CurrStsSvc != nil {
		return i.CurrStsSvc
	}
	i.CurrStsSvc = status.NewStatusService(i.Logger(), i.StatusRepository())
	return i.CurrStsSvc
}

func (i *Deps) ProjectionsService() projections.ProjectionsService {
	if i.CurrPrjSvc != nil {
		return i.CurrPrjSvc
	}
	i.CurrPrjSvc = projections.NewProjectionsService(i.Logger(), i.EventService(), i.OrderService(), i.StatusService())
	return i.CurrPrjSvc
}

func (i *Deps) Service() services.Service {
	if i.CurrSvc != nil {
		return i.CurrSvc
	}
	i.CurrSvc = services.NewService(i.Logger(), i.EventService(), i.ProjectionsService(), i.RabbitEmit())
	return i.CurrSvc
}

func (i *Deps) ArticleExistConsumer() consume.ArticleExistConsumer {
	if i.CurrArtCon != nil {
		return i.CurrArtCon
	}

	if articleExistConsumer != nil {
		return articleExistConsumer
	}

	articleExistConsumer = consume.NewArticleExistConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.Service())
	return articleExistConsumer
}

func (i *Deps) LogoutConsumer() consume.LogoutConsumer {
	if i.CurrLogCon != nil {
		return i.CurrLogCon
	}

	if logoutConsumer != nil {
		return logoutConsumer
	}

	logoutConsumer = consume.NewLogoutConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.SecurityService())
	return logoutConsumer
}

func (i *Deps) OrderPlacedConsumer() consume.OrderPlacedConsumer {
	if i.CurrOrdPlaCon != nil {
		return i.CurrOrdPlaCon
	}

	if orderPlacedConsumer != nil {
		return orderPlacedConsumer
	}

	orderPlacedConsumer = consume.NewOrderPlacedConsumer(env.Get().FluentUrl, env.Get().RabbitURL, i.Service())
	return orderPlacedConsumer
}

func (i *Deps) RabbitChannel() emit.RabbitChannel {
	if i.RabbitChannel != nil {
		return i.CurrRabbitChannel
	}
	chn, err := emit.NewChannel(env.Get().RabbitURL, i.Logger())
	if err != nil {
		i.Logger().Fatal(err)
		return nil
	}
	i.CurrRabbitChannel = chn
	return i.CurrRabbitChannel
}

func (i *Deps) RabbitEmit() emit.RabbitEmit {
	if i.CurrRabitEmit != nil {
		return i.CurrRabitEmit
	}
	i.CurrRabitEmit = emit.NewRabbitEmit(i.Logger(), i.RabbitChannel())
	return i.CurrRabitEmit
}

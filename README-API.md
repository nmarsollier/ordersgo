# OrdersGo
Microservicio de Ordenes.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article-data

#### GET
##### Summary

Mensage Rabbit order/article-data

##### Description

Cuando se consume place-order se genera la orden y se inicia el proceso.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| place-order | body | Message para Type = place-order | Yes | [consume.consumePlaceDataMessage](#consumeconsumeplacedatamessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/cart/article-data

#### PUT
##### Summary

Emite Validar Artículos a Cart cart/article-data

##### Description

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Mensage de validacion | Yes | [emit.SendValidationMessage](#emitsendvalidationmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [consume.logoutMessage](#consumelogoutmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

#### PUT
##### Summary

Mensage Rabbit

##### Description

SendOrderPlaced envía un broadcast a rabbit con logout. Esto no es Rest es RabbitMQ.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Order Placed Event | Yes | [emit.message](#emitmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### /v1/orders

#### GET
##### Summary

Ordenes de Usuario

##### Description

Busca todas las ordenes del usuario logueado.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ordenes | [ [rest.OrderListData](#restorderlistdata) ] |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/orders/:orderId

#### GET
##### Summary

Buscar Orden

##### Description

Busca una order del usuario logueado, por su id.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| orderId | path | ID de orden | Yes | string |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ordenes | [order_projection.Order](#order_projectionorder) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/orders/:orderId/payment

#### POST
##### Summary

Agrega un Pago

##### Description

Agrega un Pago

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| orderId | path | ID de orden | Yes | string |
| Authorization | header | bearer {token} | Yes | string |
| body | body | Informacion del pago | Yes | [events.PaymentEvent](#eventspaymentevent) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Ordenes | [order_projection.Order](#order_projectionorder) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [server.ErrorData](#servererrordata) |
| 404 | Not Found | [server.ErrorData](#servererrordata) |
| 500 | Internal Server Error | [server.ErrorData](#servererrordata) |

### /v1/orders/:orderId/update

#### GET
##### Summary

Actualiza la proyeccion

##### Description

Actualiza las proyecciones en caso que hayamos roto algo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |
| orderId | path | ID de orden | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | No Content |

---
### Models

#### consume.consumeArticleDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.ValidationEvent](#eventsvalidationevent) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### consume.consumePlaceDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.PlacedOrderData](#eventsplacedorderdata) |  | No |
| queue | string |  | No |
| type | string | *Example:* `"place-order"` | No |
| version | integer |  | No |

#### consume.logoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string | *Example:* `"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlbklEIjoiNjZiNjBlYzhlMGYzYzY4OTUzMzJlOWNmIiwidXNlcklEIjoiNjZhZmQ3ZWU4YTBhYjRjZjQ0YTQ3NDcyIn0.who7upBctOpmlVmTvOgH1qFKOHKXmuQCkEjMV3qeySg"` | No |
| type | string | *Example:* `"logout"` | No |

#### emit.ArticleValidationData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| referenceId | string |  | No |

#### emit.SendValidationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [emit.ArticleValidationData](#emitarticlevalidationdata) |  | No |
| queue | string |  | No |
| type | string |  | No |

#### emit.articlePlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| quantity | integer |  | No |

#### emit.message

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [emit.orderPlacedData](#emitorderplaceddata) |  | No |
| queue | string |  | No |
| type | string |  | No |

#### emit.orderPlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [emit.articlePlacedData](#emitarticleplaceddata) ] |  | No |
| cartId | string |  | No |
| orderId | string |  | No |

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### events.PaymentEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | Yes |
| method | [events.PaymentMethod](#eventspaymentmethod) |  | Yes |
| orderId | string |  | Yes |

#### events.PaymentMethod

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| events.PaymentMethod | string |  |  |

#### events.PlacePrderArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| id | string |  | Yes |
| quantity | integer |  | Yes |

#### events.PlacedOrderData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [events.PlacePrderArticleData](#eventsplaceprderarticledata) ] |  | Yes |
| cartId | string |  | Yes |
| userId | string |  | Yes |

#### events.ValidationEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| price | number |  | No |
| referenceId | string |  | No |
| stock | integer |  | No |
| valid | boolean |  | No |

#### order_projection.Article

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | Yes |
| isValid | boolean |  | No |
| isValidated | boolean |  | No |
| quantity | integer |  | Yes |
| unitaryPrice | number |  | No |

#### order_projection.Order

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [order_projection.Article](#order_projectionarticle) ] |  | No |
| cartId | string |  | Yes |
| created | string |  | No |
| id | string |  | No |
| orderId | string |  | Yes |
| payments | [ [order_projection.PaymentEvent](#order_projectionpaymentevent) ] |  | No |
| status | [order_projection.OrderStatus](#order_projectionorderstatus) |  | Yes |
| updated | string |  | No |
| userId | string |  | Yes |

#### order_projection.OrderStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| order_projection.OrderStatus | string |  |  |

#### order_projection.PaymentEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | No |
| method | [events.PaymentMethod](#eventspaymentmethod) |  | No |

#### rest.OrderListData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | integer |  | No |
| cartId | string |  | No |
| created | string |  | No |
| id | string |  | No |
| status | [order_projection.OrderStatus](#order_projectionorderstatus) |  | No |
| totalPayment | number |  | No |
| totalPrice | number |  | No |
| updated | string |  | No |

#### server.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

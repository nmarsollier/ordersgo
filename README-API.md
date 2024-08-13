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

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| place-order | body | Message para Type = place-order | Yes | [r_consume.ConsumePlaceDataMessage](#r_consumeconsumeplacedatamessage) |

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
| body | body | Mensage de validacion | Yes | [r_emit.SendValidationMessage](#r_emitsendvalidationmessage) |

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
| body | body | Estructura general del mensage | Yes | [r_consume.LogoutMessage](#r_consumelogoutmessage) |

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
| body | body | Order Placed Event | Yes | [r_emit.message](#r_emitmessage) |

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
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 200 | Ordenes | [order_proj.Order](#order_projorder) |
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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
| 200 | Ordenes | [order_proj.Order](#order_projorder) |
| 400 | Bad Request | [apperr.ValidationErr](#apperrvalidationerr) |
| 401 | Unauthorized | [engine.ErrorData](#engineerrordata) |
| 404 | Not Found | [engine.ErrorData](#engineerrordata) |
| 500 | Internal Server Error | [engine.ErrorData](#engineerrordata) |

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

#### apperr.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [apperr.errField](#apperrerrfield) ] |  | No |

#### apperr.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### engine.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

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

#### order_proj.Article

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | Yes |
| isValid | boolean |  | No |
| isValidated | boolean |  | No |
| quantity | integer |  | Yes |
| unitaryPrice | number |  | No |

#### order_proj.Order

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [order_proj.Article](#order_projarticle) ] |  | No |
| cartId | string |  | Yes |
| created | string |  | No |
| id | string |  | No |
| orderId | string |  | Yes |
| payments | [ [order_proj.PaymentEvent](#order_projpaymentevent) ] |  | No |
| status | [order_proj.OrderStatus](#order_projorderstatus) |  | Yes |
| updated | string |  | No |
| userId | string |  | Yes |

#### order_proj.OrderStatus

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| order_proj.OrderStatus | string |  |  |

#### order_proj.PaymentEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | No |
| method | [events.PaymentMethod](#eventspaymentmethod) |  | No |

#### r_consume.ConsumeArticleDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.ValidationEvent](#eventsvalidationevent) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### r_consume.ConsumePlaceDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.PlacedOrderData](#eventsplacedorderdata) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### r_consume.LogoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| type | string |  | No |

#### r_emit.ArticleValidationData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| referenceId | string |  | No |

#### r_emit.SendValidationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [r_emit.ArticleValidationData](#r_emitarticlevalidationdata) |  | No |
| queue | string |  | No |
| type | string |  | No |

#### r_emit.articlePlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| quantity | integer |  | No |

#### r_emit.message

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [r_emit.orderPlacedData](#r_emitorderplaceddata) |  | No |
| queue | string |  | No |
| type | string |  | No |

#### r_emit.orderPlacedData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [r_emit.articlePlacedData](#r_emitarticleplaceddata) ] |  | No |
| cartId | string |  | No |
| orderId | string |  | No |

#### rest.OrderListData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | integer |  | No |
| cartId | string |  | No |
| created | string |  | No |
| id | string |  | No |
| status | [order_proj.OrderStatus](#order_projorderstatus) |  | No |
| totalPayment | number |  | No |
| totalPrice | number |  | No |
| updated | string |  | No |

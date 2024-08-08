# OrdersGo
Microservicio de Ordenes.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /rabbit/article-data

#### PUT
##### Summary

Mensage Rabbit order/article-data

##### Description

Antes de iniciar las operaciones se validan los artículos contra el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| place-order | body | Message para Type = place-order | Yes | [rabbit.ConsumePlaceDataMessage](#rabbitconsumeplacedatamessage) |

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
| body | body | Mensage de validacion | Yes | [rabbit.SendValidationMessage](#rabbitsendvalidationmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### PUT
##### Summary

Mensage Rabbit

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [rabbit.LogoutMessage](#rabbitlogoutmessage) |

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
| 200 | Ordenes | [ [order_proj.Order](#order_projorder) ] |
| 400 | Bad Request | [errors.ErrValidation](#errorserrvalidation) |
| 401 | Unauthorized | [errors.ErrCustom](#errorserrcustom) |
| 404 | Not Found | [errors.ErrCustom](#errorserrcustom) |
| 500 | Internal Server Error | [errors.ErrCustom](#errorserrcustom) |

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
| 400 | Bad Request | [errors.ErrValidation](#errorserrvalidation) |
| 401 | Unauthorized | [errors.ErrCustom](#errorserrcustom) |
| 404 | Not Found | [errors.ErrCustom](#errorserrcustom) |
| 500 | Internal Server Error | [errors.ErrCustom](#errorserrcustom) |

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
| 400 | Bad Request | [errors.ErrValidation](#errorserrvalidation) |
| 401 | Unauthorized | [errors.ErrCustom](#errorserrcustom) |
| 404 | Not Found | [errors.ErrCustom](#errorserrcustom) |
| 500 | Internal Server Error | [errors.ErrCustom](#errorserrcustom) |

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

### /v1/orders_batch/payment_defined

#### GET
##### Summary

Batch Payment Defined

##### Description

Ejecuta un proceso batch que chequea ordenes en estado PAYMENT_DEFINED.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | No Content |

### /v1/orders_batch/placed

#### GET
##### Summary

Batch Placed

##### Description

Ejecuta un proceso batch para ordenes en estado PLACED.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | No Content |

### /v1/orders_batch/validated

#### GET
##### Summary

Batch Validated

##### Description

Ejecuta un proceso batch para ordenes en estado VALIDATED.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | bearer {token} | Yes | string |

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | No Content |

---
### Models

#### errors.ErrCustom

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |

#### errors.ErrField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### errors.ErrValidation

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errors.ErrField](#errorserrfield) ] |  | No |

#### events.PaymentEvent

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| amount | number |  | No |
| cartId | string |  | No |
| method | [events.PaymentMethod](#eventspaymentmethod) |  | No |

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
| isValid | boolean |  | No |
| price | number |  | No |
| stock | integer |  | No |

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

#### rabbit.ArticleValidationData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| referenceId | string |  | No |

#### rabbit.ConsumeArticleDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.ValidationEvent](#eventsvalidationevent) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### rabbit.ConsumePlaceDataMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [events.PlacedOrderData](#eventsplacedorderdata) |  | No |
| queue | string |  | No |
| type | string |  | No |
| version | integer |  | No |

#### rabbit.LogoutMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| type | string |  | No |

#### rabbit.SendValidationMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| exchange | string |  | No |
| message | [rabbit.ArticleValidationData](#rabbitarticlevalidationdata) |  | No |
| queue | string |  | No |
| type | string |  | No |

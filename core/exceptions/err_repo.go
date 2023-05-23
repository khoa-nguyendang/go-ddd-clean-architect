package exceptions

import "errors"

/**
 * errors to dispatch/repository
 */

var (
	// ErrNilJobImport ...
	ErrNilJobImport = errors.New("impoort 0 jobs")

	// ErrQueryData describe query data from database failed
	ErrQueryData = errors.New("query data by constrain failed")

	// ErrGetItemFromDB to tell caller, get item failure
	ErrGetItemFromDB = errors.New("get item from db failure")

	// ErrItemNotExistsInDB to tell there has no item by given constrain
	ErrItemNotExistsInDB = errors.New("item not exists in db")

	// ErrItemUnmarshalMap to describe get item but unmarshal failure
	ErrItemUnmarshalMap = errors.New("unmarshal item to struct failure")

	// ErrItemMarshalMap to describe struct marshal to db item  failure
	ErrItemMarshalMap = errors.New("marshal struct to db item failure")

	// ErrItemMarshalToDBFailed -
	ErrItemMarshalToDBFailed = errors.New("marshal struct to db item failed")

	// ErrPutItem to describe put item into db failure
	ErrPutItem = errors.New("put item into db failure")

	// ErrUpdateItem to describe update item from db failure
	ErrUpdateItem = errors.New("update item from db failure")

	// ErrItemNotExists to tell item is not exists in db
	ErrItemNotExists = errors.New("item not exists")

	// ErrItemIsEmpty ...
	ErrItemIsEmpty = errors.New("get empty item")

	// ErrGenUniqShipmentCode happens when reached max retry count for getting unique shipment_code
	ErrGenUniqShipmentCode = errors.New("error getting uniq shipment_code max retry reached")

	// ErrUnmarshalPagination describe unmarshal unmarshal db last evaluated key failed
	ErrUnmarshalPagination = errors.New("unmarshal dynamo pagination struct failed")

	// ErrUnmarshalMapList describe unmarshal map list from db has error
	ErrUnmarshalMapList = errors.New("unmarshal db map list failed")

	// ErrUnmarshalMap describe unmarshal map from db has error
	ErrUnmarshalMap = errors.New("unmarshal db map failed")

	// ErrNotImpl function not implement yet
	ErrNotImpl = errors.New("this function not implement yet")

	ErrDBItemJSONMarshal = errors.New("shipment db item json marshal failed")

	ErrConvertDBItemToStruct = errors.New("convert db item to shipment entity")

	ErrFailedSetTTL = errors.New("set ods assignment ttl failed")

	ErrJobNotFound = errors.New("job not found from db")

	ErrUnknownItemType = errors.New("unknown entity item type of db")

	ErrGetShipmentDetailByTTL = errors.New("unable to fetch shipments from ods ttl list")

	ErrNewHTTPRequest = errors.New("set http new request failed")

	ErrHTTPRequest = errors.New("do http request failed")

	ErrUnableReadRespBody = errors.New("unable read http response body")

	ErrJSONUnmarshal = errors.New("json unmarshal failed")

	ErrAdminAPIResponse = errors.New("fetch information from admin api has wrong http status")

	ErrAPIResponse = errors.New("api response status is not 200")

	ErrUnableFetchOriginalShipment = errors.New("unable fetch original shipment information from shipment module")

	ErrUnableFetchOriginOrder = errors.New("unable fetch original order information from order module")

	ErrUpdateRefIDs = errors.New("update shipment reference IDs failed")

	ErrOrgIDMismatched = errors.New("organisation ID mismatched")

	ErrJSONMarshal = errors.New("json marshal failed")

	ErrInvokeUpdateUnassignedShipment = errors.New("invoke lambda unassigned shipment failed")

	ErrNilConsolidatedShipment = errors.New("consolidated shipment must not be nil")

	ErrNilJob = errors.New("job must not be nil")

	ErrPackageHasMoreThanTwoJobs = errors.New("a package can only have maximum 2 jobs")

	ErrUnexpectedPagination = errors.New("unexpected pagination returned from database")

	ErrEmptyParamsInput = errors.New("empty params input")

	ErrConsolidatedShipmentRequiredFieldIsNil = errors.New("a required field for consolidated shipment is nil")

	ErrConvertConsolidatedPayloadToDBJob = errors.New("failed to generate DB job from consolidated payload")

	ErrConvertJobToDBJob = errors.New("failed to convert job to dbJob entity")

	ErrInvalidISOTimeString = errors.New("invalid ISO time string 2006-01-02T15:04:05GMT-07:00")

	// ErrBatchWriteJob ...
	ErrBatchWriteJob = errors.New("batch put job to dynamo failed")

	// ErrTransactWrite ...
	ErrTransactWrite = errors.New("transact write to dynamo failed")

	// ErrdbMarshalList ...
	ErrdbMarshalList = errors.New("db marshal list item failed")

	// ErrdbBatchGetItems -
	ErrdbBatchGetItems = errors.New("db batch get items failed")

	// ErrUpdateInboundScanResultFailed -
	ErrUpdateInboundScanResultFailed = errors.New("update inbound scan result failed")

	// ErrUpdatePODScanResultFailed -
	ErrUpdatePODScanResultFailed = errors.New("update POD scan result failed")

	// ErrUpdatePackageReportResultFailed -
	ErrUpdatePackageReportResultFailed = errors.New("update package report result failed")

	// ErrUpdateJobScanConfigsFailed -
	ErrUpdateJobScanConfigsFailed = errors.New("update job scan configs failed")

	// ErrDeleteJobRecord ...
	ErrDeleteJobRecord = errors.New("delete job record failed")

	// ErrNilETAStatusTTLRecord is returned when converting a nil TTL record to database entity
	ErrNilETAStatusTTLRecord = errors.New("ttl record for ETA status event is nil")

	// ErrConvertBizToDbEntity -
	ErrConvertBizToDbEntity = errors.New("convert biz entity to db entity")

	// ErrConvertDbToBizEntity -
	ErrConvertDbToBizEntity = errors.New("convert db entity to biz entity")

	ErrNoFieldUpdated = errors.New("no field updated")

	ErrInvalidUnoptimizedJob = errors.New("invalid Unoptimized Job")
)

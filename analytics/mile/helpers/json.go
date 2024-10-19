package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/prebid/prebid-server/v2/analytics"
)

func JsonifyAuctionObject(ao *analytics.AuctionObject, scope string) ([]MileAnalyticsEvent, error) {
	//var logEntry *MileAnalyticsEvent
	var events []MileAnalyticsEvent
	if ao != nil {

		var bidBiders []string
		//var configuredBidders []string
		if ao.Response != nil {

			for _, i := range ao.Response.SeatBid {
				bidBiders = append(bidBiders, i.Seat)
			}
			//}
		}

		fmt.Println(string(ao.RequestWrapper.Imp[0].Ext), "ext")

		var confBidders ImpressionsExt

		err := json.Unmarshal(ao.RequestWrapper.Imp[0].Ext, &confBidders)
		if err != nil {
			return nil, err
		}
		configuredBidders := make([]string, len(confBidders.Prebid.Bidder))
		i := 0
		for k := range confBidders.Prebid.Bidder {
			configuredBidders[i] = k
			i++
		}
		fmt.Println(configuredBidders)

		if ao.RequestWrapper != nil {

			for range ao.RequestWrapper.Imp {

				logEntry := MileAnalyticsEvent{
					//SessionID: ao.RequestWrapper
					Ip: ao.RequestWrapper.Device.IP,
					//ClientVersion: ao.RequestWrapper.Ext.
					Ua:                ao.RequestWrapper.Device.UA,
					ArbitraryData:     "",
					Device:            ao.RequestWrapper.Device.Model,
					Publisher:         ao.RequestWrapper.Site.Publisher.Domain,
					Site:              ao.RequestWrapper.Site.Domain,
					ReferrerURL:       ao.RequestWrapper.Site.Ref,
					AdvertiserName:    "",
					AuctionID:         ao.RequestWrapper.ID,
					Page:              ao.RequestWrapper.Site.Page,
					YetiSiteID:        ao.RequestWrapper.Site.ID,
					YetiPublisherID:   ao.RequestWrapper.Site.Publisher.ID,
					SessionID:         "",
					EventType:         "pbs_agg_adunit",
					Section:           "",
					BidBidders:        bidBiders,
					ConfiguredBidders: configuredBidders,
					ConfiguredTimeout: ao.RequestWrapper.TMax,
					MetaData:          map[string][]string{"prebid_server": []string{"1"}},
					//Viewability: ao.RequestWrapper.
					//WinningSize: ao.Response.SeatBi

				}

				events = append(events, logEntry)
			}
		}

	}

	//events = append(events, logEntry)
	return events, nil

}

func JsonifyVideoObject(vo *analytics.VideoObject, scope string) (*MileAnalyticsEvent, error) {
	var logEntry *MileAnalyticsEvent
	if vo != nil {
		//var request *openrtb2.BidRequest
		//if ao.RequestWrapper != nil {
		//	request = ao.RequestWrapper.BidRequest
		//}
		logEntry = &MileAnalyticsEvent{
			//Status:               ao.Status,
			//Errors:               ao.Errors,
			//Request:              request,
			//Response:             ao.Response,
			//Account:              ao.Account,
			//StartTime:            ao.StartTime,
			//HookExecutionOutcome: ao.HookExecutionOutcome,
		}
	}
	return logEntry, nil

}

func JsonifyAmpObject(ao *analytics.AmpObject, scope string) (*MileAnalyticsEvent, error) {
	var logEntry *MileAnalyticsEvent
	if ao != nil {
		//var request *openrtb2.BidRequest
		//if ao.RequestWrapper != nil {
		//	request = ao.RequestWrapper.BidRequest
		//}
		logEntry = &MileAnalyticsEvent{
			//Status:               ao.Status,
			//Errors:               ao.Errors,
			//Request:              request,
			//Response:             ao.Response,
			//Account:              ao.Account,
			//StartTime:            ao.StartTime,
			//HookExecutionOutcome: ao.HookExecutionOutcome,
		}
	}
	return logEntry, nil

}
func JsonifyCookieSync(cso *analytics.CookieSyncObject, scope string) (*MileAnalyticsEvent, error) {
	var logEntry *MileAnalyticsEvent
	if cso != nil {
		//var request *openrtb2.BidRequest
		//if ao.RequestWrapper != nil {
		//	request = ao.RequestWrapper.BidRequest
		//}
		logEntry = &MileAnalyticsEvent{
			//Status:               ao.Status,
			//Errors:               ao.Errors,
			//Request:              request,
			//Response:             ao.Response,
			//Account:              ao.Account,
			//StartTime:            ao.StartTime,
			//HookExecutionOutcome: ao.HookExecutionOutcome,
		}
	}
	return logEntry, nil

}
func JsonifySetUIDObject(so *analytics.SetUIDObject, scope string) (*MileAnalyticsEvent, error) {
	var logEntry *MileAnalyticsEvent
	if so != nil {
		//var request *openrtb2.BidRequest
		//if ao.RequestWrapper != nil {
		//	request = ao.RequestWrapper.BidRequest
		//}
		logEntry = &MileAnalyticsEvent{
			//Status:               ao.Status,
			//Errors:               ao.Errors,
			//Request:              request,
			//Response:             ao.Response,
			//Account:              ao.Account,
			//StartTime:            ao.StartTime,
			//HookExecutionOutcome: ao.HookExecutionOutcome,
		}
	}
	return logEntry, nil

}

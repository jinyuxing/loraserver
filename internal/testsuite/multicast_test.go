package testsuite

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/brocaar/loraserver/api/gw"
	"github.com/brocaar/loraserver/internal/config"
	"github.com/brocaar/loraserver/internal/downlink"
	"github.com/brocaar/loraserver/internal/gps"
	"github.com/brocaar/loraserver/internal/storage"
	"github.com/brocaar/loraserver/internal/test"
	"github.com/brocaar/lorawan"
)

type MulticastTestSuite struct {
	suite.Suite
	test.DatabaseTestSuiteBase

	Gateway        storage.Gateway
	MulticastGroup storage.MulticastGroup
}

func (ts *MulticastTestSuite) SetupSuite() {
	ts.DatabaseTestSuiteBase.SetupSuite()
	assert := require.New(ts.T())

	config.C.NetworkServer.Gateway.Backend.Backend = test.NewGatewayBackend()

	ts.Gateway = storage.Gateway{
		MAC: lorawan.EUI64{1, 1, 1, 1, 1, 1, 1, 1},
	}
	assert.NoError(storage.CreateGateway(ts.DB(), &ts.Gateway))

	ts.MulticastGroup = storage.MulticastGroup{
		GroupType:      storage.MulticastGroupB,
		MCAddr:         lorawan.DevAddr{1, 2, 3, 4},
		MCNetSKey:      lorawan.AES128Key{1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8},
		DR:             3,
		Frequency:      868300000,
		PingSlotPeriod: 32,
		FCnt:           10,
	}
	assert.NoError(storage.CreateMulticastGroup(ts.DB(), &ts.MulticastGroup))

	var dp storage.DeviceProfile
	var sp storage.ServiceProfile
	var rp storage.RoutingProfile

	assert.NoError(storage.CreateDeviceProfile(ts.DB(), &dp))
	assert.NoError(storage.CreateServiceProfile(ts.DB(), &sp))
	assert.NoError(storage.CreateRoutingProfile(ts.DB(), &rp))

	d := storage.Device{
		DevEUI:           lorawan.EUI64{2, 2, 2, 2, 2, 2, 2, 2},
		RoutingProfileID: rp.ID,
		ServiceProfileID: sp.ID,
		DeviceProfileID:  dp.ID,
	}
	assert.NoError(storage.CreateDevice(ts.DB(), &d))
	assert.NoError(storage.SaveDeviceGatewayRXInfoSet(ts.RedisPool(), storage.DeviceGatewayRXInfoSet{
		DevEUI: d.DevEUI,
		DR:     3,
		Items: []storage.DeviceGatewayRXInfo{
			{
				GatewayID: ts.Gateway.MAC,
				RSSI:      50,
				LoRaSNR:   5,
			},
		},
	}))

	assert.NoError(storage.AddDeviceToMulticastGroup(ts.DB(), d.DevEUI, ts.MulticastGroup.ID))
}

func (ts *MulticastTestSuite) TestHandleScheduleNextQueueItem() {
	now := time.Now().Round(time.Second).Add(-time.Second)
	nowGPS := gps.Time(now).TimeSinceGPSEpoch()
	nowGW := gw.Duration(nowGPS)
	dr, _ := config.C.NetworkServer.Band.Band.GetDataRate(3)
	fPort2 := uint8(2)

	testTable := []struct {
		Name               string
		QueueItems         []storage.MulticastQueueItem
		ExpectedError      error
		ExpectedQueueItems []storage.MulticastQueueItem
		ExpectedTXInfo     *gw.TXInfo
		ExpectedPHYPayload *lorawan.PHYPayload
	}{
		{
			Name: "nothing in queue",
		},
		{
			Name: "one item in queue",
			QueueItems: []storage.MulticastQueueItem{
				{
					ScheduleAt:              now,
					EmitAtTimeSinceGPSEpoch: &nowGPS,
					MulticastGroupID:        ts.MulticastGroup.ID,
					GatewayID:               ts.Gateway.MAC,
					FCnt:                    10,
					FPort:                   2,
					FRMPayload:              []byte{1, 2, 3, 4},
				},
			},
			ExpectedQueueItems: []storage.MulticastQueueItem{},
			ExpectedTXInfo: &gw.TXInfo{
				MAC:               ts.Gateway.MAC,
				Immediately:       false,
				TimeSinceGPSEpoch: &nowGW,
				Frequency:         ts.MulticastGroup.Frequency,
				Power:             14,
				DataRate:          dr,
				CodeRate:          "4/5",
			},
			ExpectedPHYPayload: &lorawan.PHYPayload{
				MHDR: lorawan.MHDR{
					MType: lorawan.UnconfirmedDataDown,
					Major: lorawan.LoRaWANR1,
				},
				MACPayload: &lorawan.MACPayload{
					FHDR: lorawan.FHDR{
						DevAddr: ts.MulticastGroup.MCAddr,
						FCnt:    10,
					},
					FPort:      &fPort2,
					FRMPayload: []lorawan.Payload{&lorawan.DataPayload{Bytes: []byte{1, 2, 3, 4}}},
				},
				MIC: lorawan.MIC{0x8, 0xb5, 0x29, 0xe8},
			},
		},
		{
			Name: "two items in queue",
			QueueItems: []storage.MulticastQueueItem{
				{
					ScheduleAt:              now,
					EmitAtTimeSinceGPSEpoch: &nowGPS,
					MulticastGroupID:        ts.MulticastGroup.ID,
					GatewayID:               ts.Gateway.MAC,
					FCnt:                    10,
					FPort:                   2,
					FRMPayload:              []byte{1, 2, 3, 4},
				},
				{
					ScheduleAt:              now,
					EmitAtTimeSinceGPSEpoch: &nowGPS,
					MulticastGroupID:        ts.MulticastGroup.ID,
					GatewayID:               ts.Gateway.MAC,
					FCnt:                    11,
					FPort:                   2,
					FRMPayload:              []byte{1, 2, 3, 4},
				},
			},
			ExpectedQueueItems: []storage.MulticastQueueItem{
				{
					ScheduleAt:              now,
					EmitAtTimeSinceGPSEpoch: &nowGPS,
					MulticastGroupID:        ts.MulticastGroup.ID,
					GatewayID:               ts.Gateway.MAC,
					FCnt:                    11,
					FPort:                   2,
					FRMPayload:              []byte{1, 2, 3, 4},
				},
			},
			ExpectedTXInfo: &gw.TXInfo{
				MAC:               ts.Gateway.MAC,
				Immediately:       false,
				TimeSinceGPSEpoch: &nowGW,
				Frequency:         ts.MulticastGroup.Frequency,
				Power:             14,
				DataRate:          dr,
				CodeRate:          "4/5",
			},
			ExpectedPHYPayload: &lorawan.PHYPayload{
				MHDR: lorawan.MHDR{
					MType: lorawan.UnconfirmedDataDown,
					Major: lorawan.LoRaWANR1,
				},
				MACPayload: &lorawan.MACPayload{
					FHDR: lorawan.FHDR{
						DevAddr: ts.MulticastGroup.MCAddr,
						FCnt:    10,
					},
					FPort:      &fPort2,
					FRMPayload: []lorawan.Payload{&lorawan.DataPayload{Bytes: []byte{1, 2, 3, 4}}},
				},
				MIC: lorawan.MIC{0x8, 0xb5, 0x29, 0xe8},
			},
		},
		{
			Name: "item discarded because of payload size",
			QueueItems: []storage.MulticastQueueItem{
				{
					ScheduleAt:              now,
					EmitAtTimeSinceGPSEpoch: &nowGPS,
					MulticastGroupID:        ts.MulticastGroup.ID,
					GatewayID:               ts.Gateway.MAC,
					FCnt:                    10,
					FPort:                   2,
					FRMPayload:              make([]byte, 300),
				},
			},
			ExpectedQueueItems: []storage.MulticastQueueItem{},
		},
	}

	for _, tst := range testTable {
		ts.T().Run(tst.Name, func(t *testing.T) {
			assert := require.New(t)
			assert.NoError(storage.FlushMulticastQueueForMulticastGroup(ts.DB(), ts.MulticastGroup.ID))

			for _, qi := range tst.QueueItems {
				assert.NoError(storage.CreateMulticastQueueItem(ts.DB(), &qi))
			}

			assert.Equal(tst.ExpectedError, downlink.ScheduleMulticastQueueBatch(1))

			items, err := storage.GetMulticastQueueItemsForMulticastGroup(ts.DB(), ts.MulticastGroup.ID)
			assert.NoError(err)
			assert.Len(items, len(tst.ExpectedQueueItems))

			for i := range items {
				assert.Equal(tst.ExpectedQueueItems[i].FCnt, items[i].FCnt)
				assert.Equal(tst.ExpectedQueueItems[i].FRMPayload, items[i].FRMPayload)
			}

			if tst.ExpectedTXInfo != nil && tst.ExpectedPHYPayload != nil {
				downlinkFrame := <-config.C.NetworkServer.Gateway.Backend.Backend.(*test.GatewayBackend).TXPacketChan
				assert.NotEqual(0, downlinkFrame.Token)
				assert.Equal(tst.ExpectedTXInfo, &downlinkFrame.TXInfo)
				assert.Equal(tst.ExpectedPHYPayload, &downlinkFrame.PHYPayload)
			} else {
				assert.Equal(0, len(config.C.NetworkServer.Gateway.Backend.Backend.(*test.GatewayBackend).TXPacketChan))
			}
		})
	}
}

func TestMulticast(t *testing.T) {
	suite.Run(t, new(MulticastTestSuite))
}

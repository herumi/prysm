package helpers_test

import (
	"bytes"
	"sort"
	"testing"

	"github.com/prysmaticlabs/go-bitfield"
	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/state"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/shared/bls"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/testutil"
)

func TestAttestationDataSlot_OK(t *testing.T) {
	deposits, _, _ := testutil.SetupInitialDeposits(t, 100)
	beaconState, err := state.GenesisBeaconState(deposits, uint64(0), &ethpb.Eth1Data{})
	if err != nil {
		t.Fatal(err)
	}
	offset := uint64(0)
	committeeCount, _ := helpers.CommitteeCount(beaconState, 0)
	expect := offset / (committeeCount / params.BeaconConfig().SlotsPerEpoch)
	attSlot, err := helpers.AttestationDataSlot(beaconState, &ethpb.AttestationData{
		Target: &ethpb.Checkpoint{Epoch: 0},
		Crosslink: &ethpb.Crosslink{
			Shard: 0,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if attSlot != expect {
		t.Errorf("Expected %d, received %d", expect, attSlot)
	}
}

func TestAttestationDataSlot_ReturnsErrorWithNilState(t *testing.T) {
	s, err := helpers.AttestationDataSlot(nil /*state*/, &ethpb.AttestationData{
		Target: &ethpb.Checkpoint{Epoch: 0},
		Crosslink: &ethpb.Crosslink{
			Shard: 0,
		},
	})
	if err != helpers.ErrAttestationDataSlotNilState {
		t.Errorf("Expected an error, but received %v", err)
		t.Logf("attestation slot=%v", s)
	}
}

func TestAttestationDataSlot_ReturnsErrorWithNilData(t *testing.T) {
	s, err := helpers.AttestationDataSlot(&pb.BeaconState{}, nil /*data*/)
	if err != helpers.ErrAttestationDataSlotNilData {
		t.Errorf("Expected an error, but received %v", err)
		t.Logf("attestation slot=%v", s)
	}
}

func TestAttestationDataSlot_ReturnsErrorWithErroneousTargetEpoch(t *testing.T) {
	deposits, _, _ := testutil.SetupInitialDeposits(t, 100)
	beaconState, err := state.GenesisBeaconState(deposits, uint64(0), &ethpb.Eth1Data{})
	if err != nil {
		t.Fatal(err)
	}
	s, err := helpers.AttestationDataSlot(beaconState, &ethpb.AttestationData{
		Target: &ethpb.Checkpoint{Epoch: 1<<63 - 1 /* Far future epoch */},
	})
	if err == nil {
		t.Error("Expected an error, but received nil")
		t.Logf("attestation slot=%v", s)
	}
}

func TestAttestationDataSlot_ReturnsErrorWhenTargetEpochLessThanCurrentEpoch(t *testing.T) {
	deposits, _, _ := testutil.SetupInitialDeposits(t, 100)
	beaconState, err := state.GenesisBeaconState(deposits, uint64(0), &ethpb.Eth1Data{})
	if err != nil {
		t.Fatal(err)
	}
	s, err := helpers.AttestationDataSlot(beaconState, &ethpb.AttestationData{
		Target: &ethpb.Checkpoint{Epoch: 2},
	})
	if err == nil {
		t.Error("Expected an error, but received nil")
		t.Logf("attestation slot=%v", s)
	}
}

func TestAggregateAttestation(t *testing.T) {
	tests := []struct {
		a1   *ethpb.Attestation
		a2   *ethpb.Attestation
		want *ethpb.Attestation
	}{
		{a1: &ethpb.Attestation{AggregationBits: []byte{}},
			a2:   &ethpb.Attestation{AggregationBits: []byte{}},
			want: &ethpb.Attestation{AggregationBits: []byte{}}},
		{a1: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x03}},
			a2:   &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x02}},
			want: &ethpb.Attestation{AggregationBits: []byte{0x03}}},
		{a1: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x02}},
			a2:   &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x03}},
			want: &ethpb.Attestation{AggregationBits: []byte{0x03}}},
	}
	for _, tt := range tests {
		got, err := helpers.AggregateAttestation(tt.a1, tt.a2)
		if err != nil {
			t.Fatal(err)
		}
		if !ssz.DeepEqual(got, tt.want) {
			t.Errorf("AggregateAttestation() = %v, want %v", got, tt.want)
		}
	}
}

func TestAggregateAttestation_OverlapFails(t *testing.T) {
	tests := []struct {
		a1 *ethpb.Attestation
		a2 *ethpb.Attestation
	}{
		{a1: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x1F}},
			a2: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x11}}},
		{a1: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0xFF, 0x85}},
			a2: &ethpb.Attestation{AggregationBits: bitfield.Bitlist{0x13, 0x8F}}},
	}
	for _, tt := range tests {
		_, err := helpers.AggregateAttestation(tt.a1, tt.a2)
		if err != helpers.ErrAttestationAggregationBitsOverlap {
			t.Error("Did not receive wanted error")
		}
	}
}

func bitlistWithAllBitsSet(length uint64) bitfield.Bitlist {
	b := bitfield.NewBitlist(length)
	for i := uint64(0); i < length; i++ {
		b.SetBitAt(i, true)
	}
	return b
}

func bitlistsWithSingleBitSet(length uint64) []bitfield.Bitlist {
	lists := make([]bitfield.Bitlist, length)
	for i := uint64(0); i < length; i++ {
		b := bitfield.NewBitlist(length)
		b.SetBitAt(i, true)
		lists[i] = b
	}
	return lists
}

func TestAggregateAttestations(t *testing.T) {
	// Each test defines the aggregation bitfield inputs and the wanted output result.
	tests := []struct {
		name   string
		inputs []bitfield.Bitlist
		want   []bitfield.Bitlist
	}{
		{
			name: "two attestations with no overlap",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000001, 0b1},
				bitfield.Bitlist{0b00000010, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000011, 0b1},
			},
		},
		{
			name:   "256 attestations with single bit set",
			inputs: bitlistsWithSingleBitSet(256),
			want: []bitfield.Bitlist{
				bitlistWithAllBitsSet(256),
			},
		},
		{
			name:   "1024 attestations with single bit set",
			inputs: bitlistsWithSingleBitSet(1024),
			want: []bitfield.Bitlist{
				bitlistWithAllBitsSet(1024),
			},
		},
		{
			name: "two attestations with overlap",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000101, 0b1},
				bitfield.Bitlist{0b00000110, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000101, 0b1},
				bitfield.Bitlist{0b00000110, 0b1},
			},
		},
		{
			name: "some attestations overlap",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00001001, 0b1},
				bitfield.Bitlist{0b00010110, 0b1},
				bitfield.Bitlist{0b00001010, 0b1},
				bitfield.Bitlist{0b00110001, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00111011, 0b1},
				bitfield.Bitlist{0b00011111, 0b1},
			},
		},
		{
			name: "some attestations produce duplicates which are removed",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000101, 0b1},
				bitfield.Bitlist{0b00000110, 0b1},
				bitfield.Bitlist{0b00001010, 0b1},
				bitfield.Bitlist{0b00001001, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00001111, 0b1}, // both 0&1 and 2&3 produce this bitlist
			},
		},
		{
			name: "two attestations where one is fully contained within the other",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000001, 0b1},
				bitfield.Bitlist{0b00000011, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000011, 0b1},
			},
		},
		{
			name: "two attestations where one is fully contained within the other reversed",
			inputs: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000011, 0b1},
				bitfield.Bitlist{0b00000001, 0b1},
			},
			want: []bitfield.Bitlist{
				bitfield.Bitlist{0b00000011, 0b1},
			},
		},
	}

	var makeAttestationsFromBitlists = func(bl []bitfield.Bitlist) []*ethpb.Attestation {
		atts := make([]*ethpb.Attestation, len(bl))
		for i, b := range bl {
			sk := bls.RandKey()
			sig := sk.Sign([]byte("dummy_test_data"), 0 /*domain*/)
			atts[i] = &ethpb.Attestation{
				AggregationBits: b,
				Data:            nil,
				CustodyBits:     nil,
				Signature:       sig.Marshal(),
			}
		}
		return atts
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := helpers.AggregateAttestations(makeAttestationsFromBitlists(tt.inputs))
			if err != nil {
				t.Fatal(err)
			}
			sort.Slice(got, func(i, j int) bool {
				return got[i].AggregationBits.Bytes()[0] < got[j].AggregationBits.Bytes()[0]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].Bytes()[0] < tt.want[j].Bytes()[0]
			})
			if len(got) != len(tt.want) {
				t.Logf("got=%v", got)
				t.Fatalf("Wrong number of responses. Got %d, wanted %d", len(got), len(tt.want))
			}
			for i, w := range tt.want {
				if !bytes.Equal(got[i].AggregationBits.Bytes(), w.Bytes()) {
					t.Errorf("Unexpected bitlist at index %d, got %b, wanted %b", i, got[i].AggregationBits.Bytes(), w.Bytes())
				}
			}
		})
	}
}

// Copyright 2022 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package iplibrary

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/TeaOSLab/EdgeCommon/pkg/configutils"
	"io"
	"math/big"
	"net"
	"strconv"
	"strings"
	"time"
)

type WriterV1 struct {
	writer *hashWriter
	meta   *Meta

	lastIPFrom     uint64 // 上一次的IP
	lastCountryId  int64
	lastProvinceId int64
	lastCityId     int64
	lastTownId     int64
	lastProviderId int64
}

func NewWriterV1(writer io.Writer, meta *Meta) *WriterV1 {
	if meta == nil {
		meta = &Meta{}
	}
	meta.Version = Version2
	meta.CreatedAt = time.Now().Unix()

	var libWriter = &WriterV1{
		writer: newHashWriter(writer),
		meta:   meta,
	}
	return libWriter
}

func (this *WriterV1) WriteMeta() error {
	metaJSON, err := json.Marshal(this.meta)
	if err != nil {
		return err
	}
	_, err = this.writer.Write(metaJSON)
	if err != nil {
		return err
	}
	_, err = this.writer.Write([]byte("\n"))
	return err
}

func (this *WriterV1) Write(ipFrom string, ipTo string, countryId int64, provinceId int64, cityId int64, townId int64, providerId int64) error {
	// validate IP
	var fromIP = net.ParseIP(ipFrom)
	if fromIP == nil {
		return errors.New("invalid 'ipFrom': '" + ipFrom + "'")
	}
	var fromIsIPv4 = configutils.IsIPv4(fromIP)
	var toIP = net.ParseIP(ipTo)
	if toIP == nil {
		return errors.New("invalid 'ipTo': " + ipTo)
	}
	var toIsIPv4 = configutils.IsIPv4(toIP)
	if fromIsIPv4 != toIsIPv4 {
		return errors.New("'ipFrom(" + ipFrom + ")' and 'ipTo(" + ipTo + ")' should have the same IP version")
	}

	var pieces = []string{}

	// 0
	if fromIsIPv4 {
		pieces = append(pieces, "")
	} else {
		pieces = append(pieces, "6")

		// we do NOT support v6 yet
		return nil
	}

	// 1
	var fromIPLong = this.ip2long(fromIP)
	var toIPLong = this.ip2long(toIP)

	if toIPLong < fromIPLong {
		fromIPLong, toIPLong = toIPLong, fromIPLong
	}

	if this.lastIPFrom > 0 && fromIPLong > this.lastIPFrom {
		pieces = append(pieces, "+"+this.formatUint64(fromIPLong-this.lastIPFrom))
	} else {
		pieces = append(pieces, this.formatUint64(fromIPLong))
	}
	this.lastIPFrom = fromIPLong
	if ipFrom == ipTo {
		// 2
		pieces = append(pieces, "")
	} else {
		// 2
		pieces = append(pieces, this.formatUint64(toIPLong-fromIPLong))
	}

	// 3
	if countryId > 0 {
		if countryId == this.lastCountryId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(countryId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastCountryId = countryId

	// 4
	if provinceId > 0 {
		if provinceId == this.lastProvinceId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(provinceId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastProvinceId = provinceId

	// 5
	if cityId > 0 {
		if cityId == this.lastCityId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(cityId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastCityId = cityId

	// 6
	if townId > 0 {
		if townId == this.lastTownId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(townId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastTownId = townId

	// 7
	if providerId > 0 {
		if providerId == this.lastProviderId {
			pieces = append(pieces, "+")
		} else {
			pieces = append(pieces, this.formatUint64(uint64(providerId)))
		}
	} else {
		pieces = append(pieces, "")
	}
	this.lastProviderId = providerId

	_, err := this.writer.Write([]byte(strings.TrimRight(strings.Join(pieces, "|"), "|")))
	if err != nil {
		return err
	}

	_, err = this.writer.Write([]byte("\n"))
	return err
}

func (this *WriterV1) Sum() string {
	return this.writer.Sum()
}

func (this *WriterV1) formatUint64(i uint64) string {
	return strconv.FormatUint(i, 32)
}

func (this *WriterV1) ip2long(netIP net.IP) uint64 {
	if len(netIP) == 0 {
		return 0
	}

	var b4 = netIP.To4()
	if b4 != nil {
		return uint64(binary.BigEndian.Uint32(b4.To4()))
	}

	var i = big.NewInt(0)
	i.SetBytes(netIP.To16())
	return i.Uint64()
}

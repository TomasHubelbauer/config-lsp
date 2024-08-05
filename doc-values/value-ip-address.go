package docvalues

import (
	"config-lsp/utils"
	net "net/netip"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var NonRoutableNetworks = []net.Prefix{
	net.MustParsePrefix("240.0.0.0/4"),
	net.MustParsePrefix("2001:db8::/32"),
}

type InvalidIPAddress struct{}

func (e InvalidIPAddress) Error() string {
	return "This is not a valid IP Address"
}

type IP4AddressNotAllowedError struct{}

func (e IP4AddressNotAllowedError) Error() string {
	return "IPv4 Addresses are not allowed"
}

type IP6AddressNotAllowedError struct{}

func (e IP6AddressNotAllowedError) Error() string {
	return "IPv6 Addresses are not allowed"
}

type IpRangeNotAllowedError struct{}

func (e IpRangeNotAllowedError) Error() string {
	return "IP Ranges are not allowed; Use a single IP Address instead"
}

type IPAddressNotAllowedError struct{}

func (e IPAddressNotAllowedError) Error() string {
	return "This IP Address is not allowed"
}

type IPAddressValue struct {
	AllowIPv4     bool
	AllowIPv6     bool
	AllowRange    bool
	AllowedIPs    *[]net.Prefix
	DisallowedIPs *[]net.Prefix
}

func (v IPAddressValue) GetTypeDescription() []string {
	if v.AllowedIPs != nil && len(*v.AllowedIPs) != 0 {
		return append(
			[]string{"One of the following IP Addresses (in range):"},
			utils.Map(*v.AllowedIPs, func(ip net.Prefix) string {
				return ip.String()
			})...,
		)
	}

	return []string{"An IP Address"}
}

func (v IPAddressValue) CheckIsValid(value string) []*InvalidValue {
	var ip net.Prefix

	if v.AllowRange {
		rawIP, err := net.ParsePrefix(value)

		if err != nil {
			return []*InvalidValue{
				{
					Err:   InvalidIPAddress{},
					Start: 0,
					End:   uint32(len(value)),
				},
			}
		}

		ip = rawIP
	} else {
		rawIP, err := net.ParseAddr(value)

		if err != nil {
			return []*InvalidValue{{
				Err:   InvalidIPAddress{},
				Start: 0,
				End:   uint32(len(value)),
			},
			}
		}

		ip = net.PrefixFrom(rawIP, 32)
	}

	if !ip.IsValid() {
		return []*InvalidValue{{
			Err:   InvalidIPAddress{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if v.AllowedIPs != nil {
		for _, allowedIP := range *v.AllowedIPs {
			if allowedIP.Contains(ip.Addr()) {
				return nil
			}
		}

		return []*InvalidValue{{
			Err:   IPAddressNotAllowedError{},
			Start: 0,
			End:   uint32(len(value)),
		},
		}
	}

	if v.DisallowedIPs != nil {
		for _, disallowedIP := range *v.DisallowedIPs {
			if disallowedIP.Contains(ip.Addr()) {
				return []*InvalidValue{{
					Err:   IPAddressNotAllowedError{},
					Start: 0,
					End:   uint32(len(value)),
				},
				}
			}
		}
	}

	if v.AllowIPv4 && ip.Addr().Is4() {
		return nil
	}

	if v.AllowIPv6 && ip.Addr().Is6() {
		return nil
	}

	return []*InvalidValue{{
		Err:   InvalidIPAddress{},
		Start: 0,
		End:   uint32(len(value)),
	},
	}
}

func (v IPAddressValue) FetchCompletions(line string, cursor uint32) []protocol.CompletionItem {
	if v.AllowedIPs != nil && len(*v.AllowedIPs) != 0 {
		kind := protocol.CompletionItemKindValue

		return utils.Map(*v.AllowedIPs, func(ip net.Prefix) protocol.CompletionItem {
			return protocol.CompletionItem{
				Label: ip.Addr().String(),
				Kind:  &kind,
			}
		})
	}

	return []protocol.CompletionItem{}
}
